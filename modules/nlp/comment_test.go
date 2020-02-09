package nlp

import (
	"github.com/rogeecn/aip"
	"os"
	"testing"
)

func TestCommentTag_Default(t *testing.T) {
	err := aip.Init(os.Getenv("AIP_APPKEY"), os.Getenv("AIP_APPSECRET"))
	if err != nil {
		t.Fatal(err)
	}

	var m CommentTag
	resp, err := m.Default("三星电脑电池不给力", COMMENT_TAG_TYPE_3C)
	if err != nil {
		t.Error(err)
	}
	t.Logf("resp: %+v", resp)
}
