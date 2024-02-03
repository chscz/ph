package main

import (
	"fmt"
	"log"
	"os"

	"github.com/chscz/ph/internal/auth"
	"github.com/chscz/ph/internal/config"
	"github.com/chscz/ph/internal/handler"
	"github.com/chscz/ph/internal/mysql"
	"github.com/chscz/ph/internal/router"
)

func main() {
	log.Println("start!!")
	cfg, err := config.LoadFromEnv()
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		panic(err)
	}
	log.Println(cfg)

	db, err := mysql.InitMYSQL(cfg.MySQL)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		panic(err)
	}

	ua := auth.NewUserAuth(cfg.JWT)
	uh := handler.NewUserHandler(mysql.UserRepo{DB: db}, ua, cfg.JSONRespType)
	mh := handler.NewProductHandler(mysql.ProductRepo{DB: db}, cfg.JSONRespType)

	r := router.InitGin(uh, mh, cfg.LocalDebugMode)

	if err = r.Run(); err != nil {
		fmt.Println(err)
		log.Println(err)
		os.Exit(1)
	}
}
