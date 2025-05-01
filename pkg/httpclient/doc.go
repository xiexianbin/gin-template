package httpclient

/*
package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/xiexianbin/gin-template/pkg/httpclient"
)

func main() {
	// create client
	client := httpclient.New(
		httpclient.WithMaxRetries(5),
		httpclient.WithRetryWait(1*time.Second, 30*time.Second),
	)

	// send GET request
	resp, err := client.Get(context.Background(), "https://httpbin.org/get", http.Header{})
	if err != nil {
		fmt.Printf("Request failed: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read response body failed: %v\n", err)
		return
	}

	// print response...
	fmt.Printf("Response status: %s\n, body: %s\n", resp.Status, body)
}
*/
