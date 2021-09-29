package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	first, err := ioutil.ReadFile("programa1.exe")
	if err != nil {
		log.Fatal(err)
	}
	data1 := base64.StdEncoding.EncodeToString(first)

	second, err := ioutil.ReadFile("programa2.exe")
	if err != nil {
		log.Fatal(err)
	}
	data2 := base64.StdEncoding.EncodeToString(second)

	data := fmt.Sprintf(`package main

	import (
		"io/ioutil"
		"log"
		"os"
		"os/exec"
		"encoding/base64"
	)
	
	func main() {
		tmpfile1, err := ioutil.TempFile("", "*.exe")
		if err != nil {
			log.Fatal(err)
		}
		defer os.Remove(tmpfile1.Name())

		uno, _ := base64.StdEncoding.DecodeString("%s")
	
		if _, err := tmpfile1.Write(uno); err != nil {
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

		dos, _ := base64.StdEncoding.DecodeString("%s")
	
		if _, err := tmpfile2.Write(dos); err != nil {
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
