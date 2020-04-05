package config

import (
	"encoding/json"
	"github.com/yankooo/school-eco/be/model"
	"io/ioutil"
	"os"
)

var _G_conf *model.BookSellerConf

func GlobalConf() *model.BookSellerConf {
	return _G_conf
}

func InitConfig(filePath string) (err error) {
	var (
		file *os.File
		data []byte
	)

	if file, err = os.OpenFile(filePath, os.O_RDONLY, 0644); err != nil {
		return
	}
	if data, err = ioutil.ReadAll(file); err != nil {
		return
	}

	_G_conf = &model.BookSellerConf{}
	return  json.Unmarshal(data, _G_conf)
}
