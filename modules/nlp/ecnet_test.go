package nlp

import (
	"github.com/rogeecn/aip"
	"os"
	"testing"
)

func TestEcnet_Default(t *testing.T) {
	err := aip.Init(os.Getenv("AIP_APPKEY"), os.Getenv("AIP_APPSECRET"))
	if err != nil {
		t.Fatal(err)
	}

	text := "百度是一家人工只能公司"

	var m Ecnet
	resp, err := m.Default(text)
	if err != nil {
		t.Error(err)
	}
	t.Logf("resp: %+v", resp)
}
