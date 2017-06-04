package rebase

import "github.com/ldez/go-git-cmd-wrapper/git"

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

func Continue(g *git.Cmd) {
	g.AddOptions("--continue")
}

func Abort(g *git.Cmd) {
	g.AddOptions("--abort")
}

func Quit(g *git.Cmd) {
	g.AddOptions("--quit")
}

func KeepEmpty(g *git.Cmd) {
	g.AddOptions("--keep-empty")
}

func Skip(g *git.Cmd) {
	g.AddOptions("--skip")
}

func EditTodo(g *git.Cmd) {
	g.AddOptions("--edit-todo")
}

func Merge(g *git.Cmd) {
	g.AddOptions("--merge")
}

func Quiet(g *git.Cmd) {
	g.AddOptions("--quiet")
}

func Verbose(g *git.Cmd) {
	g.AddOptions("--verbose")
}

func Stat(g *git.Cmd) {
	g.AddOptions("--stat")
}

func NoStat(g *git.Cmd) {
	g.AddOptions("--no-stat")
}

func NoVerify(g *git.Cmd) {
	g.AddOptions("--no-verify")
}

func Verify(g *git.Cmd) {
	g.AddOptions("--verify")
}

func ForceRebase(g *git.Cmd) {
	g.AddOptions("--force-rebase")
}

func ForkPoint(g *git.Cmd) {
	g.AddOptions("--fork-point")
}

func NoForkPoint(g *git.Cmd) {
	g.AddOptions("--no-fork-point")
}

func IgnoreWhitespace(g *git.Cmd) {
	g.AddOptions("--ignore-whitespace")
}

func CommitterDateIsAuthorDate(g *git.Cmd) {
	g.AddOptions("--committer-date-is-author-date")
}

func IgnoreDate(g *git.Cmd) {
	g.AddOptions("--ignore-date")
}

func Signoff(g *git.Cmd) {
	g.AddOptions("--signoff")
}

func Interactive(g *git.Cmd) {
	g.AddOptions("--interactive")
}

func PreserveMerges(g *git.Cmd) {
	g.AddOptions("--preserve-merges")
}

func Root(g *git.Cmd) {
	g.AddOptions("--root")
}

func Autosquash(g *git.Cmd) {
	g.AddOptions("--autosquash")
}

func NoAutosquash(g *git.Cmd) {
	g.AddOptions("--no-autosquash")
}

func Autostash(g *git.Cmd) {
	g.AddOptions("--autostash")
}

func NoAutostash(g *git.Cmd) {
	g.AddOptions("--no-autostash")
}

func NoFF(g *git.Cmd) {
	g.AddOptions("--no-ff")
}

func Upstream(name string) func(*git.Cmd) {
	return func(g *git.Cmd) {
		g.AddOptions(name)
	}
}

func Branch(name string) func(*git.Cmd) {
	return func(g *git.Cmd) {
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
