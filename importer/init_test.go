package importer

import (
	"os"
	"testing"

	"github.com/yaoapp/xiang/config"
	"github.com/yaoapp/xiang/model"
	"github.com/yaoapp/xiang/query"
	"github.com/yaoapp/xiang/script"
	"github.com/yaoapp/xiang/share"
)

func TestMain(m *testing.M) {
	share.DBConnect(config.Conf.DB)
	model.Load(config.Conf)
	query.Load(config.Conf)
	share.Load(config.Conf)
	script.Load(config.Conf)
	Load(config.Conf)
	code := m.Run()
	os.Exit(code)
}
