package AppStart

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func GroupFilter(c *gin.Context){
	fmt.Println("before GroupFilter middleware")
	currentStep := "---GroupFilter"
	steps , exists := c.Get("StepFilter")
	fmt.Println(steps)
	if exists{
		steps = steps.(string) + currentStep
	}else{
		steps = currentStep
	}
	c.Set("StepFilter",steps)
	c.Next()
	fmt.Println("end GroupFilter middlerware")
}
