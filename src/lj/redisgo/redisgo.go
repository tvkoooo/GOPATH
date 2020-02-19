package main

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

func main() {
	//连接到redis
	c, err := redis.Dial("tcp", "101.200.169.28:50500")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	//设置键和生命期
	_, err = c.Do("SET", "ljmykey", "superWang", "EX", "1")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
	username, err := redis.String(c.Do("GET", "ljmykey"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}
	time.Sleep(3 * time.Second)
	username, err = redis.String(c.Do("GET", "ljmykey"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}

	//删除key
	_, err = c.Do("SET", "mykeylj", "superWang")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
	username1, err := redis.String(c.Do("GET", "mykeylj"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username1)
	}
	_, err = c.Do("DEL", "mykeylj")
	if err != nil {
		fmt.Println("redis delelte failed:", err)
	}
	username1, err = redis.String(c.Do("GET", "mykeylj"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username1)
	}
	//读写json到redis
	key := "profile"
	imap := map[string]string{"username": "666", "phonenumber": "888"}
	value, _ := json.Marshal(imap)

	n, err := c.Do("SETNX", key, value)
	if err != nil {
		fmt.Println(err)
	}
	if n == int64(1) {
		fmt.Println("success")
	}

	var imapGet map[string]string

	valueGet, err := redis.Bytes(c.Do("GET", key))
	if err != nil {
		fmt.Println(err)
	}

	errShal := json.Unmarshal(valueGet, &imapGet)
	if errShal != nil {
		fmt.Println(err)
	}
	fmt.Println(imapGet["username"])
	fmt.Println(imapGet["phonenumber"])
	// 设置过期时间为20s
	nnn, _ := c.Do("EXPIRE", key, 20)
	if nnn == int64(1) {
		fmt.Println("success")
	}

	//检测值是否存在
	_, err = c.Do("SET", "mykey", "superWang")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
	is_key_exit, err := redis.Bool(c.Do("EXISTS", "mykey"))
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Printf("exists or not: %v \n", is_key_exit)
	}
	_, err = c.Do("DEL", "mykey")
	if err != nil {
		fmt.Println("redis delelte failed:", err)
	}
	//列表的操作
	_, err = c.Do("lpush", "runoobkey", "redis")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	_, err = c.Do("lpush", "runoobkey", "mongodb")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
	_, err = c.Do("lpush", "runoobkey", "mysql")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
	values, _ := redis.Values(c.Do("lrange", "runoobkey", "0", "100"))
	for _, v := range values {
		fmt.Println(string(v.([]byte)))
	}
	_, err = c.Do("DEL", "runoobkey")
	if err != nil {
		fmt.Println("redis delelte failed:", err)
	}
}
