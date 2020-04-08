package restclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	enableMock = false
	mock       = make(map[string]*MockRequest)
)

type MockRequest struct {
	Url        string
	HttpMethod string
	Response   *http.Response
	Err        error
}

func GetMockId(httpMethod string, url string) string {
	return fmt.Sprintf("%s_%s", httpMethod, url)
}

func StartMockRequest() {
	enableMock = true
}

func FlushMock() {
	mock = make(map[string]*MockRequest)
}

func StopMockRequest() {
	enableMock = false
}

func AddMockRequest(mockRequest MockRequest) {
	mock[GetMockId(mockRequest.HttpMethod, mockRequest.Url)] = &mockRequest
}

func Post(url string, body interface{}, headers http.Header) (*http.Response, error) {
	if enableMock {
		// todo return local mock without calling any external resource
		mockRequest := mock[GetMockId(http.MethodPost, url)]
		if mockRequest == nil {
			return nil, errors.New("no mock found for given request")
		}
		return mockRequest.Response, mockRequest.Err
	}

	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
	request.Header = headers

	client := http.Client{}
	return client.Do(request)
}
