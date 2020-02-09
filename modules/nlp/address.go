// 地址识别接口（邀测）
// see https://cloud.baidu.com/doc/NLP/s/vk3pmn49r#%E5%9C%B0%E5%9D%80%E8%AF%86%E5%88%AB%E6%8E%A5%E5%8F%A3%EF%BC%88%E9%82%80%E6%B5%8B%EF%BC%89

package nlp

import (
	"github.com/juju/errors"
	"github.com/rogeecn/aip"
	"github.com/rogeecn/aip/modules"
	"github.com/rogeecn/aip/utils"
	"github.com/sirupsen/logrus"
)

const (
	address = "https://aip.baidubce.com/rpc/2.0/nlp/v1/address"
)

type Address struct {
}

type AddressBody struct {
	Text string `json:"text"`
}

type AddressResponse struct {
	modules.BaseResp

	LogID int64 `json:"log_id"`
	Items []struct {
		Score float64 `json:"score"`
		Tag   string  `json:"tag"`
	} `json:"items"`
}

func (m Address) Default(text string) (AddressResponse, error) {
	var resp AddressResponse

	body := utils.MustJson(AddressBody{text})
	logrus.Debugf("[address] %s", body)

	_, respBody, errs := aip.Post(address).Send(string(body)).EndStruct(&resp)
	if len(errs) > 0 {
		return resp, errs[0]
	}
	logrus.Debugf("response body: %s", respBody)

	if resp.ErrorCode > 0 {
		return resp, errors.Errorf(resp.ErrorMsg)
	}

	return resp, nil
}
