package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"user-crud/router"

	"user-crud/config"

	"github.com/go-chi/valve"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(err.Error())
	}

	dbConfig := config.NewDBConfig()
	defer dbConfig.Db.Close()

	valv := valve.New()
	baseCtx := valv.Context()

	r := router.Router(dbConfig)

	srv := http.Server{Addr: ":8000", Handler: r}
	srv.BaseContext = func(_ net.Listener) context.Context {
		return baseCtx
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			fmt.Println("shutting down..")

			valv.Shutdown(20 * time.Second)

			ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
			defer cancel()

			srv.Shutdown(ctx)

			select {
			case <-time.After(21 * time.Second):
				fmt.Println("not all connections done")
			case <-ctx.Done():

			}
		}
	}()
	fmt.Println("Server running at port 8000")
	srv.ListenAndServe()
}
