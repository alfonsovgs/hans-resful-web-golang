package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

// Args define a new type
type Args struct{}

// TimeServer define a new type
type TimeServer int64

// GiveServerTime return the time.Now
func (t *TimeServer) GiveServerTime(args *Args, reply *int64) error {
	// Fill reply pointer to send the data back
	*reply = time.Now().Unix()
	return nil
}

func main() {
	timeServer := new(TimeServer)
	rpc.Register(timeServer)
	rpc.HandleHTTP()

	// Listen for request on port 1234
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error: ", e)
	}

	http.Serve(l, nil)
}
