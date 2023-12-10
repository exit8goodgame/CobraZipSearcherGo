package service

import (
	"reflect"
	"sync"
	"testing"
)

func TestNewSearchService(t *testing.T) {
	type args struct {
		targetCharList string
		targetCharType int
		passMaxLen     int
	}
	tests := []struct {
		name string
		args args
		want *SearchService
	}{
		{
			name: "ok",
			args: args{
				targetCharList: "",
				targetCharType: 0,
				passMaxLen:     0,
			},
			want: &SearchService{
				safeData: &SafeData{
					pass: "",
					done: false,
					mux:  sync.Mutex{},
				},
				targetChar: TargetCharDefault,
				passMaxLen: PassMaxLenDefault,
			},
		},
		{
			name: "ok. '', 1, 1",
			args: args{
				targetCharList: "",
				targetCharType: 1,
				passMaxLen:     1,
			},
			want: &SearchService{
				safeData: &SafeData{
					pass: "",
					done: false,
					mux:  sync.Mutex{},
				},
				targetChar: TargetCharShort,
				passMaxLen: 1,
			},
		},
		{
			name: "ok. '', 2, 2",
			args: args{
				targetCharList: "",
				targetCharType: 2,
				passMaxLen:     2,
			},
			want: &SearchService{
				safeData: &SafeData{
					pass: "",
					done: false,
					mux:  sync.Mutex{},
				},
				targetChar: TargetCharLong,
				passMaxLen: 2,
			},
		},
		{
			name: "ok. '', 3, 3",
			args: args{
				targetCharList: "",
				targetCharType: 3,
				passMaxLen:     3,
			},
			want: &SearchService{
				safeData: &SafeData{
					pass: "",
					done: false,
					mux:  sync.Mutex{},
				},
				targetChar: TargetCharDefault,
				passMaxLen: 3,
			},
		},
		{
			name: "ok. '123abcABC', 1, 1",
			args: args{
				targetCharList: "123abcABC",
				targetCharType: 1,
				passMaxLen:     1,
			},
			want: &SearchService{
				safeData: &SafeData{
					pass: "",
					done: false,
					mux:  sync.Mutex{},
				},
				targetChar: "123abcABC",
				passMaxLen: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSearchService(tt.args.targetCharList, tt.args.targetCharType, tt.args.passMaxLen); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSearchService() = %v, want %v", got, tt.want)
			}
		})
	}
}
