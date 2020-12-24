package imgcode

import (
	"fmt"
	"io"

	"gitlab.100bm.cn/micro-plat/vcs/vcs/modules/const/conf"
	"gitlab.100bm.cn/micro-plat/vcs/vcs/modules/sdkcache"
	"gitlab.100bm.cn/micro-plat/vcs/vcs/modules/util"
)

//Code
type Code struct {
	cfg   *conf.ImgCodeConf
	cache ICodeCache
}

//NewImgcode
func NewCode() (*Code, error) {

	cfg := conf.ImgCodeSetting
	if err := cfg.Valid(); err != nil {
		return nil, err
	}

	return &Code{
		cfg:   cfg,
		cache: NewCodeCache(),
	}, nil
}

//Get 生成图形验证码
func (s *Code) Get(i interface{}, w io.Writer, ident, account string) (err error) {

	//1.获取图形验证码
	codeByte, validCode := util.GetImgCode()

	//2.设置返回流
	_, err = util.NewImage(codeByte, 100, 40).WriteTo(w)
	if err != nil {
		return fmt.Errorf("设置图形验证码到返回的字节流中出错,err:%+v", err)
	}

	//3.获取缓存数据操作对象
	sdkCache, err := sdkcache.NewSDKCache(i)
	if err != nil {
		return err
	}

	//4.保存到缓存
	platName := sdkCache.Container.GetPlatName()
	err = s.cache.Save(sdkCache.Cache, platName, ident, account, validCode)
	if err != nil {
		return err
	}

	//5.账号错误次数清除
	return s.cache.ResetErrLimit(sdkCache.Cache, platName, ident, account)

}

//Verify 校验图形验证码
func (s *Code) Verify(i interface{}, ident, account, code string) (err error) {

	//1.获取缓存数据操作对象
	sdkCache, err := sdkcache.NewSDKCache(i)
	if err != nil {
		return
	}

	//2.保存code到缓存中
	platName := sdkCache.Container.GetPlatName()
	return s.cache.Verify(sdkCache.Cache, platName, ident, account, code)

}
