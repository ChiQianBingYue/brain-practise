package stringproblem

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "base",
			args: args{
				s: "",
			},
			want: true,
		},
		{
			name: "base2",
			args: args{
				s: "{",
			},
			want: false,
		},
		{
			name: "base3",
			args: args{
				s: "()[]{}",
			},
			want: true,
		},
		{
			name: "base4",
			args: args{
				s: "([)]",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
