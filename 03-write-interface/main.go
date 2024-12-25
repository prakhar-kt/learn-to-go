package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type person struct {
	first string
}

func (p person) writeOut(w io.Writer) {
	w.Write([]byte(p.first))
}

func main() {

	p := person{
		first: "James",
	}
	fs, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("Error %s", err)
	}

	defer fs.Close()

	var b bytes.Buffer

	p.writeOut(fs)
	p.writeOut(&b)
	fmt.Println(b.String())

}
