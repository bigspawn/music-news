package internal

import (
	"testing"

	"github.com/go-pkgz/lgr"
)

func Test_isSkippedGender(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "skip gender",
			args: args{
				data: "Genre: Gothabilly",
			},
			want: true,
		},
		{
			name: "skip gender",
			args: args{
				data: "Genre: Hillbilly / Rockabilly ",
			},
			want: true,
		},
		{
			name: "skip gender without space",
			args: args{
				data: "Genre:Gothabilly",
			},
			want: true,
		},
		{
			name: "skip gender in the middle",
			args: args{
				data: "Genre: slakajfsasf" + "Gothabilly" + "sdgadsgsdg",
			},
			want: true,
		},
		{
			name: "ok",
			args: args{
				data: "some text",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSkippedGenre(lgr.Default(), tt.args.data); got != tt.want {
				t.Errorf("isSkippedGenre() = %v, want %v", got, tt.want)
			}
		})
	}
}
