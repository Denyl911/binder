package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("Insert the name of the first file:")
	f1 := bufio.NewReader(os.Stdin)
	file1, _ := f1.ReadString('\n')
	file1 = strings.TrimSuffix(file1, "\n")

	// Open file 1
	first, err := ioutil.ReadFile(file1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Insert the name of the second file:")
	f2 := bufio.NewReader(os.Stdin)
	file2, _ := f2.ReadString('\n')
	file2 = strings.TrimSuffix(file2, "\n")

	second, err := ioutil.ReadFile(file2)
	if err != nil {
		log.Fatal(err)
	}

	fsjoin := [][]byte{first, second}
	final := bytes.Join(fsjoin, []byte(""))

	fmt.Println("Insert the name of the final file:")
	f3 := bufio.NewReader(os.Stdin)
	file3, _ := f3.ReadString('\n')
	file3 = strings.TrimSuffix(file3, "\n")

	err = ioutil.WriteFile(file3, final, 0755)
	if err != nil {
		log.Fatal(err)
	}
}
