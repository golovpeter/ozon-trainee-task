package common

import "strings"

func GetAlias(url string) string {
	urlElements := strings.Split(url, "/")
	return urlElements[len(urlElements)-1]
}
