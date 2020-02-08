package aip

import (
	"github.com/juju/errors"
	"github.com/parnurzeal/gorequest"
	"net/url"
	"time"
)

const (
	DefaultGrantType = "client_credentials"
)

type Keystone struct {
	appKey    string
	appSecret string
	grantType string

	accessToken AccessTokenInfo
}

type AccessTokenInfo struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
	RefreshToken     string `json:"refresh_token"`
	ExpiresIn        int    `json:"expires_in"`
	ExpireAt         time.Time
	Scope            string `json:"scope"`
	SessionKey       string `json:"session_key"`
	AccessToken      string `json:"access_token"`
	SessionSecret    string `json:"session_secret"`
}

func initKeystone(appKey, appSecret string) *Keystone {
	return &Keystone{appKey: appKey, appSecret: appSecret}
}

func (k *Keystone) FetchAccessToken() (AccessTokenInfo, error) {
	return k.FetchAccessTokenWithGrantType(DefaultGrantType)
}

// see https://ai.baidu.com/ai-doc/REFERENCE/Ck3dwjhhu
func (k *Keystone) FetchAccessTokenWithGrantType(grantType string) (AccessTokenInfo, error) {
	k.grantType = grantType

	baseURI := "https://aip.baidubce.com/oauth/2.0/token?"
	query := url.Values{
		"grant_type":    []string{grantType},
		"client_id":     []string{k.appKey},
		"client_secret": []string{k.appSecret},
	}

	var resp AccessTokenInfo
	gorequest.New().Get(baseURI + query.Encode()).EndStruct(&resp)
	if len(resp.Error) > 0 {
		return resp, errors.Errorf("get access token failed (%s:%s)", resp.Error, resp.ErrorDescription)
	}

	resp.ExpireAt = time.Now().Add(time.Second * time.Duration(resp.ExpiresIn))
	k.accessToken = resp
	return k.accessToken, nil
}

func (k *Keystone) AccessTokenExpired() bool {
	return k.accessToken.ExpireAt.Before(time.Now())
}

func (k *Keystone) RefreshAccessToken() (AccessTokenInfo, error) {
	return k.FetchAccessTokenWithGrantType(k.grantType)
}

func (k *Keystone) GetAccessTokenString() (string, error) {
	if k.AccessTokenExpired() {
		_, err := k.RefreshAccessToken()
		if err != nil {
			return k.accessToken.AccessToken, err
		}
	}
	return k.accessToken.AccessToken, nil
}
