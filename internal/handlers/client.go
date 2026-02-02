package handlers

import (
	"net/http"
)

const apiurl = "https://pokeapi.co/api/v2/"

func do(endpoint string) (res *http.Response, err error) {
	req, err := http.NewRequest("GET", apiurl+endpoint, nil)
	if err != nil {
		return nil, err
	}
	res, err = http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return res, err
}
