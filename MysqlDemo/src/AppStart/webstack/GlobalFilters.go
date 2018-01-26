package AppStart

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func CAServerGlobalFilterHandler() gin.HandlerFunc{
	return func(c *gin.Context){
		fmt.Println("before CAServerGlobalFilter middleware")
		currentStep := "---CAServerGlobalFilter"
		steps , exists := c.Get("StepFilter")
		fmt.Println(steps)
		if exists{
			steps = steps.(string) + currentStep
		}else{
			steps = currentStep
		}
		c.Set("StepFilter",steps)
		c.Next()
		fmt.Println("end CAServerGlobalFilter middlerware")
	}
}

func AuthorityFilterHandler() gin.HandlerFunc{
	return func(c *gin.Context){
		fmt.Println("before AuthorityFilter middleware")
		currentStep := "---AuthorityFilter"
		steps , exists := c.Get("StepFilter")
		fmt.Println(steps)
		if exists{
			steps = steps.(string) + currentStep
		}else{
			steps = currentStep
		}
		c.Set("StepFilter",steps)
		c.Next()
		fmt.Println("end AuthorityFilter middlerware")
	}
}
