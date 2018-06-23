package git

import (
	"fmt"
	"strings"

	"github.com/ldez/go-git-cmd-wrapper/add"
	"github.com/ldez/go-git-cmd-wrapper/branch"
	"github.com/ldez/go-git-cmd-wrapper/checkout"
	"github.com/ldez/go-git-cmd-wrapper/clone"
	"github.com/ldez/go-git-cmd-wrapper/commit"
	"github.com/ldez/go-git-cmd-wrapper/config"
	"github.com/ldez/go-git-cmd-wrapper/fetch"
	ginit "github.com/ldez/go-git-cmd-wrapper/init"
	"github.com/ldez/go-git-cmd-wrapper/merge"
	"github.com/ldez/go-git-cmd-wrapper/pull"
	"github.com/ldez/go-git-cmd-wrapper/push"
	"github.com/ldez/go-git-cmd-wrapper/rebase"
	"github.com/ldez/go-git-cmd-wrapper/remote"
	"github.com/ldez/go-git-cmd-wrapper/reset"
	"github.com/ldez/go-git-cmd-wrapper/revparse"
	"github.com/ldez/go-git-cmd-wrapper/worktree"
)

func ExampleInit() {
	CmdExecutor = cmdExecutorMock

	out, _ := Init(ginit.Bare, ginit.Quiet, ginit.Directory("foobar"))

	fmt.Println(out)
	// Output: git init --bare --quiet foobar
}

func ExamplePush() {
	CmdExecutor = cmdExecutorMock

	out, _ := Push(push.All, push.FollowTags, push.ReceivePack("aaa"))

	fmt.Println(out)
	// Output: git push --all --follow-tags --receive-pack=aaa
}

func ExamplePull() {
	CmdExecutor = cmdExecutorMock

	out, _ := Pull(pull.All, pull.Force, pull.Repository("upstream"), pull.Refspec("master"))

	fmt.Println(out)
	// Output: git pull --all --force upstream master
}

func ExampleClone() {
	CmdExecutor = cmdExecutorMock

	out, _ := Clone(clone.Repository("git@github.com:ldez/go-git-cmd-wrapper.git"))

	fmt.Println(out)
	// Output: git clone git@github.com:ldez/go-git-cmd-wrapper.git
}

func ExampleRemote() {
	CmdExecutor = cmdExecutorMock

	out, _ := Remote(remote.Add("upstream", "git@github.com:johndoe/go-git-cmd-wrapper.git"))

	fmt.Println(out)
	// Output: git remote add upstream git@github.com:johndoe/go-git-cmd-wrapper.git
}

func ExampleFetch() {
	CmdExecutor = cmdExecutorMock

	out, _ := Fetch(fetch.NoTags, fetch.Remote("upstream"), fetch.RefSpec("myBranchName"))

	fmt.Println(out)
	// Output: git fetch --no-tags upstream myBranchName
}

func ExampleRebase() {
	CmdExecutor = cmdExecutorMock

	out, _ := Rebase(rebase.PreserveMerges, rebase.Branch(fmt.Sprintf("%s/%s", "upstream", "master")))

	fmt.Println(out)
	// Output: git rebase --preserve-merges upstream/master
}

func ExampleCheckout() {
	CmdExecutor = cmdExecutorMock

	out, _ := Checkout(checkout.NewBranch("myBranchName"))

	fmt.Println(out)
	// Output: git checkout -b myBranchName
}

func ExampleConfig() {
	CmdExecutor = cmdExecutorMock

	out, _ := Config(config.Entry("rebase.autoSquash", "true"))

	fmt.Println(out)
	// Output: git config rebase.autoSquash true
}

func ExampleBranch() {
	CmdExecutor = cmdExecutorMock

	out, _ := Branch(branch.DeleteForce, branch.BranchName("myBranch"))

	fmt.Println(out)
	// Output: git branch -D myBranch
}

func ExampleRevParse() {
	CmdExecutor = cmdExecutorMock

	out, _ := RevParse(revparse.AbbrevRef(""), revparse.Args("HEAD"))

	fmt.Println(out)
	// Output: git rev-parse --abbrev-ref HEAD
}

func ExampleReset() {
	CmdExecutor = cmdExecutorMock

	out, _ := Reset(reset.Soft, reset.Commit("e41f083"))

	fmt.Println(out)
	// Output: git reset --soft e41f083
}

func ExampleCommit() {
	CmdExecutor = cmdExecutorMock

	out, _ := Commit(commit.Amend, commit.Message("chore: foo"))

	fmt.Println(out)
	// Output: git commit --amend --message="chore: foo"
}

func ExampleAdd() {
	CmdExecutor = cmdExecutorMock

	out, _ := Add(add.All)

	fmt.Println(out)
	// Output: git add --all
}

func ExampleMerge() {
	CmdExecutor = cmdExecutorMock

	out, _ := Merge(merge.Squash, merge.Commits("myBranch"))

	fmt.Println(out)
	// Output: git merge --squash myBranch
}

func ExampleCond() {
	CmdExecutor = cmdExecutorMock

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

func ExampleWorktree() {
	CmdExecutor = cmdExecutorMock

	out, _ := Worktree(worktree.Add("v1.0", "origin/v1.0"))

	fmt.Println(out)
	// Output: git worktree add v1.0 origin/v1.0
}

func cmdExecutorMock(name string, _ bool, args ...string) (string, error) {
	return fmt.Sprintln(name, strings.Join(args, " ")), nil
}
