package AppStart

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func SignalFilter(c *gin.Context){
	fmt.Println("before SignalFilter middleware")
	currentStep := "---SignalFilter"
	steps , exists := c.Get("StepFilter")
	fmt.Println(steps)
	if exists{
		steps = steps.(string) + currentStep
	}else{
		steps = currentStep
	}
	c.Set("StepFilter",steps)
	c.Next()
	fmt.Println("end SignalFilter middlerware")
}
