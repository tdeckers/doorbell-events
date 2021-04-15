package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPost(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(EventHandler))
	defer ts.Close()

	msg := `
	{
		"message": {
			"attributes": {
				"key": "value"
			},
			"data": "ewoJCSJldmVudElkIiA6ICI3ODhlNzViZS02ZmE5LTQzNTItOGQwOC1iN2M4MjM4ODY3ZmMiLAoJCSJ0aW1lc3RhbXAiIDogIjIwMTktMDEtMDFUMDA6MDA6MDFaIiwKCQkicmVzb3VyY2VVcGRhdGUiIDogewoJCSAgIm5hbWUiIDogImVudGVycHJpc2VzL3Byb2plY3QtaWQvZGV2aWNlcy9kZXZpY2UtaWQiLAoJCSAgImV2ZW50cyIgOiB7CgkJCSJzZG0uZGV2aWNlcy5ldmVudHMuRG9vcmJlbGxDaGltZS5DaGltZSIgOiB7CgkJCQkiZXZlbnRTZXNzaW9uSWQiIDogIkNqWTVZM1ZLYVRad1IzbzRZMTlZYlRWZk1GLi4uIiwKCQkJCSJldmVudElkIiA6ICJGT1BncW5sV09QaUxQeVJIaGcxQUtFcFZZXy4uLiIKCQkJfQoJCSAgfSwKCQkgICJ1c2VySWQiOiAiQVZQSHdFdUJmblBPblRxelZGVDRJT05YMlFxaHU5RUo0dWJPLWJOblEteWkiLAoJCSAgInJlc291cmNlR3JvdXAiOiBbCgkgICJlbnRlcnByaXNlcy9wcm9qZWN0LWlkL2RldmljZXMvZGV2aWNlLWlkIgoJICBdCgkJfQp9Cg==",
			"messageId": "2070443601311540",
			"message_id": "2070443601311540",
			"publishTime": "2021-02-26T19:13:55.749Z",
			"publish_time": "2021-02-26T19:13:55.749Z"
		},
	   "subscription": "projects/myproject/subscriptions/mysubscription"
	}
	`
	// event := `{ "message": "rrr", "subscription":"SSS"}`

	reader := strings.NewReader(msg)
	res, err := http.Post(ts.URL, "application/json", reader)
	if err != nil {
		t.Fatal(err)
	}
	_, err = io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Fatalf("Unexpected status: %s", res.Status)
	}

}

func TestGet(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(EventHandler))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	response, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Fatal(err)
	}
	if string(response) != "ok" {
		t.Errorf("unexpected response: %s", response)

	}

}
