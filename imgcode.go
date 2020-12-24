package vcs

import (
	"io"

	"gitlab.100bm.cn/micro-plat/vcs/vcs/modules/imgcode"
)

//GetImgCode 获取图形验证码
//i-->*context.Context/component.IContainer
//w-->ctx.Request.Http.GetResponse()并且Header().Set("Content-Type", "image/png")
//ident-->系统标识,account-->账号
func GetImgCode(i interface{}, w io.Writer, ident, account string) (err error) {

	obj, err := imgcode.NewCode()
	if err != nil {
		return err
	}
	return obj.Get(i, w, ident, account)
}

//VerifyImgCode 验证图形验证码
//i-->*context.Context/component.IContainer
//platName-->平台名,ident-->系统标识,account-->账号,code-->验证码
func VerifyImgCode(i interface{}, ident, account, code string) (err error) {

	obj, err := imgcode.NewCode()
	if err != nil {
		return err
	}
	return obj.Verify(i, ident, account, code)
}
