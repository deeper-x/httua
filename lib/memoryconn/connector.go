package memoryconn

import (
	"fmt"

	"github.com/go-redis/redis"
)

// Simple wrapper on Redis Client (connection, read and write)
// Usage example:
// conn := Connector{client: NewClient()}
// conn.SetValue("test", "value")
// conn.GetValue("test") // result: value

// Connector todo doc
type Connector struct {
	Client *redis.Client
}

// NewClient todo doc
func NewClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return client
}

// Ping return pong on success
func (c Connector) Ping() string {
	pong, err := c.Client.Ping().Result()
	if err != nil {
		panic(err)
	}

	return pong
}

// SetValue given an input key, set a value
func (c Connector) SetValue(key string, value interface{}) {
	err := c.Client.Set(key, value, 0).Err()
	if err != nil {
		panic(err)
	}
}

// GetValue given a key, get a value
func (c Connector) GetValue(key string) string {
	val, err := c.Client.Get(key).Result()
	if err == redis.Nil {
		return fmt.Sprintf("%s does not exist", key)
	} else if err != nil {
		panic(err)
	} else {
		return val
	}
}
