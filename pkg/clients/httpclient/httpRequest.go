package httpClient

import (
	"crypto/tls"
	"strings"

	jjson "github.com/chaolihf/udpgo/json"
	lang "github.com/chaolihf/udpgo/lang"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
)

type HttpClient struct {
	client *resty.Client
}

var logger *zap.Logger

func init() {
	logger = lang.InitProductLogger("logs/windows_helper.log", 300, 3, 10)
}

func NewHttpClient() (*HttpClient, error) {
	client := resty.New()
	return &HttpClient{
		client: client,
	}, nil
}

/*
*

	@param headers json格式的协议头
*/
func (httpClient *HttpClient) ExecuteTextRequest(url string, method string, content string, headers string) (string, error) {
	request := httpClient.client.R()
	if len(headers) > 0 {
		headInfo, err := jjson.FromBytes([]byte(headers))
		if err != nil {
			return "", err
		}
		for name := range headInfo.Attributes {
			request.SetHeader(name, headInfo.GetString(name))
		}
	}
	if len(content) > 0 {
		request.SetBody(content)
	}
	if strings.HasPrefix(url, "https") {
		httpClient.client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	var response *resty.Response
	var err error
	switch strings.ToLower(method) {
	case "get":
		{
			response, err = request.Get(url)
			break
		}
	case "post":
		{
			response, err = request.Post(url)
		}
	}
	if err != nil {
		return "", nil
	}
	return response.String(), nil
}
