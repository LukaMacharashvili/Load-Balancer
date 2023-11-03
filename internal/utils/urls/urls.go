package urls

import (
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

func GetTargetUrlsFromEnv() []string {
	commaSeparatedTargetUrls := os.Getenv("TARGET_URLS")
	return strings.Split(commaSeparatedTargetUrls, ",")
}

func GetProxyFromUrlString(targetUrl string) *httputil.ReverseProxy {
	url, _ := url.Parse(targetUrl)
	return httputil.NewSingleHostReverseProxy(url)
}
