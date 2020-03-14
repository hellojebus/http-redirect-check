package main

import (
	"fmt"
	"net/http"
)

func main() {

	nextUrl := "http://studio3marketing.com"
	var i int
	for i < 100 {
		output := ""
		client := &http.Client{
			CheckRedirect: func(req * http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
		}}

		resp, err := client.Get(nextUrl)

		if err != nil {
			fmt.Println(err)
		}

		output += string(resp.StatusCode)

		if resp.StatusCode == 200 {
			fmt.Println("Final", nextUrl)
			break
		} else {
			fmt.Println("Redirecting", nextUrl, "to", resp.Header.Get("Location"))
			nextUrl = resp.Header.Get("Location")
			i += 1
		}

		fmt.Println(output)
	}

}
