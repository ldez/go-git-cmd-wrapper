package init

import (
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/ldez/go-git-cmd-wrapper/v2/git"
)

func TestInit(t *testing.T) {
	//nolint:usetesting // Don't use `t.TempDir()` because of a bug with Windows on the CI
	dir, err := os.MkdirTemp("", "go-git-cmd-wrapper")
	if err != nil {
		t.Fatal(err)
	}

	// clean up
	t.Cleanup(func() {
		errRm := os.RemoveAll(dir)
		if errRm != nil {
			log.Println(errRm)
		}
	})

	err = os.Chdir(dir)
	if err != nil {
		t.Fatal(err)
	}

	msg, err := git.Init(Directory("test"))
	if err != nil {
		t.Fatal(msg, err)
	}

	ff, err := os.Stat(filepath.Join(dir, "test"))
	if os.IsNotExist(err) {
		t.Fatal("Repository not created.", ff)
	}
}
