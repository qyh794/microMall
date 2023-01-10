package common

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-plugins/config/source/consul/v2"
	"strconv"
)

func GetConsulConfig(host string, port int64, prefix string) (config.Config, error) {
	// tian jia pei zhi zhong xin
	// pei zhi zhong xin shi yong consul k/v mo shi
	newSource := consul.NewSource(
		consul.WithAddress(host+":"+strconv.FormatInt(port, 10)),
		consul.WithPrefix(prefix),
		consul.StripPrefix(true),
	)
	newConfig, err := config.NewConfig()
	if err != nil {
		return newConfig, err
	}
	err = newConfig.Load(newSource)
	return newConfig, err
}
