package authentication

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Authentication struct {
	config *authenticationConfig
}

type authenticationConfig  struct {
	Url string
}

const loginUrl = "https://steamcommunity.com/openid/login"

func (a *Authentication) GenLoginUrl(redirectURL string) (loginURL string) {
	params := [][]string{
		{"openid.ns" , "http://specs.openid.net/auth/2.0"},
		{"openid.mode" , "checkid_setup"},
		{"openid.return_to", redirectURL},
		{"openid.realm", redirectURL},
		{"openid.identity", "http://specs.openid.net/auth/2.0/identifier_select"},
		{"openid.claimed_id", "http://specs.openid.net/auth/2.0/identifier_select"},
	}

	loginURL = loginUrl + "?"
	for _, item  := range params  {
		appendStr := item[0] + "=" + url.QueryEscape(item[1])  + "&"
		loginURL = loginURL + appendStr
	}
	return loginURL
}
/*
将接口跳转回来的参数进行验证
signed: openid.signed
*/

func (a *Authentication) ValidateOpenID(signedParams map[string]string) (steamID string, err error ) {
	if signedParams["openid.signed"] == "" {
		return "", errors.New("signed为空，验证失败")
	}
	signedParams["openid.mode"] = "check_authentication"
	postData := url.Values{}
	for k, v := range signedParams {
		postData.Add(k,v)
	}
	//编码
	var buf io.Reader
	buf = strings.NewReader(postData.Encode())
	//todo 对http请求进行统一封装，便于多处调用 && 设置代理
	//设置代理
	//proxy, _ := url.Parse("http://127.0.0.1:1087")
	//httpTransport := &http.Transport{
	//	Proxy:         http.ProxyURL(proxy),
	//}
	httpClient := http.Client{
		//Transport:httpTransport,
		Timeout:time.Duration(5 * time.Second),
	}
	resp, err  := httpClient.Post(loginUrl,"x-www-form-urlencoded",buf)
	if err != nil {
		return "", err
	}
	fmt.Println(resp)
	return steamID, nil
}
