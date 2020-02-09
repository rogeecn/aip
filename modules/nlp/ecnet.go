// 文本纠错接口
// see https://cloud.baidu.com/doc/NLP/s/vk3pmn49r#%E6%96%87%E6%9C%AC%E7%BA%A0%E9%94%99%E6%8E%A5%E5%8F%A3

package nlp

import (
	"github.com/juju/errors"
	"github.com/rogeecn/aip"
	"github.com/rogeecn/aip/modules"
	"github.com/rogeecn/aip/utils"
	"github.com/sirupsen/logrus"
)

const (
	ecnet = "https://aip.baidubce.com/rpc/2.0/nlp/v1/ecnet"
)

type Ecnet struct {
}

type EcnetBody struct {
	Text string `json:"text"`
}

type EcnetResponse struct {
	modules.BaseResp

	LogID int64 `json:"log_id"`
	Item  struct {
		VecFragment []struct {
			OriFrag     string `json:"ori_frag"`
			BeginPos    int    `json:"begin_pos"`
			CorrectFrag string `json:"correct_frag"`
			EndPos      int    `json:"end_pos"`
		} `json:"vec_fragment"`
		Score        float64 `json:"score"`
		CorrectQuery string  `json:"correct_query"`
	} `json:"item"`
	Text string `json:"text"`
}

func (m Ecnet) Default(text string) (EcnetResponse, error) {
	var resp EcnetResponse

	body := utils.MustJson(EcnetBody{text})
	logrus.Debugf("[ecnet] %s", body)

	_, respBody, errs := aip.Post(ecnet).Send(string(body)).EndStruct(&resp)
	if len(errs) > 0 {
		return resp, errs[0]
	}
	logrus.Debugf("response body: %s", respBody)

	if resp.ErrorCode > 0 {
		return resp, errors.Errorf(resp.ErrorMsg)
	}

	return resp, nil
}
