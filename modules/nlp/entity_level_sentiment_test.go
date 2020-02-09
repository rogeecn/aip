package nlp

import (
	"github.com/rogeecn/aip"
	"os"
	"testing"
)

func TestEntityLevelSentiment_Default(t *testing.T) {
	err := aip.Init(os.Getenv("AIP_APPKEY"), os.Getenv("AIP_APPSECRET"))
	if err != nil {
		t.Fatal(err)
	}
	content := "导语：恋爱3个月就分手了？她又上节目以单身示人，公开招相亲对象。网友：我一直以为赵英俊是雪村换了个名字又再次出道了。说一个目前正在热播的综艺节目，《我家那闺女》想必很多人都知道吧？芒果台的综艺节目，邀请了吴昕、袁姗姗、傅园慧、何雯娜这4位女孩以及她们的父亲，另外还有情感观察员李维嘉、大张伟和欣然，以及飞行观察员秦海璐"
	var m EntityLevelSentiment
	resp, err := m.Default(content)
	if err != nil {
		t.Error(err)
	}
	t.Logf("resp: %+v", resp)
}
