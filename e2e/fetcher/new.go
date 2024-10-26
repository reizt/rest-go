package fetcher

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type Fetcher struct {
	baseURL string
	client  *http.Client
}

func New(baseURL string) *Fetcher {
	return &Fetcher{
		baseURL: baseURL,
		client:  http.DefaultClient,
	}
}

type Request struct {
	Method  string
	Path    string
	Body    interface{}
	Headers map[string]string
}

type Response struct {
	StatusCode int
	Body       io.ReadCloser
	Headers    map[string]string
	Cookies    map[string]string
}

func (Fetcher) parseHeaders(headers http.Header) map[string]string {
	headerMap := make(map[string]string)
	for k, v := range headers {
		headerMap[k] = v[0]
	}
	return headerMap
}

func (Fetcher) parseCookies(cookies []*http.Cookie) map[string]string {
	cookieMap := make(map[string]string)
	for _, cookie := range cookies {
		cookieMap[cookie.Name] = cookie.Value
	}
	return cookieMap
}

func (f *Fetcher) Fetch(req Request) (*Response, error) {
	var bodyJson string
	if req.Body != nil {
		jsonData, err := json.Marshal(req.Body)
		if err != nil {
			return nil, err
		}
		bodyJson = string(jsonData)
	}
	httpReq, err := http.NewRequest(req.Method, f.baseURL+req.Path, bytes.NewBuffer([]byte(bodyJson)))
	if err != nil {
		return nil, err
	}
	if httpReq.Body != nil {
		httpReq.Header.Set("Content-Type", "application/json")
	}
	resp, err := f.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return &Response{
		StatusCode: resp.StatusCode,
		Body:       resp.Body,
		Headers:    f.parseHeaders(resp.Header),
		Cookies:    f.parseCookies(resp.Cookies()),
	}, nil
}
