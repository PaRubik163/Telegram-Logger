package main

import (
	"log"
	"github.com/PaRubik163/Telegram-Logger/internal/bot"
	"github.com/PaRubik163/Telegram-Logger/internal/config"
	"github.com/PaRubik163/Telegram-Logger/internal/server"

	"github.com/joho/godotenv"
)

func main(){
	err := godotenv.Load("./config/.env")

	if err != nil{
		log.Fatal("failed to load .env file")
	}

	conf := config.NewConfig()

	tgBot, err := bot.NewBot(conf)

	if err != nil{
		log.Fatal(err)
	}

	server := server.NewServer(tgBot)
	server.Run(conf.GRPCPort)
}