package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/kvii/di/config"
	"github.com/kvii/di/database"
	"github.com/kvii/di/logger"
	"github.com/kvii/di/service"
)

func main() {
	// 主要看这里的逻辑。
	ctx := context.Background()
	ctx = config.WithContext(ctx)
	ctx = logger.WithContext(ctx)
	ctx = database.WithContext(ctx)
	s := service.FromContext(ctx)

	// curl "http://localhost:8080/hello?name=world"
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		msg, err := s.Hello(ctx, name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			fmt.Fprintln(w, msg)
		}
	})

	fmt.Println("http://localhost:8080/hello?name=world")
	http.ListenAndServe(":8080", nil)
}
