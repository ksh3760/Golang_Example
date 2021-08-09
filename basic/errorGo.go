package main

import (
	"log"
	"os"
)

func errorGo() {
	f, err := os.Open("./file/testFile.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	println(f.Name())
}

_, err := otherFunc()
switch err.(type)	{
default: // no error
	println("ok")
case MyError:
	log.Print("Log my error")
case error:
	log.Fatal(err.Error())
}
