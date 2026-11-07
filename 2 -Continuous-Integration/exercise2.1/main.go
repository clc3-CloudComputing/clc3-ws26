package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received from", r.RemoteAddr)
	t := roundToNearestMinute(time.Now())
	fmt.Fprintf(w, "Hello, it is %d:%d", t.Hour(), t.Minute())
}

func roundToNearestMinute(t time.Time) time.Time {
	return t.Round(time.Minute)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
