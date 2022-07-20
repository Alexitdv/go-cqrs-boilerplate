package main

import (
	"boilerplate/internal/app/logger"
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/kelseyhightower/envconfig"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"boilerplate/internal/app"
	server "boilerplate/internal/ports/grpc"
	pb "boilerplate/internal/ports/grpc/server/boilderplate/v1"

	_ "github.com/lib/pq"
)

type Options struct {
	GRPCPort   string `envconfig:"SRV_PORT" default:"5010"`
	HTTPPort   string `envconfig:"SRV_HTTP_PORT" default:"8080"`
	LogLevel   string `envconfig:"LOG_LEVEL" default:"info"`
	PGUser     string `envconfig:"PG_USER" required:"true"`
	PGPass     string `envconfig:"PG_PASSWORD" required:"true"`
	PGHost     string `envconfig:"PG_HOST" required:"true"`
	PGPort     string `envconfig:"PG_PORT" default:"5432"`
	PGDatabase string `envconfig:"PG_DATABASE" default:"boilerplate"`
}

func serve(ctx context.Context, conf *Options) error {
	dbURL := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s database=%s sslmode=disable",
		conf.PGUser, conf.PGPass, conf.PGHost, conf.PGPort, conf.PGDatabase,
	)

	db, err := sqlx.Open("postgres", dbURL)
	if err != nil {
		return err
	}
	defer func() {
		_ = db.Close()
	}()

	db.SetMaxIdleConns(30)

	err = db.Ping()
	if err != nil {
		return err
	}

	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", conf.GRPCPort))
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	a, err := app.NewApplication(&app.Options{
		DB: db,
	})
	if err != nil {
		return err
	}

	srv := server.NewGrpcServer(a)
	pb.RegisterBoilerplateServer(s, &srv)

	fmt.Println("GRPC Port:", conf.GRPCPort)
	fmt.Println("HTTP Port:", conf.HTTPPort)

	go func() {
		if cErr := s.Serve(listen); cErr != nil {
			logrus.Fatalf("failed to serve: %v", cErr)
		}
	}()

	mx := http.NewServeMux()
	mx.Handle("/metrics", promhttp.Handler())
	httpServer := &http.Server{Addr: fmt.Sprintf(":%s", conf.HTTPPort), Handler: mx}

	go func() {
		if cErr := httpServer.ListenAndServe(); cErr != nil {
			logrus.Fatalf("failed to serve: %v", cErr)
		}
	}()

	logrus.Info("Started")
	<-ctx.Done()
	logrus.Info("Stopped")

	s.GracefulStop()
	a.Shutdown()
	err = httpServer.Shutdown(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func main() {
	var conf Options
	err := envconfig.Process("", &conf)
	if err != nil {
		logrus.Fatal(err.Error())
	}
	err = logger.InitLogger(conf.LogLevel)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.WithField("config", conf).Info("server config ready to use")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		call := <-c
		logrus.Infof("system call: %+v", call)
		cancel()
	}()

	if err := serve(ctx, &conf); err != nil {
		logrus.Errorf("failed to serve:+%v\n", err)
	}
}
