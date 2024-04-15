package httpClient

import (
	"io"
	"net/http"
	"net/http/cookiejar"
	"strings"

	jjson "github.com/chaolihf/udpgo/json"
)

type HttpClient struct {
	client http.Client
}

func NewHttpClient() (*HttpClient, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	client := http.Client{
		Jar: jar,
	}
	return &HttpClient{
		client: client,
	}, nil
}

func (httpClient *HttpClient) ExecuteTextRequest(url string, method string, content string, headers string) (string, error) {
	var payload *strings.Reader
	if len(content) > 0 {
		payload = strings.NewReader(content)
	} else {
		payload = nil
	}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return "", err
	}
	if len(headers) > 0 {
		headInfo, err := jjson.FromBytes([]byte(headers))
		if err != nil {
			return "", err
		}
		for name := range headInfo.Attributes {
			req.Header.Set(name, headInfo.GetString(name))
		}
	}
	resp, err := httpClient.client.Do(req)
	if err != nil {
		return "", nil
	}
	defer resp.Body.Close()
	datas, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", nil
	}
	return string(datas), nil
}
