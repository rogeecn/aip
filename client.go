package aip

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"github.com/sirupsen/logrus"
	"net/url"
)

var keystone *Keystone

func Init(appKey, appSecret string) error {
	logrus.SetLevel(logrus.DebugLevel)
	if keystone != nil {
		return nil
	}

	keystone = initKeystone(appKey, appSecret)
	_, err := keystone.FetchAccessToken()
	return err
}

func buildInterface(baseInterface string) (string, error) {
	accessToken, err := keystone.GetAccessTokenString()
	if err != nil {
		return "", err
	}
	query := url.Values{
		"access_token": []string{accessToken},
		"charset":      []string{"UTF-8"},
	}

	fullInterface := fmt.Sprintf("%s?%s", baseInterface, query.Encode())
	logrus.Debugf("request full interface: %s", fullInterface)
	return fullInterface, nil
}

func Post(baseURI string) *gorequest.SuperAgent {
	uri, err := buildInterface(baseURI)
	if err != nil {
		logrus.Errorf("prepare url failed: %+v", err)
	}
	return gorequest.New().Post(uri)
}

func Get(baseURI string) *gorequest.SuperAgent {
	uri, err := buildInterface(baseURI)
	if err != nil {
		logrus.Errorf("prepare url failed: %+v", err)
	}

	return gorequest.New().Get(uri)
}
