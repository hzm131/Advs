package redis

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var Conn redis.Conn
func init(){
	var err error
	Conn, err = redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
}

func CloseConn() {
	defer Conn.Close()
}

//写入
func SaveToDB(c redis.Conn) {
	_, err := c.Do("SET", "name", "qiuqiu", "EX", "50")
	if err != nil {
		fmt.Println("redis set failed:", err)
	} else {
		fmt.Println("save success")
	}


	/*
	批量写入
	_, err := c.Do("MSET", "name", "superWang", "SEX", "F", "EX", "50")
		if err != nil {
			fmt.Println("redis set failed:", err)
		} else {
			fmt.Println("save success")
		}

	tips:EX是这个值的过期时间
	*/
}


//读取
func ReadFromDB(c redis.Conn) {
	username, err := redis.String(c.Do("GET", "name"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}
/*
批量读取
func readFromDB(c redis.Conn) {
    username, err := redis.Strings(c.Do("MGET", "SEX", "name"))
    if err != nil {
        fmt.Println("redis get failed:", err)
    } else {
        fmt.Printf("Get mykey: %v \n", username)
    }
}
*/
}

//删除
func DelFromDB(c redis.Conn) {
	_, err := c.Do("DEL", "name", "SEX")
	if err != nil {
		fmt.Println("redis delete failed:", err)
	} else {
		fmt.Println("delete success")
	}
}


//写json
func SaveJsonDataToDB(c redis.Conn) {
	imap := map[string]string{"name": "waiwaigo", "phone": "13498739038"}
	value, _ := json.Marshal(imap) //将映射转换成json
	n, err := c.Do("SETNX", "jsonkey", value)
	if err != nil {
		fmt.Println(err)
	}
	if n == int64(1) {
		fmt.Println("success")
	}
}

//读json
func ReadJsonFromDB(c redis.Conn) {
	var imapGet map[string]string
	valueGet, err := redis.Bytes(c.Do("GET", "jsonkey"))
	if err != nil {
		fmt.Println(err)
	}

	errShal := json.Unmarshal(valueGet, &imapGet)
	if errShal != nil {
		fmt.Println(err)
	}
	fmt.Println(imapGet["name"])
	fmt.Println(imapGet["phone"])
}