package public

import (
    "Oenone/common/base"
    "Oenone/model"
    "github.com/go-redis/redis"
    "github.com/thoas/go-funk"
    "math"
)

func GetHotHouse() ([]*model.HouseTrend, error) {
    rdb := base.GLOBAL_RESOURCE[base.RedisClient].(*redis.Client)
    idArr, err := rdb.ZRevRangeWithScores(base.GetHouseQualifyZSetKey(), 0, 9).Result()
    if err != nil {
        return nil, err
    }

    n10 := math.Pow10(2)
    res := funk.Map(idArr, func(z redis.Z) *model.HouseTrend {
        house, _ := HouseDetail(z.Member.(string), model.Online)
        return model.NewHouseTrend(house, math.Trunc((z.Score+0.5/n10)*n10)/n10)
    }).([]*model.HouseTrend)
    return res, nil
}
