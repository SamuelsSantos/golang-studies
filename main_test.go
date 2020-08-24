package main

import (
	"testing"
)


func Test_isBalance(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{name: "When is valid", args: args{value: "[{()}]"}, want: true, wantErr: false},
		{name: "When is invalid", args: args{value: "[[{}]"}, want: false, wantErr: true},
		{name: "When is invalid character", args: args{value: "[{|}]"}, want: false, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := isBalance(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("isBalance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("isBalance() = %v, want %v", got, tt.want)
			}
		})
	}
}
