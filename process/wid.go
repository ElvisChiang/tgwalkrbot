package process

import (
	"fmt"
	"strings"
)

const widLEN = 11
const urlPrefix = "http://walker.tie.tw/"

// GetWidURL parse and return url from private string
func GetWidURL(wid string) (url string, found bool) {
	found = false

	wid = strings.TrimSpace(wid)

	urls := strings.Fields(wid)

	for _, l := range urls {
		fmt.Printf("l = %s len = %d\n", l, len(l))
		if len(l) != widLEN {
			found = false
			return
		}
		url = url + urlPrefix + l + "\n"
	}
	fmt.Printf("url:\n%s", url)
	found = true
	return
}
