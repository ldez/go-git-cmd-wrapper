package git

import (
	"fmt"
	"strings"

	"github.com/ldez/go-git-cmd-wrapper/tag"

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
	out, _ := Init(ginit.Bare, ginit.Quiet, ginit.Directory("foobar"), CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git init --bare --quiet foobar
}

func ExamplePush() {
	out, _ := Push(push.All, push.FollowTags, push.ReceivePack("aaa"), CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git push --all --follow-tags --receive-pack=aaa
}

func ExamplePull() {
	out, _ := Pull(pull.All, pull.Force, pull.Repository("upstream"), pull.Refspec("master"), CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git pull --all --force upstream master
}

func ExampleClone() {
	out, _ := Clone(clone.Repository("git@github.com:ldez/go-git-cmd-wrapper.git"), CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git clone git@github.com:ldez/go-git-cmd-wrapper.git
}

func ExampleRemote() {
	out, _ := Remote(remote.Add("upstream", "git@github.com:johndoe/go-git-cmd-wrapper.git"), CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git remote add upstream git@github.com:johndoe/go-git-cmd-wrapper.git
}

func ExampleFetch() {
	out, _ := Fetch(fetch.NoTags, fetch.Remote("upstream"), fetch.RefSpec("myBranchName"), CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git fetch --no-tags upstream myBranchName
}

func ExampleRebase() {
	out, _ := Rebase(rebase.PreserveMerges, rebase.Branch(fmt.Sprintf("%s/%s", "upstream", "master")), CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git rebase --preserve-merges upstream/master
}

func ExampleCheckout() {
	out, _ := Checkout(checkout.NewBranch("myBranchName"), CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git checkout -b myBranchName
}

func ExampleConfig() {
	out, _ := Config(config.Entry("rebase.autoSquash", "true"), CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git config rebase.autoSquash true
}

func ExampleBranch() {
	out, _ := Branch(branch.DeleteForce, branch.BranchName("myBranch"), CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git branch -D myBranch
}

func ExampleRevParse() {
	out, _ := RevParse(revparse.AbbrevRef(""), revparse.Args("HEAD"), CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git rev-parse --abbrev-ref HEAD
}

func ExampleReset() {
	out, _ := Reset(reset.Soft, reset.Commit("e41f083"), CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git reset --soft e41f083
}

func ExampleCommit() {
	out, _ := Commit(commit.Amend, commit.Message("chore: foo"), CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git commit --amend --message="chore: foo"
}

func ExampleAdd() {
	out, _ := Add(add.All, CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git add --all
}

func ExampleMerge() {
	out, _ := Merge(merge.Squash, merge.Commits("myBranch"), CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git merge --squash myBranch
}

func ExampleCond() {
	param := false
	out, _ := Push(push.All, Cond(param, push.DryRun), push.FollowTags, push.ReceivePack("aaa"), CmdExecutor(cmdExecutorMock))

	fmt.Print(out)

	param = true
	out, _ = Push(push.All, Cond(param, push.DryRun), push.FollowTags, push.ReceivePack("aaa"), CmdExecutor(cmdExecutorMock))

	fmt.Print(out)

	// Output:
	// git push --all --follow-tags --receive-pack=aaa
	// git push --all --dry-run --follow-tags --receive-pack=aaa
}

func ExampleWorktree() {
	out, _ := Worktree(worktree.Add("v1.0", "origin/v1.0"), CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git worktree add v1.0 origin/v1.0
}

func ExampleTag() {
	out, _ := Tag(tag.List, CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git tag --list
}

func cmdExecutorMock(name string, _ bool, args ...string) (string, error) {
	return fmt.Sprintln(name, strings.Join(args, " ")), nil
}
