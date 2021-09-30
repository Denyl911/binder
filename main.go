package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	first, err := ioutil.ReadFile("programa1.exe")
	if err != nil {
		log.Fatal(err)
	}
	data1 := string(first)

	second, err := ioutil.ReadFile("programa2.exe")
	if err != nil {
		log.Fatal(err)
	}
	data2 := string(second)

	data := fmt.Sprintf(`package build

	import (
		"io/ioutil"
		"log"
		"os"
		"os/exec"
	)
	
	func main() {
		tmpfile1, err := ioutil.TempFile("", "*.exe")
		if err != nil {
			log.Fatal(err)
		}
		defer os.Remove(tmpfile1.Name())
	
		if _, err := tmpfile1.Write([]byte(%s)); err != nil {
			log.Fatal(err)
		}
		if err := tmpfile1.Close(); err != nil {
			log.Fatal(err)
		}
	
		tmpfile2, err := ioutil.TempFile("", "*.exe")
		if err != nil {
			log.Fatal(err)
		}
		defer os.Remove(tmpfile2.Name())
	
		if _, err := tmpfile2.Write([]byte(%s)); err != nil {
			log.Fatal(err)
		}
		if err := tmpfile2.Close(); err != nil {
			log.Fatal(err)
		}
	
		err = exec.Command(tmpfile1.Name()).Start()
		if err != nil {
			log.Fatal(err)
		}
		err = exec.Command(tmpfile2.Name()).Start()
		if err != nil {
			log.Fatal(err)
		}
	}
	`, data1, data2)

	err = ioutil.WriteFile("final.go", []byte(data), 0755)
	if err != nil {
		log.Fatal(err)
	}
	//err = exec.Command("go", "build", "final.go").Run()
	//if err != nil { log.Fatal(err) }

}
