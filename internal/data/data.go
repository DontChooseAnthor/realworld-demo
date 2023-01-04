package data

import (
	"realworld/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis"
	"github.com/google/wire"

	_ "github.com/go-sql-driver/mysql"
	// init mysql driver
	"realworld/internal/data/ent"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepo)

// Data .
type Data struct {
	client *ent.Client
	redis  *redis.Client
}

func NewClient(c *conf.Data, logger log.Logger) (*ent.Client, error) {
	log := log.NewHelper(logger)
	client, err := ent.Open(c.Database.Driver, c.Database.Source)
	if err != nil {
		log.Error("ent. client init error:", err)
		return nil, err
	}
	return client, nil
}

func NewRedis(c *conf.Data, logger log.Logger) *redis.Client {
	log := log.NewHelper(logger)
	opt := redis.Options{
		Addr: c.Redis.Addr,
		// 通过结构体找到对应方法
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
	}
	log.Info("create redis client")
	return redis.NewClient(&opt)
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	// 新建db连接
	entClient, err := NewClient(c, logger)

	if err != nil {
		return nil, nil, err
	}

	// 新建redis连接
	rc := NewRedis(c, logger)

	return &Data{
		client: entClient,
		redis:  rc,
	}, cleanup, nil
}
