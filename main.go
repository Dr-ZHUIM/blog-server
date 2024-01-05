package main

import (
	articleServer "article-server/internal/app"
	CustomMysql "article-server/internal/app/db"
	Log "article-server/internal/app/log"
)

func main() {
	Log.InitLogger()
	CustomMysql.Connect()
	articleServer.Launch()
}
