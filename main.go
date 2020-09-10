package main

import (
	"net/http"
	"os"
)

const s_ack = "#"

var (
	c_ack = func(id string) []byte { return []byte("[" + s_ack + id + "]" + "\n") }

	port string
)

func main() {

	if port = os.Getenv("port"); len(port) == 0 {
		port = ":8080"
	}

	print("binding to port: [" + port + "] ..." + "\n")

	print(http.ListenAndServe(port,
		(http.HandlerFunc)(func(w http.ResponseWriter, r *http.Request) {
			print(s_ack)
			if id := r.Header.Get("id"); len(id) > 0 {
				print(id)
				w.Write(c_ack(id))
			} else {
				w.WriteHeader(http.StatusOK)
			}
		}),
	))

}
