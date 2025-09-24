package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type rdb struct {
	client *redis.Client
}


// constructor ?
// primary keys?
// methods?

func ConnectToRedis() *rdb {
	r:=redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	return &rdb{client:r}
}

func (rdb *rdb) PublishToRedis(ctx context.Context, channel string, data []byte) *redis.IntCmd{
	published:=rdb.client.Publish(ctx, channel, data)
	return published
}
func (rdb *rdb) SubscribeToRedis(ctx context.Context, channel string) *redis.PubSub{
	subscribed := rdb.client.Subscribe(ctx, channel);
	return subscribed
}

func (rdb *rdb) Set(ctx context.Context, channel string, symbol string) *redis.IntCmd{
	add := rdb.client.SAdd(ctx, channel, symbol)
	return add
}

func (rdb *rdb) UnSet(ctx context.Context, channel string, symbol string) *redis.IntCmd{
	remove := rdb.client.SRem(ctx, channel, symbol)
	return remove
}