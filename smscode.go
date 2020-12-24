package vcs

import (
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/vcs/vcs/modules/smscode"
)

type SendRequest struct {
	ReqID        string `json:"req_id" m2s:"req_id" valid:"required"`
	Ident        string `json:"ident" m2s:"ident" valid:"required"`
	PhoneNo      string `json:"phone_no" m2s:"phone_no" valid:"required"`
	TemplateID   string `json:"template_id" m2s:"template_id" valid:"required"`
	Keywords     string `json:"keywords" m2s:"keywords" valid:"required"`
	DeliveryTime string `json:"delivery_time,omitempty" m2s:"delivery_time"` //格式：yyyy-mm-dd hh:mm:ss
}

type SendResult struct {
	RecordID string `json:"record_id"`
}

//SendSmsCode 发送短信验证码 *使用前,请先配置短信发送的rpc地址 SetConfig(WithSmsSendUrl(""))
//i-->*context.Context/component.IContainer
//req-->短信验证码获取实体
//实体req_id:请求id,ident:系统标识,phone_no:手机号,template_id:短信模板编号,keywords:发送内容,delivery_time:定时发送时间(可空),格式：yyyy-mm-dd hh:mm:ss
//返回值result {"record_id":"xxxx"}
func SendSmsCode(i interface{}, req *SendRequest) (result *SendResult, err error) {

	info := &smscode.SendRequest{
		ReqID:        req.ReqID,
		Ident:        req.Ident,
		PhoneNo:      req.PhoneNo,
		TemplateID:   req.TemplateID,
		Keywords:     req.Keywords,
		DeliveryTime: req.DeliveryTime,
	}

	obj, err := smscode.NewCode()
	if err != nil {
		return nil, err
	}
	r, err := obj.Send(i, info)
	if err != nil {
		return nil, err
	}

	result = &SendResult{}
	err = types.Map2Struct(r, result)
	if err != nil {
		return nil, err
	}

	return
}

//VerifySmsCode 验证短信验证码
//i-->*context.Context/component.IContainer
//ident-->系统标识,phone-->手机号,code-->验证码
func VerifySmsCode(i interface{}, ident, phone, code string) (err error) {

	obj, err := smscode.NewCode()
	if err != nil {
		return err
	}

	return obj.Validate(i, ident, phone, code)
}
