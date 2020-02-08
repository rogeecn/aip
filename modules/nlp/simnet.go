// 短文本相似度接口
// see https://cloud.baidu.com/doc/NLP/s/ak3pmn40n#%E7%9F%AD%E6%96%87%E6%9C%AC%E7%9B%B8%E4%BC%BC%E5%BA%A6%E6%8E%A5%E5%8F%A3

package nlp

import (
	"github.com/rogeecn/aip"
	"github.com/rogeecn/aip/modules"
	"github.com/rogeecn/aip/utils"
	"github.com/sirupsen/logrus"
)

const (
	simnet = "https://aip.baidubce.com/rpc/2.0/nlp/v2/simnet"
)

type Simnet struct {
}

type SimnetBody struct {
	Text1 string `json:"text_1"`
	Text2 string `json:"text_2"`
}

type SimnetResponse struct {
	modules.BaseResp

	Text  string `json:"text"`
	Items []struct {
		Word string  `json:"word"`
		Prob float64 `json:"prob"`
	} `json:"items"`
	Ppl float64 `json:"ppl"`
}

func (m Simnet) Default(first, second string) (SimnetResponse, error) {
	var resp SimnetResponse

	body := utils.MustJson(SimnetBody{first, second})
	logrus.Debugf("[simnet] %s", body)

	iresp, err := utils.CommonResponse(aip.Post(simnet).Send(string(body)), resp)
	if err != nil {
		return resp, err
	}

	finalResp, _ := iresp.(SimnetResponse)
	return finalResp, err
}
