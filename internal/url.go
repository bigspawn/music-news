package internal

import "net/url"

func EncodeQuery(u string) (string, error) {
	uu, e := url.Parse(u)
	if e != nil {
		return "", e
	}
	uu.RawQuery = uu.Query().Encode()
	return uu.String(), nil
}
