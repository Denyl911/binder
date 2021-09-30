package build

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

	if _, err := tmpfile1.Write(%s); err != nil {
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

	if _, err := tmpfile2.Write(%s); err != nil {
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
