// 评论观点抽取接口（可定制）
// see https://cloud.baidu.com/doc/NLP/s/vk3pmn49r#%E8%AF%84%E8%AE%BA%E8%A7%82%E7%82%B9%E6%8A%BD%E5%8F%96%E6%8E%A5%E5%8F%A3%EF%BC%88%E5%8F%AF%E5%AE%9A%E5%88%B6%EF%BC%89

package nlp

import (
	"github.com/juju/errors"
	"github.com/rogeecn/aip"
	"github.com/rogeecn/aip/modules"
	"github.com/rogeecn/aip/utils"
	"github.com/sirupsen/logrus"
)

const (
	comment_tag        = "https://aip.baidubce.com/rpc/2.0/nlp/v2/comment_tag"
	comment_tag_custom = "https://aip.baidubce.com/rpc/2.0/nlp/v2/comment_tag_custom"
)

const (
	COMMENT_TAG_TYPE_HOTEL     = 1 + iota // 酒店
	COMMENT_TAG_TYPE_KTV                  // KTV
	COMMENT_TAG_TYPE_BEAUTY               // 丽人
	COMMENT_TAG_TYPE_FOOD                 // 美食、餐饮
	COMMENT_TAG_TYPE_TRAVEL               // 旅行
	COMMENT_TAG_TYPE_HEALTH               // 健康
	COMMENT_TAG_TYPE_EDUCATION            // 教育
	COMMENT_TAG_TYPE_BUSINESS             // 商业
	COMMENT_TAG_TYPE_BUILDING             // 房产
	COMMENT_TAG_TYPE_CAR                  // 汽车
	COMMENT_TAG_TYPE_LIFE                 // 生活
	COMMENT_TAG_TYPE_BUY                  // 购物
	COMMENT_TAG_TYPE_3C                   // 3c
)

type CommentTag struct {
}

type CommentTagBody struct {
	Text string `json:"text"`
	Type int    `json:"type"`
}

type CommentTagResponse struct {
	modules.BaseResp

	Items []struct {
		Prop      string `json:"prop"`
		Adj       string `json:"adj"`
		Sentiment int    `json:"sentiment"`
		BeginPos  int    `json:"begin_pos"`
		EndPos    int    `json:"end_pos"`
		Abstract  string `json:"abstract"`
	} `json:"items"`
}

func (m CommentTag) Default(text string, mode int) (CommentTagResponse, error) {
	return m.doRequest(comment_tag, text, mode)
}

func (m CommentTag) Custom(text string, mode int) (CommentTagResponse, error) {
	return m.doRequest(comment_tag_custom, text, mode)
}

func (m CommentTag) doRequest(url, text string, mode int) (CommentTagResponse, error) {
	var resp CommentTagResponse

	body := utils.MustJson(CommentTagBody{text, mode})
	logrus.Debugf("[comment_tag] %s", body)

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
