package v1

import (
	"adminVideos/middleware/jwt"
	"adminVideos/models/createtable"
	"adminVideos/models/operating"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//注册
func Registered(c *gin.Context){
	//获取请求体的长度
	len := c.Request.ContentLength
	//根据长度创建字节空间
	body := make([]byte,len)
	//将请求体放到这个字节空间 这里的是前台发过来的json数据
	c.Request.Body.Read(body)

	//创建一个Users结构的实例
	var users createtable.Users
	//将json数据封装值users变量中
	json.Unmarshal(body,&users)
	//创建数据 创建成功以后，拿到id，生成token
	PersonsId,err := operating.CreateData(users)
	/*注意：这里并没有做字段验证*/
	if err != nil || PersonsId == 0{
		fmt.Errorf("查询返回id失败")
		c.Abort()
		return
	}
	fmt.Println(PersonsId)
	//b := strconv.Itoa(int(person_id))
	//fmt.Println("v2 type:", reflect.TypeOf(b)) //string

	str,err := jwt.CreateJWT(PersonsId) //返回完整token
	if err != nil {
		fmt.Errorf("失败")
		c.Abort()
	}
	fmt.Println("打印完整的token:",str) //打印token
	c.JSON(http.StatusOK, gin.H{
		"status" :200,
		"error": nil,
		"data": str,
		"persons_id":PersonsId,
	})
}



//登录
func Login(c *gin.Context){
	len := c.Request.ContentLength
	body := make([]byte,len)
	c.Request.Body.Read(body)

	var users createtable.Users
	json.Unmarshal(body,&users)  //将前台获取的数据封装到users结构中
	//查询数据库返回id  用户名 密码正确才会查到
	PersonsId,err := operating.FindId(users)
	if err != nil || PersonsId == 0 {
		fmt.Errorf("查询id失败",err)
		c.Abort()
		return
	}
	//b := strconv.Itoa(int(id))
	//生成token
	str,err := jwt.CreateJWT(PersonsId)
	if err != nil {
		fmt.Errorf("失败")
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status" :200,
		"error": nil,
		"data": str,
		"persons_id":PersonsId,
	})
}