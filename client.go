package easemob_go

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
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
type hosts struct {
	Protocol string `json:"protocol"`
	Port     string `json:"port"`
	Domain   string `json:"domain"`
	Channel  string `json:"channel"`
	Priority string `json:"priority"`
}
type rsst struct {
	Hosts []hosts `json:"hosts"`
}
type server struct {
	Rest rsst `json:"rest"`
}
type Client struct {
	http         *http.Client `json:"-"`
	baseURL      string
	appkey       string
	clientId     string
	clientSecret string
	appToken     string
}

func New(appkey, clientId, clientSecret string) (*Client, error) {
	tr := http.DefaultTransport.(*http.Transport).Clone()
	tr.MaxIdleConnsPerHost = 5
	tr.IdleConnTimeout = 59 * time.Second
	tr.ExpectContinueTimeout = 2 * time.Second
	client := &Client{
		appkey:       appkey,
		clientId:     clientId,
		clientSecret: clientSecret,
		http: &http.Client{
			Timeout:   6 * time.Second,
			Transport: tr,
		},
	}
	//match, _ := regexp.MatchString("#", appID)

	//if match { // if you are using AppKey + AppTken
	//	client.appToken = appCertificate
	//	return client, nil
	//}
	//
	//token, err := client.createAppToken(uint32(time.Now().Unix()) + 7200) // 2h
	//if err != nil {
	//	return nil, err
	//}
	//client.appToken = token
	s, err := client.getBaseURL(appkey)
	if err != nil {
		return nil, err
	}
	for _, host := range s.Rest.Hosts {
		u := host.Protocol + "://" + host.Domain + "/" + strings.Replace(appkey, "#", "/", 1)
		client.baseURL = u
		data := &ClientParam{
			GrantType:    "client_credentials",
			ClientId:     clientId,
			ClientSecret: clientSecret,
			Ttl:          "1024000",
		}
		token, err := client.createAppToken(data)
		if err == nil {
			client.appToken = token
			break
		}
	}
	return client, nil
}
func (c *Client) getBaseURL(appkey string) (*server, error) {
	var resp server
	urlStr := "https://rs.easemob.com/easemob/server.json?app_key=" + url.QueryEscape(appkey)
	req, _ := http.NewRequest("GET", urlStr, nil)
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, err
	}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, err
}
func (c *Client) createAppToken(data *ClientParam) (string, error) {
	var resp AppToken
	err := c.makeRequest(context.Background(), http.MethodPost, "token", nil, data, &resp)
	return resp.AccessToken, err
}

//
//func (c *Client) createUserToken(Uuid string, expire uint32) (string, error) {
//	return chatTokenBuilder.BuildChatUserToken(c.appID, c.appCertificate, Uuid, expire)
//}
