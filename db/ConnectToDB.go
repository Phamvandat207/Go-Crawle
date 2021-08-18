package db

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/film")

	if err != nil {
		panic(err)
	}
	fmt.Println("Database: connected")

	return db, nil
}

func CreateMovieTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS movie(movie_id int primary key auto_increment, movie_name text, movie_year varchar(10), movie_rating varchar(10) )`
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	res, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating movie table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when getting rows affected", err)
		return err
	}
	log.Printf("Rows affected when creating table: %d", rows)
	return nil
}
