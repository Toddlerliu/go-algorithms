package main

// 	 0 - 00000000 00000000 00000000 00000000 00000000 0 - 00000000 00 - 00000000 0000
// 1bit                   41bit时间戳                    10bit工作机器id 12bit序列号
// 1bit：最高位；0代表正；1代表负；
// 41bit：时间差：2^41/(1000*60*60*24*365)=69.7年；
// 10bit：2^10=1024节点，包括5位datacenterId和5位workerId；
// 12bit：2^12=4096个id，1ms 1node 产生4096个id
const (
	// 起始时间戳 (ms) 2006-03-21:20:50:14
	epoch int64 = 1136214245000
	// 工作机器id
	workerIdBits uint = 10
	// 序列号
	sequenceIdBits uint = 12
	// 最大编号
	maxWorkerId int64 = -1^(-1 << workerIdBits) // 0xfff
	// 最大序列号
	maxSequenceId int64 = -1^(-1 << sequenceIdBits) // 0x3ff
)

type WorkerNode struct {

}
