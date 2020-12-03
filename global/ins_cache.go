package global

import (
	redisCli "github.com/go-redis/redis/v8"
	"mermaid/model"
	"mermaid/utils"
)

func InitCache(redisInfo *redis) error {
	model.RedisConn = redisCli.NewClient(&redisCli.Options{
		Addr: redisInfo.Host,
	})
	_, err := model.RedisConn.Ping(utils.GetContext()).Result()
	if err != nil {
		return err
	}
	return nil
}
