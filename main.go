package main

import (
	"net/http"
	"os"
)

var (
	s_ack = "👍"
	c_ack = []byte(s_ack + "\n")

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
			w.Write(c_ack)
		}),
	))

}
