package main

import (
	// "./controllers/users"
	// "github.com/gin-gonic/gin"
	"./router"
)


func main() {
	// test api
	r:= router.Router()
	r.Run() // listen and serve on 0.0.0.0:8080
}
