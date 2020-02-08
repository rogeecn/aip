package nlp

import (
	"github.com/rogeecn/aip"
	"os"
	"testing"
)

func TestDepParser_Default(t *testing.T) {
	err := aip.Init(os.Getenv("AIP_APPKEY"), os.Getenv("AIP_APPSECRET"))
	if err != nil {
		t.Fatal(err)
	}

	var m DepParser
	resp, err := m.Default("今天天气怎么样", MODE_WEB)
	if err != nil {
		t.Error(err)
	}
	t.Logf("resp: %+v", resp)
}
