package operating

import (
	"adminVideos/models"
	"adminVideos/models/createtable"
	"fmt"
)

// 创建person表 返回person_id
func CreatePerson(id int) (int,error){
	create := models.Db.Exec("insert into persons(users_id) values(?)",id)
	if err:=create.Error;err!=nil{
		fmt.Println("创建失败")
		return 0,err
	}

	var person createtable.Persons
	per := models.Db.Raw("SELECT id FROM persons WHERE users_id = ?",id).Scan(&person)
	if err:=per.Error;err!=nil{
		fmt.Println("查询失败")
		return 0,err
	}
	PersonsId := person.ID
	return PersonsId,nil
}

//创建数据
func CreateData(user createtable.Users) (int,error){
	query := models.Db.Create(&user).Where("username = ? and password = ?",user.UserName,user.PassWord).First(&user)
	if err:=query.Error; err!=nil{
		fmt.Println("创建失败",err)
		return 0,err
	}
	userId := user.ID
	PersonId,err := CreatePerson(userId)
	if PersonId == 0 || err !=nil{
		fmt.Println("获取persons_id失败")
		return 0,nil
	}
	return PersonId,nil
}
//返回id
func FindId(user createtable.Users) (int,error) {
	//根据用户名 密码查询用户 将查询到的结果封装在user结构中
	query:= models.Db.Where("username = ? and password = ?",user.UserName,user.PassWord).First(&user)
	if err:=query.Error; err!=nil{
		fmt.Println("用户名或密码有问题",err)
		return 0,err
	}
	userId := user.ID  //拿到id
	var person createtable.Persons
	per := models.Db.Raw("SELECT id FROM persons WHERE users_id = ?",userId).Scan(&person)
	if err:=per.Error;err!=nil{
		fmt.Println("查询失败")
		return 0,err
	}
	PersonsId := person.ID
	fmt.Println("用户id",PersonsId)
	return PersonsId,nil
}