package store

import (
	"devops/internal/metrics"
	"reflect"
	"testing"
)

func TestMemStorage_Set(t *testing.T) {
	type args struct {
		key   string
		value interface{}
	}
	var m MemStorage
	m.Init()
	tests := []struct {
		name string
		c    *MemStorage
		args args
	}{
		{
			name: "Simple test gauge",
			c:    &m,
			args: args{
				key:   "Alloc",
				value: 1.555,
			},
		},
		{
			name: "Simple test counter",
			c:    &m,
			args: args{
				key:   "Alloc",
				value: 8,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.Set(tt.args.key, tt.args.value)
		})
	}
}

func TestMemStorage_Get(t *testing.T) {
	type args struct {
		key string
	}
	var m MemStorage
	m.Init()
	tests := []struct {
		name  string
		c     *MemStorage
		args  args
		want  interface{}
		want1 bool
	}{
		{
			name: "Simple test gauge exist",
			c:    &m,
			args: args{
				key: "Alloc",
			},
			want:  metrics.Gauge(0),
			want1: true,
		},
		{
			name: "Simple test gauge doesn't exist",
			c:    &m,
			args: args{
				key: "Alloccc",
			},
			want:  nil,
			want1: false,
		},
		{
			name: "Simple test count exist",
			c:    &m,
			args: args{
				key: "PollCount",
			},
			want:  metrics.Counter(0),
			want1: true,
		},
		{
			name: "Simple test count doesn't exist",
			c:    &m,
			args: args{
				key: "PollCounterrrrr",
			},
			want:  nil,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.c.Get(tt.args.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MemStorage.Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MemStorage.Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
