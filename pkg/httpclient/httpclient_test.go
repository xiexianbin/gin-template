package httpclient

import (
	"bytes"
	"context"
	"net/http"
	"testing"
)

func TestClient_Get(t *testing.T) {
	client := New()
	headers := http.Header{}
	headers.Add("Authorization", "Bearer token123")
	resp, err := client.Get(context.TODO(), "https://httpbin.org/get", headers)
	if err != nil {
		t.Error(err)
	}
	t.Logf("Get resp: %#v", resp)
}

func TestClient_Post(t *testing.T) {
	client := New()
	body := bytes.NewBufferString(`{"title":"test"}`)
	headers := http.Header{
		"X-Request-ID": []string{"123e4567-e89b-12d3-a456-426614174000"},
	}
	resp, err := client.Post(context.TODO(), "https://httpbin.org/post", "application/json", body, headers)
	if err != nil {
		t.Error(err)
	}
	t.Logf("Get resp: %#v", resp)
}
