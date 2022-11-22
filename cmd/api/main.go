package main

import (
	"github.com/mecitsemerci/consul-demo/goserver/server"
	"log"
)

func main() {

	s := server.New()

	if err := s.Run(); err != nil {
		log.Fatalln("Server error ", err)
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
