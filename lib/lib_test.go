package lib

import (
	"testing"
	"time"
	"golang-CICD/config"
)
func init(){
	path := "../config/config.yml"
	config.ReadEnv(path)
	InitRedis()
	InitMysql()
}


func TestRedisClient(t *testing.T) {

	key := "BIS:ProductPreBind"
	//key := "testsp"
	rClient := RedisClient()
	val, err := rClient.LRange(key, 0, -1).Result()
	if err != nil {
		t.Fatal(err)
	}

	resultList := make([]interface{}, 0)
	resultMap := map[string]interface{}{}
	for _, v := range val {
		if _, ok := resultMap[v]; !ok {
			resultMap[v] = 1
			resultList = append(resultList, v)
		}
	}

	rClient.Del(key)
	rClient.RPush(key, resultList...)

	t.Log("ok")
}

func TestGORM(t *testing.T) {
	conn := DBConn().Debug()
	var seckillPrice int32
	conn.Table("seckill_products").Joins("JOIN product_sell_infos ON seckill_products.pid = product_sell_infos.pid").Joins("JOIN seckill_rounds ON seckill_products.round_id = seckill_rounds.id").Where("seckill_products.pid = ? AND seckill_rounds.enabled = ? AND seckill_rounds.start_time <= ? AND seckill_rounds.end_time >= ?", "WFMXAO4B6S0C", true, time.Now().Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05")).Select("seckill_products.seckill_price").Row().Scan(&seckillPrice)
	t.Log(seckillPrice)
}
