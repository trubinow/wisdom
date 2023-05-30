package utils

import (
	"testing"
	"time"
)

func TestStringToDuration(t *testing.T) {
	var empty time.Duration
	type args struct {
		durationStr string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Duration
		wantErr bool
	}{
		{
			name:    "Valid 10s",
			args:    struct{ durationStr string }{durationStr: "10s"},
			want:    10 * time.Second,
			wantErr: false,
		},
		{
			name:    "Valid 2m",
			args:    struct{ durationStr string }{durationStr: "2m"},
			want:    2 * time.Minute,
			wantErr: false,
		},
		{
			name:    "InValid 1",
			args:    struct{ durationStr string }{durationStr: "10ss"},
			want:    empty,
			wantErr: true,
		},
		{
			name:    "InValid 2",
			args:    struct{ durationStr string }{durationStr: "1000h"},
			want:    empty,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringToDuration(tt.args.durationStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringToDuration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("StringToDuration() got = %v, want %v", got, tt.want)
			}
		})
	}
}
