package handlers

import (
	"encoding/json"
	"io"
	"net/http"
)

type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
}

func doRequest(endpoint string, result *APIResponse, obj interface{}) error {
	res, err := do(endpoint)
	if err != nil {
		return err
	}

	resBody, err := readResponseBody(res)
	if err != nil {
		return err
	}

	return buildAPIResponse(result, resBody, res, obj)
}

func readResponseBody(res *http.Response) ([]byte, error) {
	defer res.Body.Close()

	return io.ReadAll(res.Body)
}

func buildAPIResponse(
	result *APIResponse,
	resBody []byte,
	res *http.Response,
	obj interface{},
) error {

	result.Code = res.StatusCode
	result.Message = res.Status
	result.Success = res.StatusCode >= 200 && res.StatusCode < 300

	if obj != nil {
		if err := json.Unmarshal(resBody, obj); err != nil {
			return err
		}
		result.Data = obj
	}

	return nil
}
