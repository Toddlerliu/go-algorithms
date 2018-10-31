package utils

import (
	"math"
	"encoding/binary"
	"unsafe"
)

func Float64ToBytes(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, bits)
	return bytes
}

func BytesToFloat64(bytes []byte) float64 {
	bits := binary.BigEndian.Uint64(bytes)
	return math.Float64frombits(bits)
}

// String转Bytes
// runtime
//func slicebytetostring(b Slice) (s String) {
//	void *pc;
//	if(raceenabled) {
//		pc = runtime.getcallerpc(&b);
//		runtime.racereadrangepc(b.array, b.len, pc, runtime.slicebytetostring);
//	}
//	s = gostringsize(b.len);
//	runtime.memmove(s.str, b.array, s.len);
//}
//
//func stringtoslicebyte(s String) (b Slice) {
//	b.array = runtime.mallocgc(s.len, 0, FlagNoScan|FlagNoZero);
//	b.len = s.len;
//	b.cap = s.len;
//	runtime.memmove(b.array, s.str, s.len);
//}
// 无论从[]byte到string还是string到[]byte，他们的指针地址均不同。
// 在类型转换的时候，发生了值拷贝，而[]byte与string不共享内存
func StringToBytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func BytesToString(byte []byte) string {
	return *(*string)(unsafe.Pointer(&byte))
}
