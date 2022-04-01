package service

import (
	"golang.org/x/xerrors"
	"ss-quant-trade-common/model/orm"
	"ss-quant-trade-common/model/rest_model"
	"ss-quant-trade-common/utils/httputils"
)

const (
	// 查询支持的机器人策略
	taskList    string = "list_qtask_define"
	taskDetails string = "find_task_define_detail"
)

type trade struct{}

var Trade = new(trade)

func (t *trade) ListQTaskDefine(req *rest_model.UserIDRequest) ([]*orm.QTaskDefine, error) {
	resp := struct {
		rest_model.RestResult
		Data []*orm.QTaskDefine `json:"data"`
	}{}

	err := httputils.PostJSON(baseURL+taskList, req, requestTimeout, &resp)
	if err != nil {
		return nil, err
	}

	if !resp.Success {
		return nil, xerrors.New(resp.Message)
	}

	return resp.Data, nil
}

func (t *trade) QTaskDefineDetails(req *rest_model.FindTaskDefineDetailRequest) (*rest_model.TunbibaoTaskDefineDetailResponse, error) {
	resp := struct {
		rest_model.RestResult
		Data *rest_model.TunbibaoTaskDefineDetailResponse `json:"data"`
	}{}

	err := httputils.PostJSON(baseURL+taskDetails, req, requestTimeout, &resp)
	if err != nil {
		return nil, err
	}

	if !resp.Success {
		return nil, xerrors.New(resp.Message)
	}

	return resp.Data, nil
}
