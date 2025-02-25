package slack

import (
	"errors"
	"fmt"
	"github.com/go-zoox/fetch"
)

func SendWebhook(url string, payload *Payload) error {
	resp, err := fetch.Post(url, &fetch.Config{Body: payload, Headers: fetch.Headers{"Content-Type": "application/json"}})
	if err != nil {
		return err
	}
	if !resp.Ok() {
		return errors.New(fmt.Sprint("slack response error, status is not ok", resp.StatusCode()))
	}
	return nil
}
