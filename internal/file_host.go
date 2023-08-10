package internal

import "strings"

const (
	alterportalHost  = "alterportal.net"
	getrockmusicHost = "getrockmusic.net"
)

var fileHosts = []string{
	"mediafire.com",
	"zippyshare.com",
	"mega.nz",
	"solidfiles.com",
	"drive.google.com",
	"files.mail.ru",
	"disk.yandex.ru",
	"yadi.sk",
	"files.fm",
	"uppit.com",
	"filecrypt.cc",
	"turb.cc",
	"turbobit.net",
	"coreradio.ru",
	alterportalHost,
	getrockmusicHost,
	"turb.pw",
	"krakenfiles.com",
	"trbbt.net",
	"drive.google.com",
	"megaup.net",
	"1fichier.com",
}

func isAllowedFileHost(host string) bool {
	for _, s := range fileHosts {
		if strings.Contains(host, s) {
			return true
		}
	}
	return false
}
