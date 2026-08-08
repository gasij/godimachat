package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"

	"godima/internal/config"
	"godima/internal/database"
	"godima/internal/handlers"
	"godima/internal/hub"
	"godima/internal/repository"
)

func main() {
	_ = godotenv.Load()

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	pool, err := database.Connect(context.Background(), cfg.DSN())
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	userRepo := repository.NewUserRepository(pool)
	msgRepo := repository.NewMessageRepository(pool)

	chatHub := hub.New()
	go chatHub.Run()

	apiHandler := handlers.New(userRepo, msgRepo, chatHub, cfg.JWT_SECRET)
	pagesHandler, err := handlers.NewPagesHandler("web/templates")
	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()
	r.Use(chiMiddleware.Logger)
	r.Use(chiMiddleware.Recoverer)

	// Статика
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	// HTML
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	})
	r.Get("/login", pagesHandler.Login)
	r.Get("/register", pagesHandler.Register)
	r.Get("/chat", pagesHandler.Chat)

	// API
	r.Mount("/", apiHandler.Routes())

	addr := ":" + cfg.SERVER_PORT
	server := &http.Server{Addr: addr, Handler: r}

	go func() {
		fmt.Printf("Сервер: http://localhost%s\n", addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	server.Shutdown(ctx)
}
