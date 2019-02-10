package models

import (
	"github.com/astaxie/goredis"
)

const (
	URL_QUEUE = "url_queue"
	URL_SER   = "url_set"
)

var (
	client goredis.Client
)

// 连接
func ConnectRedis() {
	client.Addr = "127.0.0.1:6379"
}

//添加到队列
func PutQueue(url string) {
	client.Lpush(URL_QUEUE, []byte(url))
}

// 取出队列数据
func PopQueue() string {
	res, err := client.Rpop(URL_QUEUE)
	if err != nil {
		panic(err)
	}
	return string(res)

}

// 添加到集合
func AddSet(url string) {
	client.Sadd(URL_SER, []byte(url))

}

// 判断是否在集合中
func IsHave(url string) bool {
	has, err := client.Sismember(URL_SER, []byte(url))
	if err != nil {
		return false
	}
	return has
}

// 队列的长度
func GetQueueLength() int {
	length, err := client.Llen(URL_QUEUE)
	if err != nil {
		return 0
	}
	return length
}
