package service

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	model2 "test3/src/model"

	"github.com/go-redis/redis"
)

var Client = GetRedisClient()
var tm = 24 * time.Hour

// GetRedisClient 和redis客户端建立连接
func GetRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println(pong, err)
	}
	return client
}

// CreateCode 创建礼品码
func CreateCode(GiftCodeInfo model2.GiftCodeInfo) string {
	code := Code()
	GiftCodeInfo.GiftCode = code
	text, _ := json.Marshal(&GiftCodeInfo)
	Client.Set(code, text, tm)
	//fmt.Println(code)
	return code
}

// Code 生成8位礼品码
func Code() string {
	rand.Seed(time.Now().UnixNano())
	var letters = []byte("0123456789abcdefghijklmnopqrstuvwxyzABDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]byte, 8)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	GiftCode := string(b)
	return GiftCode
}

// Inquire 查询礼品码信息
func Inquire(GiftCode string) (GiftCodeInfo model2.GiftCodeInfo) {
	value, _ := Client.Get(GiftCode).Result()
	if value == "" {
		fmt.Println("礼品码不存在")
		return
	}

	errJson := json.Unmarshal([]byte(value), &GiftCodeInfo)
	if errJson != nil {
		fmt.Println("Unmarshal Error : ", errJson)
	}
	return GiftCodeInfo
}

// Verify 验证用户输入的礼品码
func Verify(ClientName string, GiftCode string) (GiftInfo model2.GiftInfo) {
	value, _ := Client.Get(GiftCode).Result()

	if value == "" {
		fmt.Println("礼品码不存在")
		return
	}
	var GiftCodeInfo = model2.GiftCodeInfo{}
	errJson := json.Unmarshal([]byte(value), &GiftCodeInfo)
	if errJson != nil {
		fmt.Println("Unmarshal Error")
	}
	if GiftCodeInfo.Times == 0 {
		fmt.Println("已被领取完")
		return
	}

	GiftCodeInfo.Times--
	GiftCodeInfo.ListReceived.UsersReceived += ClientName + ";"
	GiftCodeInfo.ListReceived.DateReceived += "----" + time.Now().String() + ";"

	text, _ := json.Marshal(&GiftCodeInfo)
	Client.Set(GiftCode, text, tm)
	return GiftCodeInfo.GiftText

}
