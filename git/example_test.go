package git_test

import (
	"fmt"
	"strings"

	"github.com/ldez/go-git-cmd-wrapper/git"
	"github.com/ldez/go-git-cmd-wrapper/status"
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
	out, _ := git.Init(ginit.Bare, ginit.Quiet, ginit.Directory("foobar"), git.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git init --bare --quiet foobar
}

func ExamplePush() {
	out, _ := git.Push(push.All, push.FollowTags, push.ReceivePack("aaa"), git.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git push --all --follow-tags --receive-pack=aaa
}

func ExamplePull() {
	out, _ := git.Pull(pull.All, pull.Force, pull.Repository("upstream"), pull.Refspec("master"), git.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git pull --all --force upstream master
}

func ExampleClone() {
	out, _ := git.Clone(clone.Repository("git@github.com:ldez/go-git-cmd-wrapper.git"), git.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git clone git@github.com:ldez/go-git-cmd-wrapper.git
}

func ExampleRemote() {
	out, _ := git.Remote(remote.Add("upstream", "git@github.com:johndoe/go-git-cmd-wrapper.git"), git.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git remote add upstream git@github.com:johndoe/go-git-cmd-wrapper.git
}

func ExampleFetch() {
	out, _ := git.Fetch(fetch.NoTags, fetch.Remote("upstream"), fetch.RefSpec("myBranchName"), git.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git fetch --no-tags upstream myBranchName
}

func ExampleRebase() {
	out, _ := git.Rebase(rebase.PreserveMerges, rebase.Branch(fmt.Sprintf("%s/%s", "upstream", "master")), git.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git rebase --preserve-merges upstream/master
}

func ExampleCheckout() {
	out, _ := git.Checkout(checkout.NewBranch("myBranchName"), git.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git checkout -b myBranchName
}

func ExampleConfig() {
	out, _ := git.Config(config.Entry("rebase.autoSquash", "true"), git.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git config rebase.autoSquash true
}

func ExampleBranch() {
	out, _ := git.Branch(branch.DeleteForce, branch.BranchName("myBranch"), git.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git branch -D myBranch
}

func ExampleRevParse() {
	out, _ := git.RevParse(revparse.AbbrevRef(""), revparse.Args("HEAD"), git.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git rev-parse --abbrev-ref HEAD
}

func ExampleReset() {
	out, _ := git.Reset(reset.Soft, reset.Commit("e41f083"), git.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git reset --soft e41f083
}

func ExampleCommit() {
	out, _ := git.Commit(commit.Amend, commit.Message("chore: foo"), git.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git commit --amend --message="chore: foo"
}

func ExampleAdd() {
	out, _ := git.Add(add.All, git.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git add --all
}

func ExampleMerge() {
	out, _ := git.Merge(merge.Squash, merge.Commits("myBranch"), git.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git merge --squash myBranch
}

func ExampleCond() {
	param := false
	out, _ := git.Push(push.All, git.Cond(param, push.DryRun), push.FollowTags, push.ReceivePack("aaa"), git.CmdExecutor(cmdExecutorMock))

	fmt.Print(out)

	param = true
	out, _ = git.Push(push.All, git.Cond(param, push.DryRun), push.FollowTags, push.ReceivePack("aaa"), git.CmdExecutor(cmdExecutorMock))

	fmt.Print(out)

	// Output:
	// git push --all --follow-tags --receive-pack=aaa
	// git push --all --dry-run --follow-tags --receive-pack=aaa
}

func ExampleWorktree() {
	out, _ := git.Worktree(worktree.Add("v1.0", "origin/v1.0"), git.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git worktree add v1.0 origin/v1.0
}

func ExampleTag() {
	out, _ := git.Tag(tag.List, git.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git tag --list
}

func ExampleStatus() {
	out, _ := git.Status(status.Short, status.Branch, git.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git status --short --branch
}

func cmdExecutorMock(name string, _ bool, args ...string) (string, error) {
	return fmt.Sprintln(name, strings.Join(args, " ")), nil
}
