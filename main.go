package main

import (
	"github.com/CyrusJavan/portfolio-new/src/blog"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	blog.Run()
}
