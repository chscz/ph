package main

import (
	"payhere/internal/config"
	"payhere/internal/handler"
	"payhere/internal/mysql"
	"payhere/internal/router"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		//todo
	}
	_ = cfg

	repo := mysql.InitDB()
	uh := handler.NewUserHandler(repo)
	mh := handler.NewMenuHandler(repo)

	r := router.InitGin(uh, mh)

	r.Run() // 서버가 실행 되고 0.0.0.0:8080 에서 요청을 기다립니다.
}
