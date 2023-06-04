package main

import (
	"fmt"
	"io"
	"os"
)

func fopen(name string) *os.File {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	return file
}

func main() {

	one := fopen("./samples/1.mp4")
	defer one.Close()

	two := fopen("./samples/2.mp4")
	defer two.Close()

	output, err := os.Create("./output.mp4")
	if err != nil {
		panic(err)
	}
	defer output.Close()

	n, err := io.Copy(output, one)
	if err != nil {
		panic(err)
	}

	fmt.Println(n)

	n, err = io.Copy(output, two)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)

	fmt.Println("合体")
}
