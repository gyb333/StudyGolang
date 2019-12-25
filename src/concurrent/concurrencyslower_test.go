package concurrent

import "testing"

//go test -bench .

//func BenchmarkSerialSum(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		SerialSum()
//	}
//}

func BenchmarkConcurrentSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcurrentSum()
	}
}

func BenchmarkChannelSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ChannelSum()
	}
}