package main

import (
	"fmt"
	"net/http"
)

func main() {
	client := &http.Client{
		CheckRedirect: func(req * http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
	}}

	resp, err := client.Get("http://studio3marketing.com")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("HTTP Status", resp.StatusCode)
	fmt.Println("HTTP URL", resp.Request.URL)
}
