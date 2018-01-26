package Services

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"Services/Models"
	"github.com/gin-gonic/gin/binding"
	"log"
	"time"
)

//read me: https://www.jianshu.com/p/a31e4ee25305

func RouterPropertyHandler(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello %s", name)
}

func RouterPropertySecHandler(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")
	message := name + " is " + action
	c.String(http.StatusOK, message)
}

func RouterQueryStringHandler(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "Guest")
	lastname := c.Query("lastname")
	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
}

func RouterPostBodyHandler(c *gin.Context) {
	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "anonymous")

	c.JSON(http.StatusOK, gin.H{
		"status": gin.H{
			"status_code": http.StatusOK,
			"status":      "ok",
		},
		"message": message,
		"nick":    nick,})
}

func RouterPutQueryStringBodyHandler(c *gin.Context) {
	id := c.Query("id")
	page := c.DefaultQuery("page", "0")
	name := c.PostForm("name")
	message := c.PostForm("message")
	fmt.Printf("id: %s; page: %s; name: %s; message: %s \n", id, page, name, message)
	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
	})
	//安全Json 防止拦截
	c.SecureJSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
	})
}

func RouterPropertyBindHandler(c *gin.Context) {
	var user Models.User
	var err error
	ContentType := c.Request.Header.Get("Content-Type")

	switch ContentType {
	case "application/json":
		err = c.BindJSON(&user)
	case "application/x-www-form-urlencoded":
		err = c.BindWith(&user, binding.Form)
	}

	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	user.Username = user.Username + "你好"
	user.Passwd = "******"
	user.Age = user.Age + 12

	c.JSON(http.StatusOK, gin.H{
		"user":   user.Username,
		"passwd": user.Passwd,
		"age":    user.Age,
	})
}

func RouterPropertyBindSecHandler(c *gin.Context){
	var user Models.User
	err := c.Bind(&user)
	if err != nil{
		fmt.Println(err)
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"user":   user.Username,
		"passwd": user.Passwd,
		"age":    user.Age,
	})
}

func RouterPropertyBindXMLHandler(c *gin.Context){
	contentType := c.DefaultQuery("content_type", "json")
	if contentType == "json" {
		c.JSON(http.StatusOK, gin.H{
			"user":   "rsj217",
			"passwd": "123",
		})
	} else if contentType == "xml" {
		c.XML(http.StatusOK, gin.H{
			"user":   "rsj217",
			"passwd": "123",
		})
	}
}

func RouterRedictURLHandler(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently,"http://127.0.0.1:8000/redict/TargetHandler")
}

func RouterRedictTargetURLHandler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"TargetHandler": "redict/TargetHandler",
		"user":   "rsj217",
		"passwd": "123",
	})
}

func RouterGroupV1Handler(c *gin.Context){
	c.String(http.StatusOK, "v1 RouterGroupHandler")
}

func RouterGroupV2Handler(c *gin.Context){
	c.String(http.StatusOK, "v2 RouterGroupHandler")
}

func RouterGlobalMiddleWareHandler(c *gin.Context){
	request := c.MustGet("StepFilter").(string)
	req, _ := c.Get("StepFilter")
	c.JSON(http.StatusOK, gin.H{
		"middile_request": request,
		"request": req,
	})
}

func RouterSingleMiddleWareHandler(c *gin.Context){
	request := c.MustGet("StepFilter").(string)
	req, _ := c.Get("StepFilter")
	c.JSON(http.StatusOK, gin.H{
		"middile_request": request,
		"request": req,
	})
}

func RouterGroupFilterV1handler(c *gin.Context){
	request := c.MustGet("StepFilter").(string)
	req, _ := c.Get("StepFilter")
	c.JSON(http.StatusOK, gin.H{
		"middile_request": request,
		"request": req,
	})
}

func RouterSynchandler(c *gin.Context){
	time.Sleep(5 * time.Second)
	log.Println("Done! in path" + c.Request.URL.Path)
}

func RouterAsyncchandler(c *gin.Context){
	cCp := c.Copy()
	go func() {
		time.Sleep(5 * time.Second)
		log.Println("Done! in path" + cCp.Request.URL.Path)
	}()
}
