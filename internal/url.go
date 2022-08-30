package internal

import (
	"net/url"
	"strings"
)

func EncodeQuery(u string) (string, error) {
	uu, e := url.Parse(u)
	if e != nil {
		return "", e
	}
	uu.RawQuery = uu.Query().Encode()
	return uu.String(), nil
}

func WebpToPng(s string) string {
	if !strings.HasSuffix(s, "webp") {
		return s
	}
	return strings.TrimSuffix(s, "webp") + "png"
}
