// 新闻摘要接口
// see https://cloud.baidu.com/doc/NLP/s/vk3pmn49r#%E6%96%B0%E9%97%BB%E6%91%98%E8%A6%81%E6%8E%A5%E5%8F%A3

package nlp

import (
	"github.com/rogeecn/aip"
	"github.com/rogeecn/aip/modules"
	"github.com/rogeecn/aip/utils"
	"github.com/sirupsen/logrus"
)

const (
	news_summary = "https://aip.baidubce.com/rpc/2.0/nlp/v1/news_summary"
)

type NewsSummary struct {
}

type NewsSummaryBody struct {
	Title         string `json:"title"`
	Content       string `json:"content"`
	MaxSummaryLen int    `json:"max_summary_len"`
}

type NewsSummaryResponse struct {
	modules.BaseResp

	Summary string `json:"summary"`
}

func (m NewsSummary) Default(title, content string, maxSummaryLen int) (NewsSummaryResponse, error) {
	var resp NewsSummaryResponse

	body := utils.MustJson(NewsSummaryBody{title, content, maxSummaryLen})
	logrus.Debugf("[news_summary] %s", body)

	iresp, err := utils.CommonResponse(aip.Post(news_summary).Send(string(body)), resp)
	if err != nil {
		return resp, err
	}

	finalResp, _ := iresp.(NewsSummaryResponse)
	return finalResp, err
}
