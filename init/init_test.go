package init

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/ldez/go-git-cmd-wrapper/git"
)

func TestInit(t *testing.T) {
	dir, err := ioutil.TempDir("", "go-git-cmd-wrapper")
	if err != nil {
		log.Fatal(err)
	}

	// clean up
	defer os.RemoveAll(dir)

	os.Chdir(dir)
	fmt.Println(os.Getwd())

	msg, err := git.Init(Directory("test"))

	log.Println(msg, err)
}
