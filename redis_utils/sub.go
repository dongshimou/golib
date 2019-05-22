package redis_utils

import (
	"github.com/gomodule/redigo/redis"
)

func SubTopic(conn redis.Conn, channel string, stream chan []byte) chan error {
	pc := redis.PubSubConn{Conn: conn}
	errChan := make(chan error, 8)
	if err := pc.Subscribe(channel); err != nil {
		errChan <- err
	}
	go func() {
		for {
			if pc.Conn.Err() != nil {
				errChan <- pc.Conn.Err()
				return
			}
			switch v := pc.Receive().(type) {
			case redis.Message:
				stream <- v.Data
			case redis.Subscription:
				//do nothing
			case error:
				errChan <- v
				return
			}
		}
	}()
	return errChan
}
