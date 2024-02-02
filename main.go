package main

import (
	"log"
	"os"

	"payhere/internal/auth"
	"payhere/internal/config"
	"payhere/internal/handler"
	"payhere/internal/mysql"
	"payhere/internal/router"
)

func main() {
	cfg, err := config.LoadFromEnv()
	if err != nil {
		panic(err)
	}

	db, err := mysql.InitMYSQL(cfg.MySQL)
	if err != nil {
		panic(err)
	}

	ua := auth.NewUserAuth(cfg.JWT)
	uh := handler.NewUserHandler(mysql.UserRepo{DB: db}, ua, cfg.JSONRespType)
	mh := handler.NewProductHandler(mysql.ProductRepo{DB: db}, cfg.JSONRespType)

	r := router.InitGin(uh, mh)

	if err = r.Run(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
