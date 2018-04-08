package main

// 改编自Twitter：
// https://github.com/twitter/snowflake/blob/snowflake-2010/src/main/scala/com/twitter/service/snowflake/IdWorker.scala
import (
	"errors"
	"sync"
	"time"
)

// 	 0 - 00000000 00000000 00000000 00000000 00000000 0 - 00000000 00 - 00000000 0000
// 1bit                   41bit时间戳                    10bit工作机器id 12bit序列号
// 1bit：最高位；0代表正；1代表负；
// 41bit：时间差：2^41/(1000*60*60*24*365)=69.7年；
// 10bit：2^10=1024节点，包括5位datacenterId和5位workerId；
// 12bit：2^12=4096个id，1ms 1node 产生4096个id
const (
	// 起始时间戳 (ms) 2006-01-02 15:04:05
	epoch int64 = 1136214245000
	// 工作机器id
	workerIdBits uint = 5

	datacenterIdBits uint = 5
	// 最大编号
	maxWorkerId int64 = -1 ^ (-1 << workerIdBits) // 31

	maxDatacenterId int64 = -1 ^ (-1 << datacenterIdBits) // 31
	// 序列号
	sequenceBits uint = 12
	// 最大序列号
	//maxSequenceId int64 = -1 ^ (-1 << sequenceBits)

	workerIdShift      = sequenceBits                                   //12
	datacenterIdShift  = sequenceBits + workerIdBits                    // 17
	timestampLeftShift = sequenceBits + workerIdBits + datacenterIdBits // 22
	sequenceMask       = -1 ^ (-1 << sequenceBits)                      // 0x3ff 4095
)

type IdWorker struct {
	mu            sync.Mutex
	workerId      int64
	datacenterId  int64
	lastTimestamp int64
	sequence      int64
}

func NewIdWorker(workerId, datacenterId int64) (*IdWorker, error) {
	worker := new(IdWorker)
	if workerId < 0 || workerId > maxWorkerId {
		return nil, errors.New("the workerId must between 0 and 31")
	}
	if datacenterId > maxDatacenterId || datacenterId < 0 {
		return nil, errors.New("the datacenterId must between 0 and 31")
	}
	worker.workerId = workerId
	worker.datacenterId = datacenterId
	worker.lastTimestamp = -1
	worker.sequence = 0
	return worker, nil
}

func (node *IdWorker) GenerateID() (int64, error) {
	node.mu.Lock()
	defer node.mu.Unlock()
	// ms
	now := time.Now().UnixNano() / 1000 / 1000

	if now < node.lastTimestamp {
		return 0, errors.New("error time")
	}

	if now == node.lastTimestamp {
		node.sequence = (node.sequence + 1) & sequenceMask // 1000000000000(4096) & 111111111111(4095) = 0
		// 本ms内sequence分配完(4096)
		if node.sequence == 0 {
			now = tilNextMillis(node.lastTimestamp)
		}
	} else {
		node.sequence = 0
	}

	node.lastTimestamp = now

	return (node.lastTimestamp-epoch)<<timestampLeftShift | node.datacenterId<<datacenterIdShift | node.workerId<<workerIdShift | node.sequence, nil
}

func tilNextMillis(lastTimestamp int64) int64 {
	tmp := time.Now().UnixNano() / 1000 / 1000
	for tmp <= lastTimestamp {
		tmp = time.Now().UnixNano() / 1000 / 1000
	}
	return tmp
}
