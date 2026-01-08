package handlers

import (
	"encoding/json"
	"io"
	"net/http"
)

const apiurl = "https://pokeapi.co/api/v2/"

var APIResponse struct {
	Success bool
}

func do(endpoint string, obj interface{}) error {
	req, err := http.NewRequest("GET", apiurl+endpoint, nil)
	if err != nil {
		return err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(resBody, &obj)
}
