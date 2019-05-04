package main

import (
	"os"
	"log"
	"fmt"
	"net/http"
)

// Log 用于默认输出访问日志
func Log(handler http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
	handler.ServeHTTP(w, r)
    })
}

func main() {
	color := os.Getenv("COLOR")
	hostname, err := os.Hostname()
	if err != nil{
		log.Fatal(err)
	}
	if color == ""{
		color = "CYAN"
	}
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<html><head></head><body bgcolor='%s'>hostname: %s</body></html>", color, hostname)
	})
	http.ListenAndServe(":8080", Log(http.DefaultServeMux))
}
