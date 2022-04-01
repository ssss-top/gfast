package service

import (
	"encoding/json"
	"github.com/gogf/gf/frame/g"
)

var (
	baseURL        = ""
	requestTimeout = 0
)

func init() {
	requestTimeout = g.Cfg().GetInt("quantify.requestTimeout")
	baseURL = g.Cfg().GetString("quantify.tradeQTaskWorkerUrl")
}

func UnmarshalRespData(in interface{}, out interface{}) error {
	data, err := json.Marshal(in)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, out); err != nil {
		return err
	}

	return nil
}
