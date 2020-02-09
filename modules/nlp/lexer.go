package nlp

import (
	"github.com/juju/errors"
	"github.com/rogeecn/aip"
	"github.com/rogeecn/aip/modules"
	"github.com/rogeecn/aip/utils"
	"github.com/sirupsen/logrus"
)

const (
	lexer        = "https://aip.baidubce.com/rpc/2.0/nlp/v1/lexer"
	lexer_custom = "https://aip.baidubce.com/rpc/2.0/nlp/v1/lexer_custom"
)

type Lexer struct {
}

type LexerBody struct {
	Text string `json:"text"`
}

type LexerResponse struct {
	modules.BaseResp

	Text  string `json:"text"`
	Items []struct {
		ByteLength int           `json:"byte_length"`
		ByteOffset int           `json:"byte_offset"`
		Formal     string        `json:"formal"`
		Item       string        `json:"item"`
		Ne         string        `json:"ne"`
		Pos        string        `json:"pos"`
		URI        string        `json:"uri"`
		LocDetails []interface{} `json:"loc_details"`
		BasicWords []string      `json:"basic_words"`
	} `json:"items"`
}

// see https://cloud.baidu.com/doc/NLP/s/ak3pmn40n#%E8%AF%8D%E6%B3%95%E5%88%86%E6%9E%90%E6%8E%A5%E5%8F%A3
func (m Lexer) Default(text string) (LexerResponse, error) {
	return m.doRequest(lexer, text)
}

// see https://cloud.baidu.com/doc/NLP/s/ak3pmn40n#%E8%AF%8D%E6%B3%95%E5%88%86%E6%9E%90%E6%8E%A5%E5%8F%A3
func (m Lexer) Custom(text string) (LexerResponse, error) {
	return m.doRequest(lexer_custom, text)
}

func (m Lexer) doRequest(url, text string) (LexerResponse, error) {
	var resp LexerResponse

	body := utils.MustJson(LexerBody{Text: text})
	logrus.Debugf("[lexer] %s", body)

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
