package main

import (
	"itish2003/image-primitive/router"
)

func main() {
	r := router.Router()
	r.Run(":8080")
}
