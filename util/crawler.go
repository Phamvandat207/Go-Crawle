package util

import (
	"database/sql"
	"fmt"
	"github.com/Phamvandat207/CrawlerIMDB/model"
	"github.com/gocolly/colly"
	"log"
)

func Crawler(db *sql.DB)  {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) { //Đang gửi request get HTML
		fmt.Printf("Visiting: %sn", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) { //Handle error trong quá trình craw html
		log.Println("something's wrong here:", err)
	})

	c.OnResponse(func(r *colly.Response) { //Sau khi đã lấy được HTML
		fmt.Printf("Visited: %sn", r.Request.URL)
	})

	c.OnHTML("tr", func(e *colly.HTMLElement) { //Bóc tách dữ liệu từ HTML lấy được
		m := model.Movie{}
		m.Name = e.ChildText(".titleColumn > a")
		m.Year = e.ChildText(".titleColumn .secondaryInfo")
		m.Rating = e.ChildText(".ratingColumn > strong")
		fmt.Printf("- Title: %sn- Link: %sn- Description: %sn", m.Name, m.Year, m.Rating) //In ra màn hình giá trị đã lấy được

		stmt, err1 := db.Prepare("INSERT INTO movie (movie_name, movie_year, movie_rating) values (?,?,?)") //Prepare SQL cho việc insert
		checkErr(err1) //Handle error

		res, err2 := stmt.Exec(m.Name, m.Year, m.Rating) //Binding data vào câu query
		checkErr(err2) //Handle error

		lastId, err3 := res.LastInsertId() //Lấy ra ID vừa được insert

		if err3 != nil {
			log.Fatal(err3)
		}

		fmt.Printf("=&gt;Insert ID: %dnn", lastId) //In ra màn hình ID vừa insert
	})

	c.OnScraped(func(r *colly.Response) { //Hoàn thành job craw
		fmt.Println("Finished", r.Request.URL)
	})

	c.Visit("https://www.imdb.com/chart/top/?ref_=nv_mv_250") //Trình thu thập truy cập URL đó
}

func checkErr(err error) { //Thêm function để handle error
	if err != nil {
		panic(err)
	}
}