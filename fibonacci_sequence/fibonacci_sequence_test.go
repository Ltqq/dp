package fibonacci_sequence

import "testing"

func Test_climbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "climbStairs 3",
			args: args{
				n: 3,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs(tt.args.n); got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fib(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "fib 3",
			args: args{
				n: 3,
			},
			// start with 0
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib(tt.args.n); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_recursionClimbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "recursionClimbStairs 3",
			args: args{3},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := recursionClimbStairs(tt.args.n); got != tt.want {
				t.Errorf("recursionClimbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeDupRecursionClimbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "removeDupRecursionClimbStairs 3",
			args: args{3},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDupRecursionClimbStairs(tt.args.n); got != tt.want {
				t.Errorf("removeDupRecursionClimbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_iterativeClimbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "iterativeClimbStairs 3",
			args: args{3},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := iterativeClimbStairs(tt.args.n); got != tt.want {
				t.Errorf("iterativeClimbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
