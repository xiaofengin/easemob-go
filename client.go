package easemob_go

import (
	"context"
	"net/http"
	"strings"
	"time"
)

type AppToken struct {
	Application string `json:"application"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}
type ClientParam struct {
	GrantType    string `json:"grant_type"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Ttl          string `json:"ttl"`
}
type Client struct {
	http         *http.Client `json:"-"`
	baseURL      string
	appkey       string
	clientId     string
	clientSecret string
	appToken     string
}

func New(appkey, clientId, clientSecret, domainURL string) (*Client, error) {
	tr := http.DefaultTransport.(*http.Transport).Clone()
	tr.MaxIdleConnsPerHost = 5
	tr.IdleConnTimeout = 59 * time.Second
	tr.ExpectContinueTimeout = 2 * time.Second

	baseURL := domainURL + "/" + strings.Replace(appkey, "#", "/", 1)
	client := &Client{
		appkey:       appkey,
		clientId:     clientId,
		clientSecret: clientSecret,
		baseURL:      baseURL,
		http: &http.Client{
			Timeout:   6 * time.Second,
			Transport: tr,
		},
	}
	data := &ClientParam{
		GrantType:    "client_credentials",
		ClientId:     clientId,
		ClientSecret: clientSecret,
		Ttl:          "1024000",
	}
	token, err := client.createAppToken(data)
	if err == nil {
		client.appToken = token
	}
	return client, err
}
func (c *Client) createAppToken(data *ClientParam) (string, error) {
	var resp AppToken
	err := c.makeRequest(context.Background(), http.MethodPost, "token", nil, data, &resp)
	return resp.AccessToken, err
}
