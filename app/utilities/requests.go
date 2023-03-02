package utilities

import (
	"bytes"
	"io"
	"net/http"
)

func GetRequest(url string, token string) io.ReadCloser {
	req, _ := http.NewRequest("GET", url, bytes.NewBuffer([]byte("")))
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	return resp.Body
}
