package internal

import "strings"

const (
	alterportalHost  = "alterportal.net"
	getrockmusicHost = "getrockmusic.net"
	coreradioHost    = "coreradio.online"
)

var fileHosts = map[string]struct{}{
	"1fichier.com":     {},
	"cloud.mail.ru":    {},
	"coreradio.ru":     {},
	"disk.yandex.ru":   {},
	"drive.google.com": {},
	"filecrypt.cc":     {},
	"files.fm":         {},
	"files.mail.ru":    {},
	"krakenfiles.com":  {},
	"mediafire.com":    {},
	"mega.nz":          {},
	"megaup.net":       {},
	"solidfiles.com":   {},
	"t.me":             {}, // Telegram links support
	"telegram.me":      {}, // Alternative Telegram domain
	"trbbt.net":        {},
	"turb.cc":          {},
	"turb.pw":          {},
	"turbobit.net":     {},
	"uppit.com":        {},
	"yadi.sk":          {},
	"zippyshare.com":   {},
	alterportalHost:    {},
	coreradioHost:      {},
	getrockmusicHost:   {},
}

func isAllowedFileHost(host string) bool {
	for s, _ := range fileHosts {
		if strings.Contains(host, s) {
			return true
		}
	}
	return false
}
