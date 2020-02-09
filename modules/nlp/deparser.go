//依存句法分析接口
// see https://cloud.baidu.com/doc/NLP/s/ak3pmn40n#%E4%BE%9D%E5%AD%98%E5%8F%A5%E6%B3%95%E5%88%86%E6%9E%90%E6%8E%A5%E5%8F%A3

package nlp

import (
	"github.com/juju/errors"
	"github.com/rogeecn/aip"
	"github.com/rogeecn/aip/modules"
	"github.com/rogeecn/aip/utils"
	"github.com/sirupsen/logrus"
)

const (
	depparser = "https://aip.baidubce.com/rpc/2.0/nlp/v1/depparser"
)

const (
	MODE_WEB   = 0
	MODE_QUERY = 1
)

type DepParser struct {
}

type DepParserBody struct {
	Text string `json:"text"`
	Mode int    `json:"mode"`
}

type DepParserResponse struct {
	modules.BaseResp

	Text  string `json:"text"`
	Items []struct {
		ID     string `json:"id"`
		Word   string `json:"word"`
		Postag string `json:"postag"`
		Head   string `json:"head"`
		Deprel string `json:"deprel"`
	} `json:"items"`
}

func (m DepParser) Default(text string, mode int) (DepParserResponse, error) {
	var resp DepParserResponse

	body := utils.MustJson(DepParserBody{text, mode})
	logrus.Debugf("[depparser] %s", body)

	_, respBody, errs := aip.Post(depparser).Send(string(body)).EndStruct(&resp)
	if len(errs) > 0 {
		return resp, errs[0]
	}
	logrus.Debugf("response body: %s", respBody)

	if resp.ErrorCode > 0 {
		return resp, errors.Errorf(resp.ErrorMsg)
	}

	return resp, nil
}
