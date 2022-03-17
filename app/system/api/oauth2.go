package api

import (
	"encoding/json"
	"fmt"
	"gfast/app/system/model"
	"gfast/app/system/service"
	"gfast/library"
	"io/ioutil"

	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"
	"golang.org/x/xerrors"
)

//后台用户Oauth登陆验证
func (c *auth) loginForOauth(r *ghttp.Request) (string, interface{}) {
	var apiReq *model.LoginForOauthReq
	if err := r.Parse(&apiReq); err != nil {
		c.FailJsonExit(r, err.(gvalid.Error).Current().Error())
	}

	oauthUser, err := userProfileFromOauthSrv(apiReq.AccessToken)
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}

	user, err := service.SysUser.GetUserInfoByEmail(oauthUser.Email)
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}

	if user == nil {
		if err := service.SysUser.AddUserFromOauthUser(oauthUser); err != nil {
			c.FailJsonExit(r, err.Error())
		}

		user, err = service.SysUser.GetUserInfoByEmail(oauthUser.Email)
		if err != nil {
			c.FailJsonExit(r, err.Error())
		}
	}

	ip := library.GetClientIp(r)
	userAgent := r.Header.Get("User-Agent")
	//更新用户登录记录 写入日志信息
	service.SysUser.UpdateLoginInfo(user.Id, user.UserName, ip, userAgent, "登录成功", "系统后台")
	var keys string
	if MultiLogin {
		keys = gconv.String(user.Id) + "-" + gmd5.MustEncryptString(user.UserEmail) + gmd5.MustEncryptString(user.UserPassword+ip)
	} else {
		keys = gconv.String(user.Id) + "-" + gmd5.MustEncryptString(user.UserEmail) + gmd5.MustEncryptString(user.UserPassword)
	}

	return keys, user
}

func userProfileFromOauthSrv(accessToken string) (*model.OauthUser, error) {
	jsonData := []byte(`{"operationName":null,"variables":{},"query":"{\n  profile {\n    id\n    email\n    given_name\n    family_name\n    picture\n  }\n}\n"}`)

	gc := ghttp.NewClient().SetHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	resp, err := gc.Post("http://localhost:8080/graphql", jsonData)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var res model.GraphqlQueryUserProfileRes
	if err := json.Unmarshal(bytes, &res); err != nil {
		return nil, err
	}

	if res.Message == "unauthorized" || res.Data.OauthUser == nil {
		return nil, xerrors.New(res.Message)
	}

	return res.Data.OauthUser, nil
}
