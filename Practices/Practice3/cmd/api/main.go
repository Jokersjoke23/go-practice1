package main

import (
	"Practice3/internal/handlers"
	"Practice3/internal/middleware"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// ✅ Открытый маршрут (доступен всем)
	mux.HandleFunc("/public", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "This is a public page — no API key needed.")
	})

	// ✅ Защищённый маршрут — middleware применяется только к /user
	mux.Handle("/user", middleware.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Protected user route: authorized access!")
	})))

	// ✅ Остальные маршруты (если есть)
	mux.HandleFunc("/", handlers.HomeHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Println("✅ Server is running on http://localhost:8080")

	if err := server.ListenAndServe(); err != nil {
		fmt.Println("❌ Server error:", err)
	}
}
