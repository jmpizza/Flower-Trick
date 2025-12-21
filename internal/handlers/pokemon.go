package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const apiurl = "https://pokeapi.co/api/v2/"

func Handler() {
	req, err := http.NewRequest("GET", apiurl+"ability/1", nil)
	if err != nil {
		log.Fatalf("client: error %v", err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("client: error %v", err)
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("client: error %v", err)
	}
	fmt.Println("client: response body: \n", string(resBody))

	fmt.Println(res.StatusCode)
	fmt.Println(`
	
	POKEMON API
	
	`)

}
