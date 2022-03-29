package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		getContextData := r.Context().Value("data")
		fmt.Println(getContextData)
		io.WriteString(rw, `{"message": "hello word.."}`)
	})

	m := TheLogger(r)
	http.ListenAndServe(":8000", m)

}

func TheLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "data", "data:test")
		fmt.Println(fmt.Sprintf("%s %s", r.Method, r.URL.Path))
		next.ServeHTTP(rw, r.WithContext(ctx))
	})
}
