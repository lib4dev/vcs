package sdkcache

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/cache"
	"gitlab.100bm.cn/micro-plat/vcs/vcs/modules/const/conf"
)

//SDKCache sdk缓存操作对象
type SDKCache struct {
	Cache     cache.ICache
	Container context.IContainer
}

//NewSDKCache 初始化对象
func NewSDKCache(c interface{}) (*SDKCache, error) {
	switch v := c.(type) {
	case *context.Context:
		i := v.GetContainer()
		cache, err := i.GetCache(conf.CacheName)
		return &SDKCache{Cache: cache, Container: i}, err
	case component.IContainer:
		cache, err := v.GetCache(conf.CacheName)
		return &SDKCache{Cache: cache, Container: v}, err
	default:
		return nil, fmt.Errorf("不支持的参数类型")
	}
}
