package common

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-plugins/config/source/consul/v2"
	"strconv"
)

func GetConsulConfig(host string, port int64, prefix string) (config.Config, error) {
	// NewSource 创建一个consul
	consulSource := consul.NewSource(
		// 127.0.0.1：8500,设置consul的地址，也就是consul部署的地址
		consul.WithAddress(host+":"+strconv.FormatInt(port, 10)),
		// 设置要使用的前缀key
		consul.WithPrefix(prefix),
		// 是否从配置项中删除前缀
		consul.StripPrefix(true),
	)
	// 返回一个新的config
	cf, err := config.NewConfig()
	if err != nil {
		return cf, err
	}
	// 加载配置源
	err = cf.Load(consulSource)
	return cf, err
}
