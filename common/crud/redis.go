package crud

import (
	"Oenone/common/base"
	"github.com/go-redis/redis"
)

func ZSetIncrBy(key string, member string, score float64) error {
	rdb := base.GLOBAL_RESOURCE[base.RedisClient].(*redis.Client)
	return rdb.ZIncrBy(key, score, member).Err()
}

func DelZSetMember(key string, member ...string) error {
	rdb := base.GLOBAL_RESOURCE[base.RedisClient].(*redis.Client)
	return rdb.ZRem(key, member).Err()
}
