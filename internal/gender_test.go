package internal

import "testing"

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
			name: "skip it",
			args: args{
				data: skipGenders[0],
			},
			want: true,
		},
		{
			name: "skip it long",
			args: args{
				data: "slakajfsasf" + skipGenders[0] + "sdgadsgsdg",
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
			if got := isSkippedGender(tt.args.data); got != tt.want {
				t.Errorf("isSkippedGender() = %v, want %v", got, tt.want)
			}
		})
	}
}
