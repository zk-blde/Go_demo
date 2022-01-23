package main

import (
	"github.com/gin-gonic/gin"

	"casbin_demo/lib"
)

func main() {
	// /depts
	//sub:= "shenyi" // 想要访问资源的用户。
	//obj:= "/depts" // 将被访问的资源。
	//act:= "POST" // 用户对资源执行的操作。
	//e:= casbin.NewEnforcer("resources/model.conf","resources/p.csv")

	//ok:= e.Enforce(sub, obj, act)
	//if ok {
	//	log.Println("运行通过")
	//}

	r := gin.New()
	r.Use(lib.Middlewares()...)
	// r.GET("/depts", func(context *gin.Context) {
	// 	context.JSON(200, gin.H{"result": "部门列表"})
	// })
	// r.GET("/depts/:id", func(context *gin.Context) {
	// 	context.JSON(200, gin.H{"result": "部门详细"})
	// })
	// r.POST("/depts", func(context *gin.Context) {
	// 	context.JSON(200, gin.H{"result": "批量修改部门列表"})
	// })

	r.GET("/:domain/depts", func(context *gin.Context) {
		context.JSON(200, gin.H{"result": "部门列表--" + context.Param("domain")})
	})

	r.POST("/:domain/depts", func(context *gin.Context) {
		context.JSON(200, gin.H{"result": "批量修改部门列表" + context.Param("domain")})
	})
	r.Run(":8080")

}
