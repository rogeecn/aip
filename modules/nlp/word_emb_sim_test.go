package nlp

import (
	"github.com/rogeecn/aip"
	"os"
	"testing"
)

func TestWordEmbSim_Default(t *testing.T) {
	err := aip.Init(os.Getenv("AIP_APPKEY"), os.Getenv("AIP_APPSECRET"))
	if err != nil {
		t.Fatal(err)
	}

	var m WordEmbSim
	resp, err := m.Default("北京", "上海")
	if err != nil {
		t.Error(err)
	}
	t.Logf("resp: %+v", resp)
}
