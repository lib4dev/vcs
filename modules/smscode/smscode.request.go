package smscode

import (
	"encoding/json"
	"fmt"

	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/vcs/vcs/modules/const/conf"
)

func (s *Code) SendRequest(info *SendRequest, c context.IContainer) (r *SendResult, err error) {

	requestData, err := types.Struct2Map(info)
	if err != nil {
		return
	}

	url := conf.SmsCodeSetting.SmsCodeSendRequestURL
	status, resultVal, _, err := c.Request(url, "POST", nil, requestData, true)
	if err != nil || status != 200 {
		return nil, fmt.Errorf("发送短信请求错误,status:%d,url:%s,params:%+v,err:%+v", status, url, requestData, err)
	}

	r = &SendResult{}

	if err = json.Unmarshal([]byte(resultVal), r); err != nil {
		return nil, fmt.Errorf("发送短信返回错误,resultVal:%s,err:%+v", resultVal, err)
	}

	return
}
