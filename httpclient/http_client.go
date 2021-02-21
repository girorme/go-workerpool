package httpclient

import (
	"fmt"
	"net/http"
)

type HttpClient struct {
}

func (h *HttpClient) Get(url string) string {
	resp, err := http.Get(url)

	if err != nil {
		return fmt.Sprintf("[-][%s] err: %s\n", url, err.Error())

	}

	return fmt.Sprintf("[+][%s] http_code %d - %s \n", url, resp.StatusCode, http.StatusText(resp.StatusCode))
}
