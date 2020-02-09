// 对话情绪识别接口
// see https://cloud.baidu.com/doc/NLP/s/vk3pmn49r#%E5%AF%B9%E8%AF%9D%E6%83%85%E7%BB%AA%E8%AF%86%E5%88%AB%E6%8E%A5%E5%8F%A3

package nlp

import (
	"github.com/rogeecn/aip"
	"github.com/rogeecn/aip/modules"
	"github.com/rogeecn/aip/utils"
	"github.com/sirupsen/logrus"
)

const (
	emotion = "https://aip.baidubce.com/rpc/2.0/nlp/v1/emotion"
)

const (
	SCENE_TALK             = "talk"
	SCENE_TASK             = "task"
	SCENE_CUSTOMER_SERVICE = "customer_service"
)

type Emotion struct {
}

type EmotionBody struct {
	Text  string `json:"text"`
	Scene string `json:"scene"`
}

type EmotionResponse struct {
	modules.BaseResp

	Province     string `json:"province"`
	City         string `json:"city"`
	ProvinceCode string `json:"province_code"`
	LogID        int64  `json:"log_id"`
	Text         string `json:"text"`
	Town         string `json:"town"`
	Phonenum     string `json:"phonenum"`
	Detail       string `json:"detail"`
	County       string `json:"county"`
	Person       string `json:"person"`
	TownCode     string `json:"town_code"`
	CountyCode   string `json:"county_code"`
	CityCode     string `json:"city_code"`
}

func (m Emotion) Default(text, scene string) (EmotionResponse, error) {
	var resp EmotionResponse

	body := utils.MustJson(EmotionBody{text, scene})
	logrus.Debugf("[emotion] %s", body)

	iresp, err := utils.CommonResponse(aip.Post(emotion).Send(string(body)), resp)
	if err != nil {
		return resp, err
	}

	finalResp, _ := iresp.(EmotionResponse)
	return finalResp, err
}
