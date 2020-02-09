package nlp

import (
	"github.com/rogeecn/aip"
	"os"
	"testing"
)

func TestAddress_Default(t *testing.T) {
	err := aip.Init(os.Getenv("AIP_APPKEY"), os.Getenv("AIP_APPSECRET"))
	if err != nil {
		t.Fatal(err)
	}

	var m Address
	resp, err := m.Default("上海市浦东新区纳贤路701号百度上海研发中心 F4A000 张三")
	if err != nil {
		t.Error(err)
	}
	t.Logf("resp: %+v", resp)
}
