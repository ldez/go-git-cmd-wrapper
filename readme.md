# Go Git Cmd Wrapper

[![Build Status](https://github.com/ldez/go-git-cmd-wrapper/workflows/Main/badge.svg?branch=master)](https://github.com/ldez/go-git-cmd-wrapper/actions)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/ldez/go-git-cmd-wrapper)](https://pkg.go.dev/github.com/ldez/go-git-cmd-wrapper/v2)
[![Go Report Card](https://goreportcard.com/badge/github.com/ldez/go-git-cmd-wrapper)](https://goreportcard.com/report/github.com/ldez/go-git-cmd-wrapper)

[![Sponsor](https://img.shields.io/badge/Sponsor%20me-%E2%9D%A4%EF%B8%8F-pink.svg)](https://github.com/sponsors/ldez)

It's a simple wrapper around `git` command.

Import `github.com/ldez/go-git-cmd-wrapper/v2/git`.

```go
// clone
output, err := git.Clone(clone.Repository("https://github.com/ldez/prm"))
// with debug option
output, err := git.Clone(clone.Repository("https://github.com/ldez/prm"), git.Debug)
output, err := git.Clone(clone.Repository("https://github.com/ldez/prm"), git.Debugger(true))

// fetch
output, err = git.Fetch(fetch.NoTags, fetch.Remote("upstream"))
output, err = git.Fetch(fetch.NoTags, fetch.Remote("upstream"), fetch.RefSpec("master"))

// add a remote
output, err = git.Remote(remote.Add, remote.Name("upstream"), remote.URL("https://github.com/ldez/prm"))
```

More examples: [Documentation](https://pkg.go.dev/github.com/ldez/go-git-cmd-wrapper/v2/git)
