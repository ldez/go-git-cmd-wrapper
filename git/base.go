package git

import (
	"context"
	"io"
	"log"
	"slices"

	"github.com/ldez/go-git-cmd-wrapper/v2/types"
)

// Init https://git-scm.com/docs/git-init
func Init(options ...types.Option) (string, error) {
	return command(context.Background(), "init", options...)
}

// InitWithContext https://git-scm.com/docs/git-init
func InitWithContext(ctx context.Context, options ...types.Option) (string, error) {
	return command(ctx, "init", options...)
}

// Push https://git-scm.com/docs/git-push
func Push(options ...types.Option) (string, error) {
	return command(context.Background(), "push", options...)
}

// PushWithContext https://git-scm.com/docs/git-push
func PushWithContext(ctx context.Context, options ...types.Option) (string, error) {
	return command(ctx, "push", options...)
}

// Pull https://git-scm.com/docs/git-pull
func Pull(options ...types.Option) (string, error) {
	return command(context.Background(), "pull", options...)
}

// PullWithContext https://git-scm.com/docs/git-pull
func PullWithContext(ctx context.Context, options ...types.Option) (string, error) {
	return command(ctx, "pull", options...)
}

// Clone https://git-scm.com/docs/git-clone
func Clone(options ...types.Option) (string, error) {
	return command(context.Background(), "clone", options...)
}

// CloneWithContext https://git-scm.com/docs/git-clone
func CloneWithContext(ctx context.Context, options ...types.Option) (string, error) {
	return command(ctx, "clone", options...)
}

// Remote https://git-scm.com/docs/git-remote
func Remote(options ...types.Option) (string, error) {
	return command(context.Background(), "remote", options...)
}

// RemoteWithContext https://git-scm.com/docs/git-remote
func RemoteWithContext(ctx context.Context, options ...types.Option) (string, error) {
	return command(ctx, "remote", options...)
}

// Fetch https://git-scm.com/docs/git-fetch
func Fetch(options ...types.Option) (string, error) {
	return command(context.Background(), "fetch", options...)
}

// FetchWithContext https://git-scm.com/docs/git-fetch
func FetchWithContext(ctx context.Context, options ...types.Option) (string, error) {
	return command(ctx, "fetch", options...)
}

// Rebase https://git-scm.com/docs/git-rebase
func Rebase(options ...types.Option) (string, error) {
	return command(context.Background(), "rebase", options...)
}

// RebaseWithContext https://git-scm.com/docs/git-rebase
func RebaseWithContext(ctx context.Context, options ...types.Option) (string, error) {
	return command(ctx, "rebase", options...)
}

// Checkout https://git-scm.com/docs/git-checkout
func Checkout(options ...types.Option) (string, error) {
	return command(context.Background(), "checkout", options...)
}

// CheckoutWithContext https://git-scm.com/docs/git-checkout
func CheckoutWithContext(ctx context.Context, options ...types.Option) (string, error) {
	return command(ctx, "checkout", options...)
}

// Config https://git-scm.com/docs/git-config
func Config(options ...types.Option) (string, error) {
	return command(context.Background(), "config", options...)
}

// ConfigWithContext https://git-scm.com/docs/git-config
func ConfigWithContext(ctx context.Context, options ...types.Option) (string, error) {
	return command(ctx, "config", options...)
}

// Branch https://git-scm.com/docs/git-branch
func Branch(options ...types.Option) (string, error) {
	return command(context.Background(), "branch", options...)
}

// BranchWithContext https://git-scm.com/docs/git-branch
func BranchWithContext(ctx context.Context, options ...types.Option) (string, error) {
	return command(ctx, "branch", options...)
}

// RevParse https://git-scm.com/docs/git-rev-parse
func RevParse(options ...types.Option) (string, error) {
	return command(context.Background(), "rev-parse", options...)
}

// RevParseWithContext https://git-scm.com/docs/git-rev-parse
func RevParseWithContext(ctx context.Context, options ...types.Option) (string, error) {
	return command(ctx, "rev-parse", options...)
}

// Reset https://git-scm.com/docs/git-reset
func Reset(options ...types.Option) (string, error) {
	return command(context.Background(), "reset", options...)
}

// ResetWithContext https://git-scm.com/docs/git-reset
func ResetWithContext(ctx context.Context, options ...types.Option) (string, error) {
	return command(ctx, "reset", options...)
}

// Commit https://git-scm.com/docs/git-commit
func Commit(options ...types.Option) (string, error) {
	return command(context.Background(), "commit", options...)
}

// CommitWithContext https://git-scm.com/docs/git-commit
func CommitWithContext(ctx context.Context, options ...types.Option) (string, error) {
	return command(ctx, "commit", options...)
}

// Add https://git-scm.com/docs/git-add
func Add(options ...types.Option) (string, error) {
	return command(context.Background(), "add", options...)
}

// AddWithContext https://git-scm.com/docs/git-add
func AddWithContext(ctx context.Context, options ...types.Option) (string, error) {
	return command(ctx, "add", options...)
}

// Tag https://git-scm.com/docs/git-tag
func Tag(options ...types.Option) (string, error) {
	return command(context.Background(), "tag", options...)
}

// TagWithContext https://git-scm.com/docs/git-tag
func TagWithContext(ctx context.Context, options ...types.Option) (string, error) {
	return command(ctx, "tag", options...)
}

// Merge https://git-scm.com/docs/git-merge
func Merge(options ...types.Option) (string, error) {
	return command(context.Background(), "merge", options...)
}

// MergeWithContext https://git-scm.com/docs/git-merge
func MergeWithContext(ctx context.Context, options ...types.Option) (string, error) {
	return command(ctx, "merge", options...)
}

// Worktree https://git-scm.com/docs/git-worktree
func Worktree(options ...types.Option) (string, error) {
	return command(context.Background(), "worktree", options...)
}

// WorktreeWithContext https://git-scm.com/docs/git-worktree
func WorktreeWithContext(ctx context.Context, options ...types.Option) (string, error) {
	return command(ctx, "worktree", options...)
}

// Status https://git-scm.com/docs/git-status
func Status(options ...types.Option) (string, error) {
	return command(context.Background(), "status", options...)
}

// StatusWithContext https://git-scm.com/docs/git-status
func StatusWithContext(ctx context.Context, options ...types.Option) (string, error) {
	return command(ctx, "status", options...)
}

// Notes https://git-scm.com/docs/git-notes
func Notes(subCommand ...types.Option) (string, error) {
	return command(context.Background(), "notes", subCommand...)
}

// NotesWithContext https://git-scm.com/docs/git-notes
func NotesWithContext(ctx context.Context, subCommand ...types.Option) (string, error) {
	return command(ctx, "notes", subCommand...)
}

// LsFiles https://git-scm.com/docs/git-ls-files
func LsFiles(subCommand ...types.Option) (string, error) {
	return command(context.Background(), "ls-files", subCommand...)
}

// LsFilesWithContext https://git-scm.com/docs/git-ls-files
func LsFilesWithContext(ctx context.Context, subCommand ...types.Option) (string, error) {
	return command(ctx, "ls-files", subCommand...)
}

// Stash https://git-scm.com/docs/git-stash
func Stash(options ...types.Option) (string, error) {
	return command(context.Background(), "stash", options...)
}

// StashWithContext https://git-scm.com/docs/git-stash
func StashWithContext(ctx context.Context, options ...types.Option) (string, error) {
	return command(ctx, "stash", options...)
}

// Raw use to execute arbitrary git commands.
func Raw(cmd string, options ...types.Option) (string, error) {
	return command(context.Background(), cmd, options...)
}

// RawWithContext use to execute arbitrary git commands.
func RawWithContext(ctx context.Context, cmd string, options ...types.Option) (string, error) {
	return command(ctx, cmd, options...)
}

// Debug display command line.
func Debug(g *types.Cmd) {
	g.Debug = true
}

// Debugger display command line.
func Debugger(debug bool) types.Option {
	return func(g *types.Cmd) {
		g.Debug = debug
	}
}

// Cond apply conditionally some options.
func Cond(apply bool, options ...types.Option) types.Option {
	if apply {
		return func(g *types.Cmd) {
			g.ApplyOptions(options...)
		}
	}
	return NoOp
}

// NoOp do nothing.
func NoOp(_ *types.Cmd) {}

// LogOutput Writer used by the internal logger.
func LogOutput(w io.Writer) types.Option {
	return func(g *types.Cmd) {
		g.Logger = log.New(w, "", 0)
	}
}

// CmdExecutor Allow to override the Git command call (useful for testing purpose).
func CmdExecutor(executor types.Executor) types.Option {
	return func(g *types.Cmd) {
		g.Executor = executor
	}
}

func command(ctx context.Context, name string, options ...types.Option) (string, error) {
	g := types.NewCmd(name)
	g.ApplyOptions(options...)

	return g.Exec(ctx, g.Base, g.Debug, slices.Concat(g.BaseOptions, g.Options)...)
}
