package nlp

import (
	"github.com/rogeecn/aip"
	"os"
	"testing"
)

func TestDnn_Default(t *testing.T) {
	err := aip.Init(os.Getenv("AIP_APPKEY"), os.Getenv("AIP_APPSECRET"))
	if err != nil {
		t.Fatal(err)
	}

	var m Dnn
	resp, err := m.Default("床前明月光")
	if err != nil {
		t.Error(err)
	}
	t.Logf("resp: %+v", resp)
}
