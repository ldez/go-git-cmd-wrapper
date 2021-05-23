package git_test

import (
	"context"
	"fmt"
	"strings"

	"github.com/ldez/go-git-cmd-wrapper/v2/add"
	"github.com/ldez/go-git-cmd-wrapper/v2/branch"
	"github.com/ldez/go-git-cmd-wrapper/v2/checkout"
	"github.com/ldez/go-git-cmd-wrapper/v2/clone"
	"github.com/ldez/go-git-cmd-wrapper/v2/commit"
	"github.com/ldez/go-git-cmd-wrapper/v2/config"
	"github.com/ldez/go-git-cmd-wrapper/v2/fetch"
	"github.com/ldez/go-git-cmd-wrapper/v2/git"
	ginit "github.com/ldez/go-git-cmd-wrapper/v2/init"
	"github.com/ldez/go-git-cmd-wrapper/v2/merge"
	"github.com/ldez/go-git-cmd-wrapper/v2/pull"
	"github.com/ldez/go-git-cmd-wrapper/v2/push"
	"github.com/ldez/go-git-cmd-wrapper/v2/rebase"
	"github.com/ldez/go-git-cmd-wrapper/v2/remote"
	"github.com/ldez/go-git-cmd-wrapper/v2/reset"
	"github.com/ldez/go-git-cmd-wrapper/v2/revparse"
	"github.com/ldez/go-git-cmd-wrapper/v2/status"
	"github.com/ldez/go-git-cmd-wrapper/v2/tag"
	"github.com/ldez/go-git-cmd-wrapper/v2/types"
	"github.com/ldez/go-git-cmd-wrapper/v2/worktree"
)

func ExampleInit() {
	out, _ := git.Init(ginit.Bare, ginit.Quiet, ginit.Directory("foobar"), git.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git init --bare --quiet foobar
}

func ExampleInitWithContext() {
	out, _ := git.InitWithContext(context.Background(), ginit.Bare, ginit.Quiet, ginit.Directory("foobar"), git.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git init --bare --quiet foobar
}

func ExamplePush() {
	out, _ := git.Push(push.All, push.FollowTags, push.ReceivePack("aaa"), git.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git push --all --follow-tags --receive-pack=aaa
}

func ExamplePushWithContext() {
	out, _ := git.PushWithContext(context.Background(), push.All, push.FollowTags, push.ReceivePack("aaa"), git.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git push --all --follow-tags --receive-pack=aaa
}

func ExamplePull() {
	out, _ := git.Pull(pull.All, pull.Force, pull.Repository("upstream"), pull.Refspec("master"), git.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git pull --all --force upstream master
}

func ExamplePullWithContext() {
	out, _ := git.PullWithContext(context.Background(), pull.All, pull.Force, pull.Repository("upstream"), pull.Refspec("master"), git.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git pull --all --force upstream master
}

func ExampleClone() {
	out, _ := git.Clone(clone.Repository("git@github.com:ldez/go-git-cmd-wrapper.git"), git.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git clone git@github.com:ldez/go-git-cmd-wrapper.git
}

func ExampleCloneWithContext() {
	out, _ := git.CloneWithContext(context.Background(), clone.Repository("git@github.com:ldez/go-git-cmd-wrapper.git"), git.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git clone git@github.com:ldez/go-git-cmd-wrapper.git
}

func ExampleRemote() {
	out, _ := git.Remote(remote.Add("upstream", "git@github.com:johndoe/go-git-cmd-wrapper.git"), git.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git remote add upstream git@github.com:johndoe/go-git-cmd-wrapper.git
}

func ExampleRemoteWithContext() {
	out, _ := git.RemoteWithContext(context.Background(), remote.Add("upstream", "git@github.com:johndoe/go-git-cmd-wrapper.git"), git.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git remote add upstream git@github.com:johndoe/go-git-cmd-wrapper.git
}

func ExampleFetch() {
	out, _ := git.Fetch(fetch.NoTags, fetch.Remote("upstream"), fetch.RefSpec("myBranchName"), git.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git fetch --no-tags upstream myBranchName
}

func ExampleFetchWithContext() {
	out, _ := git.FetchWithContext(context.Background(), fetch.NoTags, fetch.Remote("upstream"), fetch.RefSpec("myBranchName"), git.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git fetch --no-tags upstream myBranchName
}

func ExampleRebase() {
	out, _ := git.Rebase(rebase.PreserveMerges, rebase.Branch(fmt.Sprintf("%s/%s", "upstream", "master")), git.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git rebase --preserve-merges upstream/master
}

func ExampleRebaseWithContext() {
	out, _ := git.RebaseWithContext(context.Background(), rebase.PreserveMerges, rebase.Branch(fmt.Sprintf("%s/%s", "upstream", "master")), git.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git rebase --preserve-merges upstream/master
}

func ExampleCheckout() {
	out, _ := git.Checkout(checkout.NewBranch("myBranchName"), git.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git checkout -b myBranchName
}

func ExampleCheckoutWithContext() {
	out, _ := git.CheckoutWithContext(context.Background(), checkout.NewBranch("myBranchName"), git.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git checkout -b myBranchName
}

func ExampleConfig() {
	out, _ := git.Config(config.Entry("rebase.autoSquash", "true"), git.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git config rebase.autoSquash true
}

func ExampleConfigWithContext() {
	out, _ := git.ConfigWithContext(context.Background(), config.Entry("rebase.autoSquash", "true"), git.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git config rebase.autoSquash true
}

func ExampleBranch() {
	out, _ := git.Branch(branch.DeleteForce, branch.BranchName("myBranch"), git.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git branch -D myBranch
}

func ExampleBranchWithContext() {
	out, _ := git.BranchWithContext(context.Background(), branch.DeleteForce, branch.BranchName("myBranch"), git.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git branch -D myBranch
}

func ExampleRevParse() {
	out, _ := git.RevParse(revparse.AbbrevRef(""), revparse.Args("HEAD"), git.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git rev-parse --abbrev-ref HEAD
}

func ExampleRevParseWithContext() {
	out, _ := git.RevParseWithContext(context.Background(), revparse.AbbrevRef(""), revparse.Args("HEAD"), git.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git rev-parse --abbrev-ref HEAD
}

func ExampleReset() {
	out, _ := git.Reset(reset.Soft, reset.Commit("e41f083"), git.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git reset --soft e41f083
}

func ExampleResetWithContext() {
	out, _ := git.ResetWithContext(context.Background(), reset.Soft, reset.Commit("e41f083"), git.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git reset --soft e41f083
}

func ExampleCommit() {
	out, _ := git.Commit(commit.Amend, commit.Message("foo"), git.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git commit --amend --message=foo
}

func ExampleCommitWithContext() {
	out, _ := git.CommitWithContext(context.Background(), commit.Amend, commit.Message("foo"), git.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git commit --amend --message=foo
}

func ExampleAdd() {
	out, _ := git.Add(add.All, git.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git add --all
}

func ExampleAddWithContext() {
	out, _ := git.AddWithContext(context.Background(), add.All, git.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git add --all
}

func ExampleMerge() {
	out, _ := git.Merge(merge.Squash, merge.Commits("myBranch"), git.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git merge --squash myBranch
}

func ExampleMergeWithContext() {
	out, _ := git.MergeWithContext(context.Background(), merge.Squash, merge.Commits("myBranch"), git.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git merge --squash myBranch
}

func ExampleWorktree() {
	out, _ := git.Worktree(worktree.Add("v1.0", "origin/v1.0"), git.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git worktree add v1.0 origin/v1.0
}

func ExampleWorktreeWithContext() {
	out, _ := git.WorktreeWithContext(context.Background(), worktree.Add("v1.0", "origin/v1.0"), git.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git worktree add v1.0 origin/v1.0
}

func ExampleTag() {
	out, _ := git.Tag(tag.List, git.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git tag --list
}

func ExampleTagWithContext() {
	out, _ := git.TagWithContext(context.Background(), tag.List, git.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git tag --list
}

func ExampleStatus() {
	out, _ := git.Status(status.Short, status.Branch, git.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git status --short --branch
}

func ExampleStatusWithContext() {
	out, _ := git.StatusWithContext(context.Background(), status.Short, status.Branch, git.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git status --short --branch
}

func ExampleRaw() {
	out, _ := git.Raw("stash", git.CmdExecutor(cmdExecutorMock), func(g *types.Cmd) {
		g.AddOptions("list")
		g.AddOptions("--pretty=format:'%Cblue%gd%Creset%Cred:%Creset %C(yellow)%s%Creset'")
	})

	fmt.Println(out)
	// Output: git stash list --pretty=format:'%Cblue%gd%Creset%Cred:%Creset %C(yellow)%s%Creset'
}

func ExampleRawWithContext() {
	out, _ := git.RawWithContext(context.Background(), "stash", git.CmdExecutor(cmdExecutorMock), func(g *types.Cmd) {
		g.AddOptions("list")
		g.AddOptions("--pretty=format:'%Cblue%gd%Creset%Cred:%Creset %C(yellow)%s%Creset'")
	})

	fmt.Println(out)
	// Output: git stash list --pretty=format:'%Cblue%gd%Creset%Cred:%Creset %C(yellow)%s%Creset'
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

func cmdExecutorMock(_ context.Context, name string, _ bool, args ...string) (string, error) {
	return fmt.Sprintln(name, strings.Join(args, " ")), nil
}
