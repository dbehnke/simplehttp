package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	var dir = flag.String("dir", "./", "directory to serve files from")
	var hostport = flag.String("hostport", ":8099", "host:port to serve on")
	flag.Parse()
	fmt.Printf("dir='%s', hostport='%s'\n", *dir, *hostport)

	r := gin.Default()
	r.StaticFS("/", http.Dir(*dir))
	r.Run(*hostport)
}
