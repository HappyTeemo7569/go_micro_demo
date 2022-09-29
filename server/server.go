package server

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
	"net/http"
)

func RunHttpByGin() {

	cr := consul.NewRegistry(
		registry.Addrs("121.199.53.29:8500"))

	// 使用gin作为路由
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		c.String(http.StatusOK, "user api")
	})

	httpServer := web.NewService(
		web.Name("productService"), // 当前微服务服务名
		web.Registry(cr),           // 注册到consul
		web.Address(":8081"),       // 端口
		web.Metadata(map[string]string{"protocol": "http"}), // 元信息
		web.Handler(r)) // 路由

	_ = httpServer.Run()

}

func RunHttpByMicro() {
	httpServer := web.NewService(web.Address(":8081")) // 路由
	httpServer.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello go micro"))
	})
	_ = httpServer.Run()
}
