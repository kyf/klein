package logic

import (
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/golang/protobuf/proto"
)

type OfflineStore struct {
	pool *redis.Pool
	ctx  *Server
}

var (
	redisHost, redisAuth string
)

const (
	MaxRedisConnIdle   = 1000
	MaxRedisConnActive = 0

	RedisConnIdleTimeout = time.Minute * 10

	RedisConnectTimeout = time.Second * 10
	RedisReadTimeout    = time.Second * 10
	RedisWriteTimeout   = time.Second * 10
)

func dialRedis() (redis.Conn, error) {
	c, err := redis.Dial("tcp", redisHost,
		redis.DialConnectTimeout(RedisConnectTimeout),
		redis.DialReadTimeout(RedisReadTimeout),
		redis.DialWriteTimeout(RedisWriteTimeout),
		redis.DialPassword(redisAuth))

	if err != nil {
		return nil, err
	}

	return c, nil
}

func checkValidRedisConn(c redis.Conn, t time.Time) error {
	if time.Since(t) < time.Minute {
		return nil
	}

	_, err := c.Do("PING")
	return err
}

func NewOfflineStore(ctx *Server) *OfflineStore {
	pool := &redis.Pool{
		MaxIdle:      MaxRedisConnIdle,
		Dial:         dialRedis,
		TestOnBorrow: checkValidRedisConn,
		MaxActive:    MaxRedisConnActive,
		IdleTimeout:  RedisConnIdleTimeout,
	}

	redisHost, redisAuth = ctx.conf.RedisHost, ctx.conf.RedisAuth
	return &OfflineStore{pool: pool, ctx: ctx}
}

func (this *OfflineStore) Put(key string, m *MessageRequest) error {
	c := this.pool.Get()
	defer c.Close()

	_, err := c.Do("lpush", key, proto.MarshalTextString(m))
	return err
}

func (this *OfflineStore) Get(key string) ([]*MessageRequest, error) {
	c := this.pool.Get()
	defer c.Close()

	result := make([]*MessageRequest, 0)
	list, err := redis.Strings(c.Do("lrange", key, 0, -1))
	if err != nil {
		return nil, err
	}

	for _, it := range list {
		var m MessageRequest
		err := proto.UnmarshalText(it, &m)
		if err == nil {
			result = append(result, &m)
		}
	}

	return result, nil
}

func (this *OfflineStore) Remove(key string) error {
	c := this.pool.Get()
	defer c.Close()

	_, err := c.Do("del", key)
	return err
}

func (this *OfflineStore) Stop() {
	this.pool.Close()
}
