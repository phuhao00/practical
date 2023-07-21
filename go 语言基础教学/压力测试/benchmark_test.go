package 压力测试

import "testing"

func BenchmarkB(b *testing.B) {
	for i := 0; i < b.N; i++ {

	}
}
