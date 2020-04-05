package main

import (
	"github.com/heroku/go-getting-started/src/blog"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	blog.Run()
}
