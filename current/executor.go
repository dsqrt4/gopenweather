package current

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Executor func(query url.URL) (io.Reader, error)

var DefaultExecutor Executor = func(query url.URL) (data io.Reader, err error) {
	r, err := getResponse(query)
	if err != nil {
		return nil, err
	}

	if r.StatusCode != 200 {
		return nil, fmt.Errorf("http %s", r.Status)
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(body), nil
}

func getResponse(query url.URL) (*http.Response, error) {
	request, err := http.NewRequest(http.MethodGet, query.String(), nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("accept", "application/json")

	c := http.Client{}
	return c.Do(request)
}
