package boot

import (
	"errors"
	"mermaid/global"
)

func bootCache() error {
	redisInfo := global.Conf.Redis
	if redisInfo.Host == "" {
		return errors.New("Redis host is empty")
	}
	return global.InitCache(redisInfo)
}
