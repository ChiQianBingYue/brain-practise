package stringproblem

import (
	"reflect"
	"testing"
)

func Test_generateSpiralMatrix(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			name: "basic",
			args: args{
				n: 3,
			},
			want: [][]int{
				[]int{1, 2, 3},
				[]int{8, 9, 4},
				[]int{7, 6, 5},
			},
		},
		{
			name: "basic",
			args: args{
				n: 0,
			},
			want: [][]int{},
		},
		{
			name: "basic",
			args: args{
				n: 1,
			},
			want: [][]int{
				[]int{1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateSpiralMatrix(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateSpiralMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
