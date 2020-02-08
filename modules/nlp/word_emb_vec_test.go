package nlp

import (
	"github.com/rogeecn/aip"
	"os"
	"testing"
)

func TestWordEmbVec_Default(t *testing.T) {
	err := aip.Init(os.Getenv("AIP_APPKEY"), os.Getenv("AIP_APPSECRET"))
	if err != nil {
		t.Fatal(err)
	}

	var m WordEmbVec
	resp, err := m.Default("张飞")
	if err != nil {
		t.Error(err)
	}
	t.Logf("resp: %+v", resp)
}
