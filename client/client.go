package client

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"strconv"
	"time"
)

const VERSION = "0.0.1"

type LearnClient struct {
	rdb     *redis.Client
	ctx     context.Context
	verbose bool
}

func NewLearnClient(addr string, port int) *LearnClient {
	options := redis.Options{
		Addr:         fmt.Sprintf("%s:%d", addr, port),
		ReadTimeout:  time.Duration(ViperGetInt64("read_timeout")) * time.Second,
		WriteTimeout: time.Duration(ViperGetInt64("write_timeout")) * time.Second,
	}

	rdb := redis.NewClient(&options)
	if rdb != nil {
		client := LearnClient{
			rdb:     rdb,
			ctx:     context.Background(),
			verbose: ViperGetBool("verbose"),
		}
		return &client
	}
	return nil
}

func (c *LearnClient) Ping() error {
	if c.verbose {
		fmt.Print("ping...")
	}
	pong, err := c.rdb.Ping(c.ctx).Result()
	if err != nil {
		return Fatal(err)
	}
	if c.verbose {
		fmt.Printf("%+v\n", pong)
	}
	return nil
}

func (c *LearnClient) Keys(pattern string) ([]string, error) {
	keys, err := c.rdb.Keys(c.ctx, pattern).Result()
	if err != nil {
		return nil, Fatal(err)
	}
	if c.verbose {
		for i, key := range keys {
			fmt.Printf("%d %s\n", i, key)
		}
	}
	return keys, nil
}

func (c *LearnClient) Classes(address string) (map[string]int64, error) {
	counts := map[string]int64{}
	classes, err := c.rdb.HGetAll(c.ctx, "RS"+address).Result()
	if err != nil {
		return nil, Fatal(err)
	}
	for key, value := range classes {
		count, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return nil, Fatal(err)
		}
		counts[key] = count
	}
	return counts, nil
}
