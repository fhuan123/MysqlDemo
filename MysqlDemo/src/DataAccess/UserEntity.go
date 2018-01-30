package DataAccess

import (
	"time"
	"github.com/xormplus/xorm"
	"log"
	"fmt"
	"github.com/xormplus/core"
	"strconv"
)

type User struct {
	Id         int64
	Name       string    `xorm:"varchar(25) notnull unique 'user_name'"`
	UserId     int       `xorm:"int(11) notnull 'user_id'"`
	CreateDate time.Time `xorm:"created notnull 'create_date'"`
	UpdateDate time.Time `xorm:"updated notnull 'update_date'"`
	IsDeleted  string    `xorm:"varchar(10) notnull default '0' 'is_deleted'"`
}

var engine *xorm.Engine
var err error

func TestUser(){
	var driverName string = "mysql"
	var dataSourceName string = "saasoadbdev:CpwxbMszy3Qc@tcp(rm-2ze9a47310707n81iho.mysql.rds.aliyuncs.com:3306)/test?charset=utf8"

	engine, err = xorm.NewMySQL(driverName,dataSourceName)//.NewEngine(driverName, dataSourceName)
	if err != nil {
		log.Println(err)
	}
	engine.Ping()
	engine.ShowSQL(true)
	log.Println("db sync!")
	fmt.Println(engine.DBMetas())
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "go_")
	engine.SetTableMapper(tbMapper)

	//CreateTable()

	//AddObject()
	//AddObjectList()
	//QueryObject("冯焕")
	//UpdateObjects(24,"fenghuan1000")
	//DeleteObejct(30, "fenghuan1000")
}

func CreateTable() {
	var boolvar bool = false
	boolvar, err = engine.IsTableExist(&User{})
	if err != nil {
		fmt.Println()
	} else {
		fmt.Println(boolvar)
		return
	}

	err = engine.CreateTables(&User{})
	if err != nil {
		fmt.Println()
	}
}

func AddObject() {
	user := new(User)
	user.Name = "冯焕"
	user.UserId = 100
	user.CreateDate = time.Now()
	user.UpdateDate = time.Now()
	affected, err := engine.Insert(user)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(affected)
	}
	fmt.Println(user.Id)
}

func AddObjectList() {
	users := make([]*User, 10)
	sum := 0
	for ; sum < 10; sum++ {
		users[sum] = new(User)
		users[sum].Name = "冯焕" + strconv.Itoa(sum)
		users[sum].UserId = 100 + sum
		//users[sum].CreateDate = time.Now()
		users[sum].UpdateDate = time.Now()
	}

	affected, err := engine.Insert(&users)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("count(%d) success!!!\n", affected)
	}

	for index, value := range users {
		fmt.Printf("arr[%d]=User.id(PK)=%d \n", index, value.Id)
	}




}

func QueryObject(userName string) {
	//主键id查询
	fmt.Println("主键id查询")
	user := new(User)
	engine.Id(23).Get(user)
	fmt.Printf("query1: id[%d]=%+v\n", 23, user)

	//其他字段条件查询
	fmt.Println("其他字段条件查询")
	user1 := new(User)
	user1.Name = userName
	engine.Get(user1)
	fmt.Printf("query2: name[%s]=%+v\n", userName, user1)

	//其他字段条件查询（limit1）
	fmt.Println("其他字段条件查询（limit1）")
	user2 := new(User)
	engine.Where("user_name=?", userName).Get(user2)
	fmt.Printf("query3: name[%s]=%+v\n", userName, user2)

	//其他字段条件查询（集合）
	useritems := make([]User, 0)
	engine.Where("user_name=?", userName).Find(&useritems)
	fmt.Println("其他字段条件查询(集合)")
	for index, useritem := range useritems {
		fmt.Printf("query4[%d]: name[%s]=%+v\n", index, userName, useritem)
	}

	//其他字段条件模糊查询返回记录条数
	fmt.Println("其他字段条件模糊查询返回记录条数")
	user3 := new(User)
	count, err := engine.Where("user_name like ? ", "%"+userName+"%").Count(user3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("query5: name like %s=%d\n", userName, count)

}

func UpdateObjects(id int64, userName string) {
	//根据主键ID进行修改（构造体指针修改）
	fmt.Println("根据主键ID进行修改（构造体指针修改）")
	user := new(User)
	user.Name = "fenghuan1000"
	affected, err := engine.Id(id).Update(user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("update: id= %d update %s, count(%d)\n", id, userName, affected)

	//根据条件进行修改（构造体指针修改）
	fmt.Println("根据条件进行修改多条记录（构造体指针修改）")
	user2 := new(User)
	user2.IsDeleted = "1"
	affected2, err := engine.Where("user_name=?", userName).Update(user2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("update: id= %d update %s, update items count(%d)\n", id, userName, affected2)
	//根据userName进行修改（Map修改）
}

func DeleteObejct(id int64, userName string) {
	//根据主键ID进行修改（构造体指针修改）
	fmt.Println("根据主键ID进行删除（构造体指针修改）")
	user := new(User)
	user.Id = id
	affected, err := engine.Delete(user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("update: id= %d update %s, count(%d)\n", id, userName, affected)

	//根据条件进行修改（构造体指针修改）
	fmt.Println("根据条件进行修改多条记录（构造体指针修改）")
	user2 := new(User)
	affected2, err := engine.Where("user_name=?", userName).Delete(user2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("update: id= %d update %s, update items count(%d)\n", id, userName, affected2)
}

