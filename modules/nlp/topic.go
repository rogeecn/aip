// 文章分类接口
// see https://cloud.baidu.com/doc/NLP/s/vk3pmn49r#%E6%96%87%E7%AB%A0%E5%88%86%E7%B1%BB%E6%8E%A5%E5%8F%A3

package nlp

import (
	"github.com/rogeecn/aip"
	"github.com/rogeecn/aip/modules"
	"github.com/rogeecn/aip/utils"
	"github.com/sirupsen/logrus"
)

const (
	topic = "https://aip.baidubce.com/rpc/2.0/nlp/v1/topic"
)

type Topic struct {
}

type TopicBody struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type TopicResponse struct {
	modules.BaseResp

	LogID int64 `json:"log_id"`
	Item  struct {
		Lv2TagList []struct {
			Score float64 `json:"score"`
			Tag   string  `json:"tag"`
		} `json:"lv2_tag_list"`
		Lv1TagList []struct {
			Score float64 `json:"score"`
			Tag   string  `json:"tag"`
		} `json:"lv1_tag_list"`
	} `json:"item"`
}

func (m Topic) Default(title, content string) (TopicResponse, error) {
	var resp TopicResponse

	body := utils.MustJson(TopicBody{title, content})
	logrus.Debugf("[topic] %s", body)

	iresp, err := utils.CommonResponse(aip.Post(topic).Send(string(body)), resp)
	if err != nil {
		return resp, err
	}

	finalResp, _ := iresp.(TopicResponse)
	return finalResp, err
}
