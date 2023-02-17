package api

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Request interface {
	Do(...string) (*http.Response, error)
}

func Post(url string, r Request) (*http.Response, error) {
	httpClient := &http.Client{}

	toJson, err := ToJson(r)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(toJson))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-sah-ws-4-call+json")
	req.Header.Set("Authorization", "X-Sah-Login")

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func PostAuth(url string, r Request, contextID string, cookie string) (*http.Response, error) {
	httpClient := &http.Client{}

	toJson, err := ToJson(r)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(toJson))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-sah-ws-4-call+json")
	req.Header.Set("Authorization", "X-Sah "+contextID)
	req.Header.Set("X-Context", contextID)
	req.Header.Set("Cookie", cookie)

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func ToJson(r Request) ([]byte, error) {
	return json.Marshal(r)
}
