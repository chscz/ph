package main

import (
	"payhere/internal/auth"
	"payhere/internal/config"
	"payhere/internal/handler"
	"payhere/internal/mysql"
	"payhere/internal/router"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	db, err := mysql.InitMYSQL(cfg.MySQL)
	if err != nil {
		panic(err)
	}

	ua := auth.NewUserAuth(cfg.JWT)
	uh := handler.NewUserHandler(mysql.UserRepo{DB: db}, ua)
	mh := handler.NewProductHandler(mysql.ProductRepo{DB: db})
	//uh := handler.NewUserHandler(nil)
	//mh := handler.NewProductHandler(nil)

	r := router.InitGin(uh, mh)

	r.Run() // 서버가 실행 되고 0.0.0.0:8080 에서 요청을 기다립니다.
}
