package cache

import (
	"api/configs"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v7"
	"time"
)

type cache struct {
	client *redis.Client
}

type Cache interface {
	Close() error
	Set(key, value string, ttl time.Duration) error
	Get(key string) (string, error)
	LPush(key string, value ...interface{}) error
	LRange(key string, start, stop int64) ([]string, error)
	HGetAll(key string) (map[string]string, error)
	HExist(key string, field string) (bool, error)
	HSetNX(key string, field string, value interface{}) error
	HGet(key, field string) (string, error)
	HDel(key string, field ...string) error
	SAdd(key string, value ...interface{}) error
	SMembers(key string) ([]string, error)
	SRandMember(key string) (string, error)
}

// NewCache 创建cache
func NewCache() (Cache, error) {
	client, err := redisConnection()
	if err != nil {
		return nil, err
	}
	return &cache{client: client}, nil
}

func (c *cache) Close() error {
	return c.client.Close()
}

// Set 保存key
func (c *cache) Set(key, value string, ttl time.Duration) error {
	if err := c.client.Set(key, value, ttl).Err(); err != nil {
		return errors.New(fmt.Sprintf("cache set key: %s error.", key))
	}
	return nil
}

// Get 读取key
func (c *cache) Get(key string) (string, error) {
	value, err := c.client.Get(key).Result()
	if err != nil {
		return "", errors.New(fmt.Sprintf("cache get key: %s error.", key))
	}
	return value, nil
}

// LPush 存储数据到list
func (c *cache) LPush(key string, value ...interface{}) error {
	return c.client.LPush(key, value...).Err()
}

// LRange 从list中读取指定范围的数据
func (c *cache) LRange(key string, start, stop int64) ([]string, error) {
	return c.client.LRange(key, start, stop).Result()
}

// HGetAll 从hash从获取全部数据
func (c *cache) HGetAll(key string) (map[string]string, error) {
	return c.client.HGetAll(key).Result()
}

// HExist 判断key是否存在
func (c *cache) HExist(key string, field string) (bool, error) {
	return c.client.HExists(key, field).Result()
}

func (c *cache) HSetNX(key string, field string, value interface{}) error {
	return c.client.HSetNX(key, field, value).Err()
}

func (c *cache) HGet(key, field string) (string, error) {
	return c.client.HGet(key, field).Result()
}

func (c *cache) HDel(key string, field ...string) error {
	return c.client.HDel(key, field...).Err()
}

func (c *cache) SAdd(key string, value ...interface{}) error {
	return c.client.SAdd(key, value...).Err()
}

func (c *cache) SMembers(key string) ([]string, error) {
	return c.client.SMembers(key).Result()
}

func (c *cache) SRandMember(key string) (string, error) {
	return c.client.SRandMember(key).Result()
}

func redisConnection() (*redis.Client, error) {
	cfg := configs.Get().Redis
	client := redis.NewClient(&redis.Options{
		Addr:         cfg.Addr,
		Password:     cfg.Pass,
		DB:           cfg.DB,
		MaxRetries:   cfg.MaxRetries,
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
	})

	if err := client.Ping().Err(); err != nil {
		return nil, errors.New(fmt.Sprintf("ping cache: %s error", cfg.Addr))
	}
	return client, nil
}
