package global

import (
	redisCli "github.com/go-redis/redis/v8"
	"mermaid/utils"
)

var redisConn *redisCli.Client

func InitCache(redisInfo *redis) error {
	redisConn = redisCli.NewClient(&redisCli.Options{
		Addr:               redisInfo.Host,
	})
	pong, err := redisConn.Ping(utils.GetContext()).Result()
	if err != nil {
		return err
	}
	Logger.Error(pong)
	return nil
}
