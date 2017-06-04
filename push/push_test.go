package push

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/ldez/go-git-cmd-wrapper/git"
)

func TestPush(t *testing.T) {
	dir, err := ioutil.TempDir("", "go-git-cmd-wrapper")
	if err != nil {
		log.Fatal(err)
	}

	// clean up
	defer os.RemoveAll(dir)

	os.Chdir(dir)
	fmt.Println(os.Getwd())

	msg, err := git.Push(All, FollowTags, ReceivePack("fjd"))

	log.Println(msg, err)
}
