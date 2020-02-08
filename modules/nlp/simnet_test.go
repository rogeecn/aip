package nlp

import (
	"github.com/rogeecn/aip"
	"os"
	"testing"
)

func TestSimnet_Default(t *testing.T) {
	err := aip.Init(os.Getenv("AIP_APPKEY"), os.Getenv("AIP_APPSECRET"))
	if err != nil {
		t.Fatal(err)
	}

	var m Simnet
	resp, err := m.Default("浙富股份", "万事通自考网")
	if err != nil {
		t.Error(err)
	}
	t.Logf("resp: %+v", resp)
}
