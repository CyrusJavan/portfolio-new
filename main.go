package main

import (
	"github.com/CyrusJavan/blog/src/blog"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	blog.Run()
}
