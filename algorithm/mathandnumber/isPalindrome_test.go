package mathandnumber

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "basic",
			args: args{
				x: 121,
			},
			want: true,
		},
		{
			name: "negative",
			args: args{
				x: -121,
			},
			want: false,
		},
		{
			name: "zero",
			args: args{
				x: 0,
			},
			want: true,
		},
		{
			name: "trailingZero",
			args: args{
				x: 10,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
