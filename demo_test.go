package goroutine_demo

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestSumShared(t *testing.T) {
	type args struct {
		allNums [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "should return 500,005,000,000,000",
			args: args{
				allNums: generateVLA(),
			},
			want: 500005000000000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SumShared(tt.args.allNums)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSumChannel(t *testing.T) {
	type args struct {
		allNums [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "should return 500,005,000,000,000",
			args: args{
				allNums: generateVLA(),
			},
			want: 500005000000000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SumChannel(tt.args.allNums)
			assert.Equal(t, tt.want, got)
		})
	}
}

//func TestSumSharedNoMutex(t *testing.T) {
//	type args struct {
//		allNums [][]int
//	}
//	tests := []struct {
//		name string
//		args args
//		want int
//	}{
//		{
//			name: "should return 500,005,000,000,000",
//			args: args{
//				allNums: generateVLA(),
//			},
//			want: 500005000000000,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got := SumSharedNoMutex(tt.args.allNums)
//			assert.Equal(t, tt.want, got)
//		})
//	}
//}

func generateVLA() [][]int {
	highest := 100000
	baseArr := make([]int, highest)
	for i := 0; i < highest; i++ {
		baseArr[i] = i + 1
	}

	finalArr := make([][]int, highest)
	for i := 0; i < highest; i++ {
		finalArr[i] = baseArr
	}

	return finalArr
}
