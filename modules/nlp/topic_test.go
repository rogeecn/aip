package nlp

import (
	"github.com/rogeecn/aip"
	"os"
	"testing"
)

func TestTopic_Default(t *testing.T) {
	err := aip.Init(os.Getenv("AIP_APPKEY"), os.Getenv("AIP_APPSECRET"))
	if err != nil {
		t.Fatal(err)
	}

	title := "欧洲冠军联赛"
	content := "欧洲冠军联赛是欧洲足球协会联盟主办的年度足球比赛，代表欧洲俱乐部足球最高荣誉和水平，被认为是全世界最高素质、最具影响力以及最高水平的俱乐部赛事，亦是世界上奖金最高的足球赛事和体育赛事之一"

	var m Topic
	resp, err := m.Default(title, content)
	if err != nil {
		t.Error(err)
	}
	t.Logf("resp: %+v", resp)
}
