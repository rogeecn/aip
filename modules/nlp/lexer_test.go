package nlp

import (
	"github.com/rogeecn/aip"
	"os"
	"testing"
)

func TestLexer_Default(t *testing.T) {
	err := aip.Init(os.Getenv("AIP_APPKEY"), os.Getenv("AIP_APPSECRET"))
	if err != nil {
		t.Fatal(err)
	}

	var lexer Lexer
	resp, err := lexer.Default("百度是一家高科技公司")
	if err != nil {
		t.Error(err)
	}
	t.Logf("resp: %+v", resp)

	resp, err = lexer.Custom("百度是一家高科技公司")
	if err != nil {
		t.Error(err)
	}
	t.Logf("resp: %+v", resp)
}
