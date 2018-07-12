package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//http get returns the result and an error variable
	//putting the blank identifier, '_',  in place of the error variable allows
	//you to ignore it
	var result, _ = http.Get("https://github.com/DarkenedLight")
	var page, _ = ioutil.ReadAll(result.Body)
	result.Body.Close()
	fmt.Printf("%s", page)
}

//blank identifier, '_', says you are not using a variable
