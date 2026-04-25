package misc

import (
	"regexp"
	"strings"
)

func AdaptUrlToFileName(url string) string {
	re := regexp.MustCompile(`[^\pL]+`)
	url = strings.ReplaceAll(url, `\`, "")
	return re.ReplaceAllString(url, "_")
}
