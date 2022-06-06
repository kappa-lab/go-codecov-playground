package main

import "testing"

func TestIsTrueByString(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "`true` is true",
			args: args{value: "true"},
			want: true,
		},
		{
			name: "`TRUE` is true",
			args: args{value: "TRUE"},
			want: true,
		},
		{
			name: "`false` is false",
			args: args{value: "false"},
			want: false,
		},
		{
			name: "`` is false",
			args: args{value: ""},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsTrueByString(tt.args.value); got != tt.want {
				t.Errorf("IsTrueByString() = %v, want %v", got, tt.want)
			}
		})
	}
}
