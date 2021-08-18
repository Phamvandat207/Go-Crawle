package main

import (
	"github.com/Phamvandat207/CrawlerIMDB/db"
	"github.com/Phamvandat207/CrawlerIMDB/util"
	"time"
)


func main()  {
	database, _ := db.ConnectDB()
	db.CreateMovieTable(database)
	time.Sleep(2 * time.Second)
	util.Crawler(database)
}
