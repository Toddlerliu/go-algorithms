package main

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
	workerIdBits uint = 10
	// 序列号
	sequenceIdBits uint = 12
	// 最大编号
	maxWorkerId int64 = -1 ^ (-1 << workerIdBits) // 0xfff 1024
	// 最大序列号
	maxSequenceId int64 = -1 ^ (-1 << sequenceIdBits) // 0x3ff 4096
)

// 工作节点 最大1024个
type WorkerNode struct {
	mu            sync.Mutex
	nodeId        int64
	lastTimestamp int64
	sequence      int64
}

func NewWorkerNode(nodeId int64) (*WorkerNode, error) {
	worker := new(WorkerNode)
	if nodeId < 0 || nodeId > maxWorkerId {
		return nil, errors.New("the nodeId must between 0 and 1024")
	}
	worker.nodeId = nodeId
	worker.lastTimestamp = 0
	worker.sequence = 0
	return worker, nil
}

func (node *WorkerNode) GenerateID() (int64, error) {
	node.mu.Lock()
	defer node.mu.Unlock()
	// ms
	now := time.Now().UnixNano() / 1000 / 1000
	if node.lastTimestamp == now {
		node.sequence =(node.sequence + 1) & maxWorkerId
		// 本ms内sequence分配完（>4096）
		if node.sequence > maxSequenceId {
			for node.lastTimestamp >= now {
				now = time.Now().UnixNano()
			}
		}
	} else if node.lastTimestamp > now {
		return 0, errors.New("error time")
	} else {
		// 本毫秒内 sequence 用完
		node.sequence = 0
	}
	node.lastTimestamp = now
	return (now-epoch)<<(workerIdBits+sequenceIdBits) | node.nodeId<<12 | node.sequence, nil
}
