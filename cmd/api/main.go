package main

import (
	"flag"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"todo/internal/data"
)

type config struct {
	env  string
	port int
	db   struct {
		dsn string
	}
}

type application struct {
	cfg    config
	logger *log.Logger
	models data.Models
}

func main() {

	// Initialize config
	var cfg config
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|testing|production)")
	flag.IntVar(&cfg.port, "port", 8080, "API server port")
	dsn := fmt.Sprintf("root:pass@tcp(127.0.0.1:3306)/todolist?charset=utf8mb4&parseTime=True&loc=Local")
	flag.StringVar(&cfg.db.dsn, "db-dsn", dsn, "MySQL DSN")
	flag.Parse()

	// Create a logger
	logger := log.New(os.Stdout, "[log]", log.Ldate|log.Ltime)

	// Connect to db
	db := openDB(cfg)
	logger.Println("connect to db successfully")
	// Initialize application
	app := &application{
		cfg:    cfg,
		logger: logger,
		models: data.NewModel(db),
	}

	// Router
	r := app.routes()
	r.Run(fmt.Sprintf(":%d", app.cfg.port))
}

func openDB(cfg config) *gorm.DB {
	db, err := gorm.Open(mysql.Open(cfg.db.dsn), &gorm.Config{})
	if err != nil {
		panic("cannot connect to db")
	}
	return db
}
