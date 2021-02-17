package restclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	enabledMocks = false
	mocks        = make(map[string]*Mock)
)

//Mock struct
type Mock struct {
	URL        string
	HTTPMethod string
	Response   *http.Response
	Err        error
}

//GetMockID gets the mock id
func GetMockID(httpMethod string, url string) string {
	return fmt.Sprintf("%s_%s", httpMethod, url)
}

//StartMockups fucntion to test start
func StartMockups() {
	enabledMocks = true
}

//StopMockups fucntion to test stop
func StopMockups() {
	enabledMocks = false
}

//FlushMockups ..
func FlushMockups() {
	mocks = make(map[string]*Mock)
}

//AddMockup add mock
func AddMockup(mock Mock) {
	mocks[GetMockID(mock.HTTPMethod, mock.URL)] = &mock
}

//Post posts
func Post(url string, body interface{}, headers http.Header) (*http.Response, error) {

	if enabledMocks {
		//TODO: return local mock without calling any external resource!
		mock := mocks[GetMockID(http.MethodPost, url)]
		if mock == nil {
			return nil, errors.New("no mockup found for given request")
		}
		return mock.Response, mock.Err
	}

	jsonBytes, err := json.Marshal(body)
	if err != nil {
		//if we get an error in here, means that we dont have a valid body interface
		return nil, err
	}
	//NewReader returns a new Reader reading from jsonBytes
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
	request.Header = headers

	client := http.Client{}

	return client.Do(request)
}
