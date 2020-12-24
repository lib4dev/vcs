package smscode

import (
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/vcs/vcs/modules/const/conf"
	"gitlab.100bm.cn/micro-plat/vcs/vcs/modules/sdkcache"
)

type Code struct {
	cfg   *conf.SmsCodeConf
	cache ICodeCache
}

//NewCode
func NewCode() (*Code, error) {

	cfg := conf.SmsCodeSetting
	if err := cfg.Valid(); err != nil {
		return nil, err
	}

	return &Code{
		cfg:   cfg,
		cache: NewCodeCache(),
	}, nil

}

func (s *Code) Send(i interface{}, info *SendRequest) (result types.XMap, err error) {

	//1.验证参数
	if err = info.Valid(); err != nil {
		return
	}

	//2.获取缓存数据操作对象
	sdkCache, err := sdkcache.NewSDKCache(i)
	if err != nil {
		return nil, err
	}

	//3.请求发送短信
	r, err := s.SendRequest(info, sdkCache.Container)
	if err != nil {
		return
	}

	//4 保存验证码到缓存
	platName := sdkCache.Container.GetPlatName()
	err = s.cache.Save(sdkCache.Cache, platName, info.Ident, info.PhoneNo, info.Keywords)
	if err != nil {
		return
	}

	//5 返回值
	return types.XMap{
		"record_id": r.RecordID,
	}, nil
}

func (s *Code) Validate(i interface{}, ident, phone, code string) (err error) {

	//1.获取缓存数据操作对象
	sdkCache, err := sdkcache.NewSDKCache(i)
	if err != nil {
		return err
	}

	//2 校验错误次数
	platName := sdkCache.Container.GetPlatName()
	errCount, err := s.cache.CheckErrorLimit(sdkCache.Cache, platName, ident, phone)
	if err != nil {
		return
	}

	//3. 校验验证码
	err = s.cache.Verify(sdkCache.Cache, platName, ident, phone, code, errCount)
	if err != nil {
		return
	}

	return
}
