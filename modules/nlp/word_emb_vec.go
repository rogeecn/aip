// 词向量表示接口
// see: https://cloud.baidu.com/doc/NLP/s/ak3pmn40n#%E8%AF%8D%E5%90%91%E9%87%8F%E8%A1%A8%E7%A4%BA%E6%8E%A5%E5%8F%A3

package nlp

import (
	"github.com/juju/errors"
	"github.com/rogeecn/aip"
	"github.com/rogeecn/aip/modules"
	"github.com/rogeecn/aip/utils"
	"github.com/sirupsen/logrus"
)

const (
	word_emb_vec = "https://aip.baidubce.com/rpc/2.0/nlp/v2/word_emb_vec"
)

type WordEmbVec struct {
}

type WordEmbVecBody struct {
	Word string `json:"word"`
}

type WordEmbVecResponse struct {
	modules.BaseResp

	Word string    `json:"word"`
	Vec  []float64 `json:"vec"`
}

func (m WordEmbVec) Default(word string) (WordEmbVecResponse, error) {
	var resp WordEmbVecResponse

	body := utils.MustJson(WordEmbVecBody{Word: word})
	logrus.Debugf("[word_emb_vec] %s", body)

	_, respBody, errs := aip.Post(word_emb_vec).Send(string(body)).EndStruct(&resp)
	if len(errs) > 0 {
		return resp, errs[0]
	}
	logrus.Debugf("response body: %s", respBody)

	if resp.ErrorCode > 0 {
		return resp, errors.Errorf(resp.ErrorMsg)
	}

	return resp, nil
}
