package mediatypes

import (
	"reflect"
	"testing"
)

func TestByExtension(t *testing.T) {
	type args struct {
		ext string
	}
	tests := []struct {
		name string
		args args
		want []MediaType
	}{
		{
			name: "empty extension",
			args: args{ext: ""},
			want: nil,
		},
		{
			name: "unknown extension",
			args: args{ext: "unknown"},
			want: nil,
		},
		{
			name: "singular",
			args: args{ext: "gif"},
			want: []MediaType{
				mediaTypes[2155],
			},
		},
		{
			name: "multiple",
			args: args{ext: "xml"},
			want: []MediaType{
				mediaTypes[65], mediaTypes[414], mediaTypes[1886], mediaTypes[2524],
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := ByExtension(tt.args.ext); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("ByExtension() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
