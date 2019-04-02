package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo"
	"log"
	"os"
)

func main() {
	e := echo.New()

	// Static Files
	e.Static("/public", "public")

	// Define the HTTP routes
	e.File("/", "public/index.html")
	e.PUT("/note", func(c echo.Context) error {

		m := echo.Map{}
		if err := c.Bind(&m); err != nil {
			return err
		}

		note := m["note"]

		// ENV
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		dbDatabse := os.Getenv("DB_DATABASE")
		dbUsername := os.Getenv("DB_USERNAME")
		dbPassowrd := os.Getenv("DB_PASSWORD")

		// DB CONNECT
		db, err := sql.Open("mysql", dbUsername+":"+dbPassowrd+"@tcp("+dbHost+":"+dbPort+")/"+dbDatabse)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		// INSERT
		result, err := db.Exec("INSERT INTO comento_members_note (note, date) VALUES (?, now())", note)
		if err != nil {
			log.Fatal(err)
		}

		return c.JSON(200, result)
	})

	e.Logger.Fatal(e.Start(":1323"))
}