package conf

//CacheName 缓存配置名称配置名称
var CacheName = "cache"

func Config(c string) {
	CacheName = c
	return
}
