package url

import "net/url"

func ParseUrl(rawUrl string) (hostName string, err error) {
	u, err := url.Parse(rawUrl)
	if err != nil {
		return "", err
	}

	hostName = u.Hostname()
	return
}

