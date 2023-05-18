package databases

import (
	"fmt"
	"net"

	"github.com/go-redis/redis"
)

// Redis ...
func Redis(params Dependencies) (rdb *redis.Client) {
	rdb = redis.NewClient(&redis.Options{
		Addr: net.JoinHostPort(
			params.Config.Redis.Host, fmt.Sprint(params.Config.Redis.Port),
		),
		Password: params.Config.Redis.Password,
		DB:       params.Config.Redis.DBIndex,
	})

	pong, err := rdb.Ping().Result()
	if err != nil {
		panic(err)
	}

	params.Logger.Info("Redis ", pong)
	return
}
