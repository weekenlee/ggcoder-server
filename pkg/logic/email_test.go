// +build ignore

package logic_test

import (
	. "github.com/polaris1119/config"
	"github.com/polaris1119/logger"

	"github.com/guanggu-coder/ggcoder-server/pkg/logic"
	"testing"
)

func TestSendAuthMail(t *testing.T) {
	logger.Init(ROOT+"/log", ConfigFile.MustValue("global", "log_level", "DEBUG"))

	err := logic.DefaultEmail.SendAuthMail("中文test", "内容test content，收到？", []string{"xuxinhua@zhimadj.com"})
	if err != nil {
		t.Error(err)
	} else {
		t.Log("successful")
	}
}
