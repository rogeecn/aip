package utils

import (
	"github.com/parnurzeal/gorequest"
	"github.com/rogeecn/aip/modules"
	"github.com/sirupsen/logrus"
	"github.com/juju/errors"
)

func CommonResponse(agent *gorequest.SuperAgent, resp interface{}) (interface{}, error) {
	_, respBody, errs := agent.EndStruct(&resp)
	if len(errs) > 0 {
		return resp, errs[0]
	}
	logrus.Debugf("response body: %s", respBody)

	baseResp, _ := resp.(modules.BaseResp)
	if baseResp.ErrorCode != 0 {
		return resp, errors.Errorf("lexer response error: %s", baseResp.ErrorMsg)
	}

	return resp, nil
}
