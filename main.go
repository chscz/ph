package main

import (
	"payhere/internal/config"
	"payhere/internal/handler"
	"payhere/internal/http"
	"payhere/internal/mysql"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		//todo
	}
	_ = cfg

	repo := mysql.InitDB()
	mh := handler.NewMenuHandler(repo)
	uh := http.NewUserHandler(repo)

	r := http.InitGin(uh, mh)

	r.Run() // 서버가 실행 되고 0.0.0.0:8080 에서 요청을 기다립니다.
}
