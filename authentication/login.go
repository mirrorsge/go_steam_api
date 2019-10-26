package authentication

import (
	"errors"
	"net/url"
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

func (a *Authentication) ValidateOpenID(signed string) (steamID string, err error ) {
	if signed == "" {
		return "", errors.New("signed为空，验证失败")
	}

	return steamID, nil
}
