// Rebase https://git-scm.com/docs/git-rebase
// git rebase [-i | --interactive] [options] [--exec <cmd>] [--onto <newbase>] [<upstream> [<branch>]]
// git rebase [-i | --interactive] [options] [--exec <cmd>] [--onto <newbase>] --root [<branch>]
// git rebase --continue | --skip | --abort | --quit | --edit-todo
// --onto <newbase>
// <upstream>
// <branch>
// --continue
// --abort
// --quit
// --keep-empty
// --skip
// --edit-todo
// --merge
// --strategy=<strategy>
// --strategy-option=<strategy-option>
// --gpg-sign[=<keyid>]
// --quiet
// --verbose
// --stat
// --no-stat
// --no-verify
// --verify
// -C<n>
// --force-rebase
// --fork-point
// --no-fork-point
// --ignore-whitespace
// --whitespace=<option>
// --committer-date-is-author-date
// --ignore-date
// --signoff
// --interactive
// --preserve-merges
// -x <cmd>
// --exec <cmd>
// --root
// --autosquash
// --no-autosquash
// --autostash
// --no-autostash
// --no-ff
package rebase

import "github.com/ldez/go-git-cmd-wrapper/types"

func Continue(g *types.Cmd) {
	g.AddOptions("--continue")
}

func Abort(g *types.Cmd) {
	g.AddOptions("--abort")
}

func Quit(g *types.Cmd) {
	g.AddOptions("--quit")
}

func KeepEmpty(g *types.Cmd) {
	g.AddOptions("--keep-empty")
}

func Skip(g *types.Cmd) {
	g.AddOptions("--skip")
}

func EditTodo(g *types.Cmd) {
	g.AddOptions("--edit-todo")
}

func Merge(g *types.Cmd) {
	g.AddOptions("--merge")
}

func Quiet(g *types.Cmd) {
	g.AddOptions("--quiet")
}

func Verbose(g *types.Cmd) {
	g.AddOptions("--verbose")
}

func Stat(g *types.Cmd) {
	g.AddOptions("--stat")
}

func NoStat(g *types.Cmd) {
	g.AddOptions("--no-stat")
}

func NoVerify(g *types.Cmd) {
	g.AddOptions("--no-verify")
}

func Verify(g *types.Cmd) {
	g.AddOptions("--verify")
}

func ForceRebase(g *types.Cmd) {
	g.AddOptions("--force-rebase")
}

func ForkPoint(g *types.Cmd) {
	g.AddOptions("--fork-point")
}

func NoForkPoint(g *types.Cmd) {
	g.AddOptions("--no-fork-point")
}

func IgnoreWhitespace(g *types.Cmd) {
	g.AddOptions("--ignore-whitespace")
}

func CommitterDateIsAuthorDate(g *types.Cmd) {
	g.AddOptions("--committer-date-is-author-date")
}

func IgnoreDate(g *types.Cmd) {
	g.AddOptions("--ignore-date")
}

func Signoff(g *types.Cmd) {
	g.AddOptions("--signoff")
}

func Interactive(g *types.Cmd) {
	g.AddOptions("--interactive")
}

func PreserveMerges(g *types.Cmd) {
	g.AddOptions("--preserve-merges")
}

func Root(g *types.Cmd) {
	g.AddOptions("--root")
}

func Autosquash(g *types.Cmd) {
	g.AddOptions("--autosquash")
}

func NoAutosquash(g *types.Cmd) {
	g.AddOptions("--no-autosquash")
}

func Autostash(g *types.Cmd) {
	g.AddOptions("--autostash")
}

func NoAutostash(g *types.Cmd) {
	g.AddOptions("--no-autostash")
}

func NoFF(g *types.Cmd) {
	g.AddOptions("--no-ff")
}

func Upstream(name string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(name)
	}
}

func Branch(name string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(name)
	}
}

//--onto <newbase>
//--strategy=<strategy>
//--strategy-option=<strategy-option>
//--gpg-sign[=<keyid>]
//--whitespace=<option>
//-C<n>
//-x <cmd>
//--exec <cmd>
