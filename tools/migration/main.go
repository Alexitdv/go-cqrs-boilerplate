package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kelseyhightower/envconfig"

	"github.com/pressly/goose"
)

var (
	flags = flag.NewFlagSet("migration", flag.ExitOnError)
	dir   = flags.String("dir", "./migrations/sql", "directory with migration files")
)

type Options struct {
	PGUser     string `envconfig:"PG_USER" required:"true"`
	PGPass     string `envconfig:"PG_PASSWORD" required:"true"`
	PGHost     string `envconfig:"PG_HOST" required:"true"`
	PGPort     string `envconfig:"PG_PORT" default:"5432"`
	PGDatabase string `envconfig:"PG_DATABASE" default:"boilerplate"`
}

func main() {
	err := flags.Parse(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	args := flags.Args()

	var conf Options
	err = envconfig.Process("", &conf)
	if err != nil {
		log.Fatal(err.Error())
	}

	if len(args) < 1 {
		flags.Usage()
		return
	}

	command := args[0]
	var dbURL = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?sslmode=disable",
		conf.PGUser, conf.PGPass, conf.PGHost, conf.PGPort, conf.PGDatabase)

	db, err := sql.Open("mysql", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if cErr := db.Close(); cErr != nil {
			log.Fatalf("migration: failed to close DB: %v\n", err)
		}
	}()

	err = db.Ping()
	if err != nil {
		log.Panic(err)
	}

	arguments := []string{}
	if len(args) > 1 {
		arguments = append(arguments, args[1:]...)
	}

	if err := goose.SetDialect("postgres"); err != nil {
		log.Panic(err)
	}

	if err := goose.Run(command, db, *dir, arguments...); err != nil {
		log.Panicf("migration %v: %v", command, err)
	}
}
