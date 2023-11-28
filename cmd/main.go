package main

import (
	articleServer "article-server/internal/app"
	CustomMysql "article-server/pkg/Mysql"
)

func main() {
	CustomMysql.Connect()
	articleServer.Launch()
}
