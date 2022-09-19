package main

import (
	"embed"
	"io/ioutil"
	"log"
)

//go:embed include/*
var included embed.FS

type Scripts struct{}

func (sc *Scripts) Read(fname string) []byte {
	file, err := included.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	return content
}

// Read the example.py shown in the input box by default
func (sc *Scripts) ReadExample() string {
	return string(sc.Read("include/example.py"))
}

func NewScripts() Scripts {
	return Scripts{}
}
