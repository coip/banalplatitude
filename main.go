package main

import (
	"net/http"
	"os"
)

const s_ack = "#"

var (
	c_ack = func(id string) []byte { return []byte("["+s_ack+id+"]"+"\n") }

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
			if id, present := r.Header["id"]; present {
				print(id[0])
				w.Write(c_ack(id[0]))
			} else {
				w.WriteHeader(http.StatusOK)
			}	
		}),
	))

}
