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
		{
			name: "skip one line with new line",
			args: args{
				data: "Album: Stroke in Time \nGenre: Blues Rock / Country Rock \nCountry: USA \nReleased: 2023 \nQuality: MP3 320 / FLAC \nTracklist: \n01. Leave It to Chance \n02. Shiny Globe \n03. Sweet Spot \n04. Lady and Lasalle \n05. Blind Men \n06. Sweeping (Out) the Corners \n07. Backyard Burning \n08. Someone Else's War \n",
			},
			want: true,
		},
		{
			name: "skip new line",
			args: args{
				data: `Album: Stroke in Time 
Genre: Blues Rock / Country Rock 
Country: USA 
Released: 2023 
Quality: MP3 320 / FLAC 
Tracklist: 
01. Leave It to Chance 
02. Shiny Globe 
03. Sweet Spot 
04. Lady and Lasalle 
05. Blind Men 
06. Sweeping (Out) the Corners 
07. Backyard Burning 
08. Someone Else's War 
`,
			},
			want: true,
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
