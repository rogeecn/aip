package nlp

import (
	"github.com/rogeecn/aip"
	"os"
	"testing"
)

func TestEmotion_Default(t *testing.T) {
	err := aip.Init(os.Getenv("AIP_APPKEY"), os.Getenv("AIP_APPSECRET"))
	if err != nil {
		t.Fatal(err)
	}

	var m Emotion
	resp, err := m.Default("本来今天高高兴兴", SCENE_TALK)
	if err != nil {
		t.Error(err)
	}
	t.Logf("resp: %+v", resp)
}
