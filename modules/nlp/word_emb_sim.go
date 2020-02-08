// 词义相似度接口
// see https://cloud.baidu.com/doc/NLP/s/ak3pmn40n#%E8%AF%8D%E4%B9%89%E7%9B%B8%E4%BC%BC%E5%BA%A6%E6%8E%A5%E5%8F%A3

package nlp

import (
	"github.com/rogeecn/aip"
	"github.com/rogeecn/aip/modules"
	"github.com/rogeecn/aip/utils"
	"github.com/sirupsen/logrus"
)

const (
	word_emb_sim = "https://aip.baidubce.com/rpc/2.0/nlp/v2/word_emb_sim"
)

type WordEmbSim struct {
}

type WordEmbSimBody struct {
	Word1 string `json:"word_1"`
	Word2 string `json:"word_2"`
}

type WordEmbSimResponse struct {
	modules.BaseResp

	Score float64 `json:"score"`
	Words struct {
		Word1 string `json:"word_1"`
		Word2 string `json:"word_2"`
	} `json:"words"`
}

func (m WordEmbSim) Default(first, second string) (WordEmbSimResponse, error) {
	var resp WordEmbSimResponse

	body := utils.MustJson(WordEmbSimBody{first, second})
	logrus.Debugf("[word_emb_sim] %s", body)

	iresp, err := utils.CommonResponse(aip.Post(word_emb_sim).Send(string(body)), resp)
	if err != nil {
		return resp, err
	}

	finalResp, _ := iresp.(WordEmbSimResponse)
	return finalResp, err
}
