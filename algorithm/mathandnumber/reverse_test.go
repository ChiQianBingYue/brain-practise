package mathandnumber

import "testing"

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "basic",
			args: args{
				x: 123,
			},
			want: 321,
		},
		{
			name: "nega",
			args: args{
				x: -1234,
			},
			want: -4321,
		},
		{
			name: "zero",
			args: args{
				x: 1534236469,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
