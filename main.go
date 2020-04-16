package main

import (
	"github.com/CyrusJavan/portfolio-new/src/blog"
	_ "github.com/heroku/x/hmetrics/onload"
	_ "github.com/lib/pq"
)

func main() {
	blog.Run()
}
