package main 

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handlefunc("/", func (w http.ResponseWriter, r *http.Request))
}