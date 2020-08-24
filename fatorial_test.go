package main

import "testing"

func TestFatorial(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "When calc is correct", args: args{n: 4}, want: 24},
		{name: "When calc is correct", args: args{n: 5}, want: 120},
		{name: "When calc is correct", args: args{n: 0}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fatorial(tt.args.n); got != tt.want {
				t.Errorf("Fatorial() = %v, want %v", got, tt.want)
			}
		})
	}
}
