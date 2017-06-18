package git

import (
	"fmt"
	"strings"

	"github.com/ldez/go-git-cmd-wrapper/branch"
	"github.com/ldez/go-git-cmd-wrapper/checkout"
	"github.com/ldez/go-git-cmd-wrapper/clone"
	"github.com/ldez/go-git-cmd-wrapper/config"
	"github.com/ldez/go-git-cmd-wrapper/fetch"
	ginit "github.com/ldez/go-git-cmd-wrapper/init"
	"github.com/ldez/go-git-cmd-wrapper/push"
	"github.com/ldez/go-git-cmd-wrapper/rebase"
	"github.com/ldez/go-git-cmd-wrapper/remote"
	"github.com/ldez/go-git-cmd-wrapper/revparse"
)

func ExampleInit() {
	cmdExecutor = cmdExecutorMock

	out, _ := Init(ginit.Bare, ginit.Quiet, ginit.Directory("foobar"))

	fmt.Println(out)
	// Output: git init --bare --quiet foobar
}

func ExamplePush() {
	cmdExecutor = cmdExecutorMock

	out, _ := Push(push.All, push.FollowTags, push.ReceivePack("aaa"))

	fmt.Println(out)
	// Output: git push --all --follow-tags --receive-pack=aaa
}

func ExampleClone() {
	cmdExecutor = cmdExecutorMock

	out, _ := Clone(clone.Repository("git@github.com:ldez/go-git-cmd-wrapper.git"))

	fmt.Println(out)
	// Output: git clone git@github.com:ldez/go-git-cmd-wrapper.git
}

func ExampleRemote() {
	cmdExecutor = cmdExecutorMock

	out, _ := Remote(remote.Add("upstream", "git@github.com:johndoe/go-git-cmd-wrapper.git"))

	fmt.Println(out)
	// Output: git remote add upstream git@github.com:johndoe/go-git-cmd-wrapper.git
}

func ExampleFetch() {
	cmdExecutor = cmdExecutorMock

	out, _ := Fetch(fetch.NoTags, fetch.Remote("upstream"), fetch.RefSpec("myBranchName"))

	fmt.Println(out)
	// Output: git fetch --no-tags upstream myBranchName
}

func ExampleRebase() {
	cmdExecutor = cmdExecutorMock

	out, _ := Rebase(rebase.PreserveMerges, rebase.Branch(fmt.Sprintf("%s/%s", "upstream", "master")))

	fmt.Println(out)
	// Output: git rebase --preserve-merges upstream/master
}

func ExampleCheckout() {
	cmdExecutor = cmdExecutorMock

	out, _ := Checkout(checkout.NewBranch, checkout.Branch("myBranchName"))

	fmt.Println(out)
	// Output: git checkout -b myBranchName
}

func ExampleConfig() {
	cmdExecutor = cmdExecutorMock

	out, _ := Config(config.Entry("rebase.autoSquash", "true"))

	fmt.Println(out)
	// Output: git config rebase.autoSquash true
}

func ExampleBranch() {
	cmdExecutor = cmdExecutorMock

	out, _ := Branch(branch.DeleteForce, branch.BranchName("myBranch"))

	fmt.Println(out)
	// Output: git branch -D myBranch
}

func ExampleRevParse() {
	cmdExecutor = cmdExecutorMock

	out, _ := RevParse(revparse.AbbrevRef(""), revparse.Args("HEAD"))

	fmt.Println(out)
	// Output: git rev-parse --abbrev-ref HEAD
}

func ExampleCond() {
	cmdExecutor = cmdExecutorMock

	param := false
	out, _ := Push(push.All, Cond(param, push.DryRun), push.FollowTags, push.ReceivePack("aaa"))

	fmt.Print(out)

	param = true
	out, _ = Push(push.All, Cond(param, push.DryRun), push.FollowTags, push.ReceivePack("aaa"))

	fmt.Print(out)

	// Output:
	// git push --all --follow-tags --receive-pack=aaa
	// git push --all --dry-run --follow-tags --receive-pack=aaa
}

func cmdExecutorMock(name string, _ bool, args ...string) (string, error) {
	return fmt.Sprintln(name, strings.Join(args, " ")), nil
}
