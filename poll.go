package main

import (
//	"encoding/json"
	"fmt"
	"database/sql"
	"realtime-poll-go-pusher/handlers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3"
)

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	if err != nil {
		panic(err)
	}

	if db == nil {
		panic("db nil")
	}

	return db
}

func migrate(db *sql.DB) {
	fmt.Printf("\nprint: func migrate ")
	sql := `
        CREATE TABLE IF NOT EXISTS polls(
                id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
                name VARCHAR NOT NULL,
                topic VARCHAR NOT NULL,
                src VARCHAR NOT NULL,
                upvotes INTEGER NOT NULL,
                downvotes INTEGER NOT NULL,
                UNIQUE(name)
        );

        INSERT OR IGNORE INTO polls(name, topic, src, upvotes, downvotes) VALUES('Angular','Awesome Angular', '', 1, 0);

        INSERT OR IGNORE INTO polls(name, topic, src, upvotes, downvotes) VALUES('Vue', 'Voguish Vue','', 1, 0);

        INSERT OR IGNORE INTO polls(name, topic, src, upvotes, downvotes) VALUES('React','Remarkable React','', 1, 0);

        INSERT OR IGNORE INTO polls(name, topic, src, upvotes, downvotes) VALUES('Ember','Excellent Ember','', 1, 0);

        INSERT OR IGNORE INTO polls(name, topic, src, upvotes, downvotes) VALUES('Knockout','Knightly Knockout','', 1, 0);
   `
	_, err := db.Exec(sql)
	//fmt.Println(string(bytes))
	//panic(err)
	if err != nil {
		panic(err)
	}
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Initialize the database
	db := initDB("storage.db")
	fmt.Printf("print: func main ")
	migrate(db)

	// Define the HTTP routes
	e.File("/", "public/index.html")
	e.GET("/polls", handlers.GetPolls(db))
        fmt.Printf("print: db: %#v",db)
	e.PUT("/update/:index", handlers.UpdatePoll(db))

	// Start server
	e.Logger.Fatal(e.Start(":9000"))
}
