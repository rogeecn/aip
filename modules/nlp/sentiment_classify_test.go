package nlp

import (
	"github.com/rogeecn/aip"
	"os"
	"testing"
)

func TestSentimentClassify_Default(t *testing.T) {
	err := aip.Init(os.Getenv("AIP_APPKEY"), os.Getenv("AIP_APPSECRET"))
	if err != nil {
		t.Fatal(err)
	}

	var m SentimentClassify
	resp, err := m.Default("苹果是一家伟大的公司")
	if err != nil {
		t.Error(err)
	}
	t.Logf("resp: %+v", resp)
}
