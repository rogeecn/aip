// 情感倾向分析接口（可定制）
// see https://cloud.baidu.com/doc/NLP/s/vk3pmn49r#%E6%83%85%E6%84%9F%E5%80%BE%E5%90%91%E5%88%86%E6%9E%90%E6%8E%A5%E5%8F%A3%EF%BC%88%E5%8F%AF%E5%AE%9A%E5%88%B6%EF%BC%89

package nlp

import (
	"github.com/juju/errors"
	"github.com/rogeecn/aip"
	"github.com/rogeecn/aip/modules"
	"github.com/rogeecn/aip/utils"
	"github.com/sirupsen/logrus"
)

const (
	sentiment_classify        = "https://aip.baidubce.com/rpc/2.0/nlp/v2/comment_tag"
	sentiment_classify_custom = "https://aip.baidubce.com/rpc/2.0/nlp/v2/comment_tag_custom"
)

type SentimentClassify struct {
}

type SentimentClassifyBody struct {
	Text string `json:"text"`
}

type SentimentClassifyResponse struct {
	modules.BaseResp

	Text  string `json:"text"`
	Items []struct {
		Sentiment    int     `json:"sentiment"`
		Confidence   float64 `json:"confidence"`
		PositiveProb float64 `json:"positive_prob"`
		NegativeProb float64 `json:"negative_prob"`
	} `json:"items"`
}

func (m SentimentClassify) Default(text string) (SentimentClassifyResponse, error) {
	return m.doRequest(sentiment_classify, text)
}

func (m SentimentClassify) Custom(text string) (SentimentClassifyResponse, error) {
	return m.doRequest(sentiment_classify_custom, text)
}

func (m SentimentClassify) doRequest(url, text string) (SentimentClassifyResponse, error) {
	var resp SentimentClassifyResponse

	body := utils.MustJson(SentimentClassifyBody{text})
	logrus.Debugf("[sentiment_classify] %s", body)

	_, respBody, errs := aip.Post(url).Send(string(body)).EndStruct(&resp)
	if len(errs) > 0 {
		return resp, errs[0]
	}
	logrus.Debugf("response body: %s", respBody)

	if resp.ErrorCode > 0 {
		return resp, errors.Errorf(resp.ErrorMsg)
	}

	return resp, nil
}
