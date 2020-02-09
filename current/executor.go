package current

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Executor func(query url.URL) (io.Reader, error)

var DefaultExecutor Executor = func(query url.URL) (data io.Reader, err error) {
	request, err := http.NewRequest(http.MethodGet, query.String(), nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("accept", "application/json")

	c := http.Client{}
	r, err := c.Do(request)
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

	return strings.NewReader(string(body)), nil
}
