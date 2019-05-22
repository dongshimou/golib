package main

import (
	"github.com/dongshimou/golib/config"
	"github.com/dongshimou/golib/redis_utils"
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
)

var pool *redis.Pool

type Config struct {
	Address  string `json:"address"`
	Password string `json:"password"`
}
func Init()error{

	cfg:=Config{}

	if err:=config.Read(&cfg);err!=nil{
		return err
	}

	pool=&redis.Pool{
		MaxIdle:     100,
		MaxActive:   100,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp",
				cfg.Address,
				redis.DialConnectTimeout(240*time.Second),
				redis.DialReadTimeout(60*time.Second),
				redis.DialWriteTimeout(60*time.Second))
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("AUTH", cfg.Password); err != nil {
				c.Close()
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	return nil
}
func main(){

	if err:=Init();err!=nil{
		log.Println(err.Error())
		return
	}

	datas:=make(chan []byte,1024)

	go func() {
		for{
			func() {
				conn := pool.Get()
				defer conn.Close()

				errChan:=redis_utils.SubTopic(conn,"fuck",datas)
				if err:=<-errChan;err!=nil{
					log.Println(err.Error())
				}
			}()
		}
	}()

	done:=make(chan bool,1)
	go func() {
		conn:=pool.Get()
		defer conn.Close()
		for i:=0;i<100;i++{
			conn.Do("publish","fuck",i)
		}
		done<-true
	}()

	func() {
		for {
			select {
			case data := <-datas:
				log.Println(string(data))
			case <-done:
				return
			}
		}
	}()
}