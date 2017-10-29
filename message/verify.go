package message

import (
	"fmt"
	"net/http"
)

func VerifyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Responding to challenge")

	res := http.Response{
		Status:  "HTTP 200 OK",
		Body:  r.Body,
	}
	err := res.Write(w)
	if err != nil {
		// handle
	}
	fmt.Println("Verified!")
}
