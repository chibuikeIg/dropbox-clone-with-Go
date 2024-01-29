package app

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func SendRequest(method string, url string, body io.Reader, headers http.Header) (any, int, error) {

	req, err := http.NewRequest(method, url, body)

	if err != nil {
		return nil, 0, errors.New("internal server error occured")
	}

	req.Header = headers
	// for k, v := range headers {
	// 	req.Header.Add(k, v)
	// }

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, 0, errors.New("internal server error occured")
	}

	defer resp.Body.Close()

	resp_body, err := io.ReadAll(resp.Body)

	if err != nil {

		return nil, 0, errors.New("internal server error occured")
	}

	var data any
	json.Unmarshal(resp_body, &data)

	return data, resp.StatusCode, nil
}
