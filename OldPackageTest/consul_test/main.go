package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
	"net/http"
)

func Register(address string, port int, name string, tags []string, id string) error {
	cfg := api.DefaultConfig()
	cfg.Address = "192.168.28.100:8500"

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	//生成对应的检查对象
	check := &api.AgentServiceCheck{
		HTTP:                           "http://192.168.0.105:8021/health",
		Timeout:                        "5s",
		Interval:                       "5s", //5s检查一次
		DeregisterCriticalServiceAfter: "10s",
	}

	//生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = name
	registration.ID = id
	registration.Port = port
	registration.Tags = tags
	registration.Address = address
	registration.Check = check

	err = client.Agent().ServiceRegister(registration)
	//client.Agent().ServiceDeregister()
	if err != nil {
		panic(err)
	}
	return nil
}

//func AllServices() {
//	cfg := api.DefaultConfig()
//	cfg.Address = "192.168.28.100:8500"
//
//	client, err := api.NewClient(cfg)
//	if err != nil {
//		panic(err)
//	}
//
//	data, err := client.Agent().Services()
//	if err != nil {
//		panic(err)
//	}
//	for key, _ := range data {
//		fmt.Println(key)
//	}
//}
//func FilterSerivice() {
//	cfg := api.DefaultConfig()
//	cfg.Address = "192.168.1.103:8500"
//
//	client, err := api.NewClient(cfg)
//	if err != nil {
//		panic(err)
//	}
//
//	data, err := client.Agent().ServicesWithFilter(`Service == "user-web"`)
//	if err != nil {
//		panic(err)
//	}
//	for key, _ := range data {
//		fmt.Println(key)
//	}
//}

func main() {
	_ = Register("192.168.0.105", 8021, "user-web", []string{"mxshop", "bobby"}, "user-web")
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "test",
		})
	})
	err := r.Run(":8021")
	if err != nil {
		panic(err)
	}
	//AllServices()
	//FilterSerivice()
	//fmt.Println(fmt.Sprintf(`Service == "%s"`, "user-srv"))
}
