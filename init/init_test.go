package init

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"testing"

	"github.com/ldez/go-git-cmd-wrapper/git"
)

func TestInit(t *testing.T) {
	dir, err := ioutil.TempDir("", "go-git-cmd-wrapper")
	if err != nil {
		t.Fatal(err)
	}

	// clean up
	defer func() {
		errRm := os.RemoveAll(dir)
		log.Println(errRm)
	}()

	err = os.Chdir(dir)
	if err != nil {
		t.Fatal(err)
	}

	msg, err := git.Init(Directory("test"))
	if err != nil {
		t.Fatal(msg, err)
	}

	if ff, err := os.Stat(path.Join(dir, "test")); os.IsNotExist(err) {
		t.Fatal("Repositoty not created.", ff)
	}
}
