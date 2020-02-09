// DNN语言模型接口
// see https://cloud.baidu.com/doc/NLP/s/ak3pmn40n#dnn%E8%AF%AD%E8%A8%80%E6%A8%A1%E5%9E%8B%E6%8E%A5%E5%8F%A3

package nlp

import (
	"github.com/juju/errors"
	"github.com/rogeecn/aip"
	"github.com/rogeecn/aip/modules"
	"github.com/rogeecn/aip/utils"
	"github.com/sirupsen/logrus"
)

const (
	dnnlm_cn = "https://aip.baidubce.com/rpc/2.0/nlp/v2/word_emb_sim"
)

type Dnn struct {
}

type DnnBody struct {
	Text string `json:"text"`
}

type DnnResponse struct {
	modules.BaseResp

	Text  string `json:"text"`
	Items []struct {
		Word string  `json:"word"`
		Prob float64 `json:"prob"`
	} `json:"items"`
	Ppl float64 `json:"ppl"`
}

func (m Dnn) Default(text string) (DnnResponse, error) {
	var resp DnnResponse

	body := utils.MustJson(DnnBody{text})
	logrus.Debugf("[word_emb_sim] %s", body)

	_, respBody, errs := aip.Post(dnnlm_cn).Send(string(body)).EndStruct(&resp)
	if len(errs) > 0 {
		return resp, errs[0]
	}
	logrus.Debugf("response body: %s", respBody)

	if resp.ErrorCode > 0 {
		return resp, errors.Errorf(resp.ErrorMsg)
	}

	return resp, nil
}
