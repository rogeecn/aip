// 多实体情感倾向分析接口（可定制）【邀测】
// see https://cloud.baidu.com/doc/NLP/s/vk3pmn49r#%E5%A4%9A%E5%AE%9E%E4%BD%93%E6%83%85%E6%84%9F%E5%80%BE%E5%90%91%E5%88%86%E6%9E%90%E6%8E%A5%E5%8F%A3%EF%BC%88%E5%8F%AF%E5%AE%9A%E5%88%B6%EF%BC%89%E3%80%90%E9%82%80%E6%B5%8B%E3%80%91

package nlp

import (
	"github.com/rogeecn/aip"
	"github.com/rogeecn/aip/modules"
	"github.com/rogeecn/aip/utils"
	"github.com/sirupsen/logrus"
)

const (
	entity_level_sentiment             = "https://aip.baidubce.com/rpc/2.0/nlp/v1/entity_level_sentiment"
	entity_level_sentiment_add         = "https://aip.baidubce.com/rpc/2.0/nlp/v1/entity_level_sentiment/add"
	entity_level_sentiment_list        = "https://aip.baidubce.com/rpc/2.0/nlp/v1/entity_level_sentiment/list"
	entity_level_sentiment_delete      = "https://aip.baidubce.com/rpc/2.0/nlp/v1/entity_level_sentiment/delete"
	entity_level_sentiment_delete_repo = "https://aip.baidubce.com/rpc/2.0/nlp/v1/entity_level_sentiment/delete_repo"
	entity_level_sentiment_query       = "https://aip.baidubce.com/rpc/2.0/nlp/v1/entity_level_sentiment/query"
)

type EntityLevelSentiment struct {
}

type EntityLevelSentimentBody struct {
	Content string `json:"content"`
}

type EntityLevelSentimentResponse struct {
	modules.BaseResp

	Title   string `json:"title"`
	Content string `json:"content"`
	Items   []struct {
		PvalNeg   float64 `json:"pval_neg"`
		PvalPos   float64 `json:"pval_pos"`
		Sentiment int     `json:"sentiment"`
		Status    int     `json:"status"`
		Entity    string  `json:"entity"`
	} `json:"items"`
}

func (m EntityLevelSentiment) Default(content string) (EntityLevelSentimentResponse, error) {
	var resp EntityLevelSentimentResponse

	body := utils.MustJson(EntityLevelSentimentBody{content})
	logrus.Debugf("[entity_level_sentiment] %s", body)

	iresp, err := utils.CommonResponse(aip.Post(entity_level_sentiment).Send(string(body)), resp)
	if err != nil {
		return resp, err
	}

	finalResp, _ := iresp.(EntityLevelSentimentResponse)
	return finalResp, err
}

// 实体库新增接口
// see https://cloud.baidu.com/doc/NLP/s/vk3pmn49r#%E5%AE%9E%E4%BD%93%E5%BA%93%E6%96%B0%E5%A2%9E%E6%8E%A5%E5%8F%A3%E8%AF%B7%E6%B1%82%E8%AF%B4%E6%98%8E
type EntityLevelSentimentAddBody struct {
	Repository string   `json:"repository"`
	Entities   []string `json:"entities"`
}

type EntityLevelSentimentAddResponse struct {
	modules.BaseResp
}

func (m EntityLevelSentiment) Add(repository string, entities []string) (EntityLevelSentimentAddResponse, error) {
	var resp EntityLevelSentimentAddResponse

	body := utils.MustJson(EntityLevelSentimentAddBody{repository, entities})
	logrus.Debugf("[entity_level_sentiment_add] %s", body)

	iresp, err := utils.CommonResponse(aip.Post(entity_level_sentiment_add).Send(string(body)), resp)
	if err != nil {
		return resp, err
	}

	finalResp, _ := iresp.(EntityLevelSentimentAddResponse)
	return finalResp, err
}

// 实体库新增接口
// see https://cloud.baidu.com/doc/NLP/s/vk3pmn49r#%E5%AE%9E%E4%BD%93%E5%BA%93%E6%96%B0%E5%A2%9E%E6%8E%A5%E5%8F%A3%E8%AF%B7%E6%B1%82%E8%AF%B4%E6%98%8E
type EntityLevelSentimentListBody struct {
	Repository string `json:"repository"`
}

type EntityLevelSentimentListResponse struct {
	modules.BaseResp
	Repositories []string `json:"repositories"`
}

func (m EntityLevelSentiment) List(repository string) (EntityLevelSentimentListResponse, error) {
	var resp EntityLevelSentimentListResponse

	body := utils.MustJson(EntityLevelSentimentListBody{repository})
	logrus.Debugf("[entity_level_sentiment_list] %s", body)

	iresp, err := utils.CommonResponse(aip.Post(entity_level_sentiment_list).Send(string(body)), resp)
	if err != nil {
		return resp, err
	}

	finalResp, _ := iresp.(EntityLevelSentimentListResponse)
	return finalResp, err
}

// 实体库删除接口请求说明
// see https://cloud.baidu.com/doc/NLP/s/vk3pmn49r#%E5%AE%9E%E4%BD%93%E5%BA%93%E5%88%A0%E9%99%A4%E6%8E%A5%E5%8F%A3%E8%AF%B7%E6%B1%82%E8%AF%B4%E6%98%8E
type EntityLevelSentimentDeleteRepoBody struct {
	Repositories []string `json:"repositories"`
}

type EntityLevelSentimentDeleteRepoResponse struct {
	modules.BaseResp
}

func (m EntityLevelSentiment) DeleteRepo(repository []string) (EntityLevelSentimentDeleteRepoResponse, error) {
	var resp EntityLevelSentimentDeleteRepoResponse

	body := utils.MustJson(EntityLevelSentimentDeleteRepoBody{repository})
	logrus.Debugf("[entity_level_sentiment_list] %s", body)

	iresp, err := utils.CommonResponse(aip.Post(entity_level_sentiment_delete_repo).Send(string(body)), resp)
	if err != nil {
		return resp, err
	}

	finalResp, _ := iresp.(EntityLevelSentimentDeleteRepoResponse)
	return finalResp, err
}

// 实体名单查询接口请求说明
// see https://cloud.baidu.com/doc/NLP/s/vk3pmn49r#%E5%AE%9E%E4%BD%93%E5%90%8D%E5%8D%95%E6%9F%A5%E8%AF%A2%E6%8E%A5%E5%8F%A3%E8%AF%B7%E6%B1%82%E8%AF%B4%E6%98%8E
type EntityLevelSentimentQueryBody struct {
	Repository string `json:"repository"`
}

type EntityLevelSentimentQueryResponse struct {
	modules.BaseResp
	Entities []string `json:"entities"`
}

func (m EntityLevelSentiment) Query(repository string) (EntityLevelSentimentQueryResponse, error) {
	var resp EntityLevelSentimentQueryResponse

	body := utils.MustJson(EntityLevelSentimentQueryBody{repository})
	logrus.Debugf("[entity_level_sentiment_list] %s", body)

	iresp, err := utils.CommonResponse(aip.Post(entity_level_sentiment_query).Send(string(body)), resp)
	if err != nil {
		return resp, err
	}

	finalResp, _ := iresp.(EntityLevelSentimentQueryResponse)
	return finalResp, err
}

// 实体名单删除接口请求说明
// see https://cloud.baidu.com/doc/NLP/s/vk3pmn49r#%E5%AE%9E%E4%BD%93%E5%90%8D%E5%8D%95%E5%88%A0%E9%99%A4%E6%8E%A5%E5%8F%A3%E8%AF%B7%E6%B1%82%E8%AF%B4%E6%98%8E
type EntityLevelSentimentDeleteBody struct {
	Repository string   `json:"repository"`
	Entities   []string `json:"entities"`
}

type EntityLevelSentimentDeleteResponse struct {
	modules.BaseResp
	Entities []string `json:"entities"`
}

func (m EntityLevelSentiment) Delete(repository string, entities []string) (EntityLevelSentimentDeleteResponse, error) {
	var resp EntityLevelSentimentDeleteResponse

	body := utils.MustJson(EntityLevelSentimentDeleteBody{repository, entities})
	logrus.Debugf("[entity_level_sentiment_list] %s", body)

	iresp, err := utils.CommonResponse(aip.Post(entity_level_sentiment_delete).Send(string(body)), resp)
	if err != nil {
		return resp, err
	}

	finalResp, _ := iresp.(EntityLevelSentimentDeleteResponse)
	return finalResp, err
}
