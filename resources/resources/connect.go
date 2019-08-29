package resources

import (
	"bytes"
	"fmt"
	"net/http"
)

type TenableIORequest struct {
	Endpoint    string
	Method      string
	Credentials struct {
		AccessKey string
		SecretKey string
	}
	Body []byte
}

func (T *TenableIORequest) Do() (*http.Response, error) {
	var (
		r   *http.Request
		err error
	)

	client := &http.Client{}

	if T.Body != nil {
		r, err = http.NewRequest(T.Method, T.Endpoint, bytes.NewBuffer(T.Body))
		if err != nil {
			return nil, fmt.Errorf("error creating a new http '%s' request to '%s' due to error: %v", T.Method, T.Endpoint, err)
		}
	} else {
		r, err = http.NewRequest(T.Method, T.Endpoint, nil)
		if err != nil {
			return nil, fmt.Errorf("error creating a new http '%s' request to '%s' due to error: %v", T.Method, T.Endpoint, err)
		}
	}

	r.Header.Add("accept", "application/json")
	r.Header.Add("content-type", "application/json")
	r.Header.Add("x-apiKeys", fmt.Sprintf("accessKey=%s;secretKey=%s", T.Credentials.AccessKey, T.Credentials.SecretKey))

	resp, err := client.Do(r)
	if err != nil {
		return nil, fmt.Errorf("error sending a new http '%s' request to '%s' due to error: %v", T.Method, T.Endpoint, err)
	}
	return resp, nil
}
