package helpers

import "net/url"

func parseURL(u string) (*url.URL, error) {
	pu, err := url.Parse(u)
	if err != nil {
		return nil, err
	}
	return pu, nil
}

// GetPathname returns pathname
func GetPathname() {}

// GetHost returns host
func GetHost() {}

// GetURLData returns the url data
func GetURLData() {}
