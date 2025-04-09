package main

import (
	"math/rand"
	"sync"
	"github.com/gin-gonic/gin"
)

type Storage interface {
	Shorten(url string, expSecond int64, count uint8) (string, error)
	ShortLinkInfo (sid string)(*entidy.UrlDetailInfo, error)
	UnShorten (sid string)(string, error)
}
type RedisStorage struct {
	redisClient *redis.Client
	storage *Storage
	map map[string]string
	mu sync.RWMutex
}
func (store *RedisStorage) Shorten(url string, expSeconds int, count uint8) (string, error) {
	key := generateRandomString(8)
	if store.map[key] != "" {
		if count > 5 {
			return "", errors.New("exceed max count")
		}
		return store.Shorten(url, expSeconds)
	}
	store.mu.Lock()
	defer store.mu.Unlock()

	return "", nil
}
func (store *RedisStorage) ShortLinkInfo(sid string) (*entidy.UrlDetailInfo, error) {

}
func (store *RedisStorage) UnShorten(sid string) (string, error) {

}

func generateRandomString(len uint8) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, len)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

func main() {
	router := gin.Default()
	storage := &RedisStorage{
		map: 			make(map[string]string),
		mu:  			sync.RWMutex{},
		storage: 	&Storage{}
	}

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hash": Shorten("http://www.baidu.com", 3600),
			"msg":  "pong",
		})
	})

	router.Run()
}
