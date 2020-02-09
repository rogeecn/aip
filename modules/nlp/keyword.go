// 文章标签接口
// see https://cloud.baidu.com/doc/NLP/s/vk3pmn49r#%E6%96%87%E7%AB%A0%E6%A0%87%E7%AD%BE%E6%8E%A5%E5%8F%A3

package nlp

import (
	"github.com/rogeecn/aip"
	"github.com/rogeecn/aip/modules"
	"github.com/rogeecn/aip/utils"
	"github.com/sirupsen/logrus"
)

const (
	keyword = "https://aip.baidubce.com/rpc/2.0/nlp/v1/keyword"
)

type Keyword struct {
}

type KeywordBody struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type KeywordResponse struct {
	modules.BaseResp

	LogID int64 `json:"log_id"`
	Items []struct {
		Score float64 `json:"score"`
		Tag   string  `json:"tag"`
	} `json:"items"`
}

func (m Keyword) Default(title, content string) (KeywordResponse, error) {
	var resp KeywordResponse

	body := utils.MustJson(KeywordBody{title, content})
	logrus.Debugf("[keyword] %s", body)

	iresp, err := utils.CommonResponse(aip.Post(keyword).Send(string(body)), resp)
	if err != nil {
		return resp, err
	}

	finalResp, _ := iresp.(KeywordResponse)
	return finalResp, err
}
