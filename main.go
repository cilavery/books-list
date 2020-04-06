package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Book struct {
	ID int `json:id`
	Title string `json:title`
	Author string `jsong:author`
	Year string `json:year`
}

var books []Book

func main() {

}
