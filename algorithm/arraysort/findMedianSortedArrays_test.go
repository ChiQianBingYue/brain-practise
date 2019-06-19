package arraysort

import (
	"reflect"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "simple",
			args: args{nums1: []int{1, 2}, nums2: []int{2, 3}},
			want: 2.0,
		},
		// {name: "long",
		// 	args: args{nums1: []int{1, 5}, nums2: []int{2, 3, 4, 6}},
		// 	want: 3.5,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
