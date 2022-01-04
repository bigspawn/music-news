package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractLinkAlt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "alterPortal",
			args: args{
				s: "https://alterportal.net/index.php?do=go&amp;url=aHR0cHM6Ly9jbG91ZC5tYWlsLnJ1L3B1YmxpYy9LdTVmL2trR0dnM1VzWg%3D%3D",
			},
			want: "aHR0cHM6Ly9jbG91ZC5tYWlsLnJ1L3B1YmxpYy9LdTVmL2trR0dnM1VzWg==",
		},
		{
			name: "getRock",
			args: args{
				s: "https://getrockmusic.net/index.php?do=go&url=aHR0cHM6Ly90dXJiLmNjL2dkY3l0cWp3cWZlcC5odG1s",
			},
			want: "aHR0cHM6Ly90dXJiLmNjL2dkY3l0cWp3cWZlcC5odG1s",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExtractLinkFromParamURL(tt.args.s)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestDecodeBase64StdPadding(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "1",
			args: args{
				s: "aHR0cHM6Ly9jbG91ZC5tYWlsLnJ1L3B1YmxpYy9LdTVmL2trR0dnM1VzWg==",
			},
			want:    "https://cloud.mail.ru/public/Ku5f/kkGGg3UsZ",
			wantErr: nil,
		},
		{
			name: "2",
			args: args{
				s: "aHR0cHM6Ly90dXJiLmNjL2dkY3l0cWp3cWZlcC5odG1s",
			},
			want:    "https://turb.cc/gdcytqjwqfep.html",
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DecodeBase64StdPadding(tt.args.s)
			assert.NoError(t, err)
			assert.Equalf(t, tt.want, got, "DecodeBase64Std(%v)", tt.args.s)
		})
	}
}
