# Contributing

## Add a new command

I recommend opening an issue before creating a PR.

1. You must create a directory with the name of the command.
2. Inside this directory, you must create a `doc.go` with a summary of the git command (based of the help of this command).
3. You must define all the options manually inside [descriptions.json](https://github.com/ldez/go-git-cmd-wrapper/blob/main/internal/descriptions.json).
4. Optionally, you could create a `<command>_custom.go` inside the command directory to add option(s) which cannot be generated.
5. You must run `make generate` to generate the `<command>_gen.go` related to the `descriptions.json`.
6. Inside the command directory, you must add test `<command>_test.go`.
7. You must add examples inside [example_test.go](https://github.com/ldez/go-git-cmd-wrapper/blob/main/git/example_test.go)

Additionally, I use [regular expressions](https://github.com/ldez/go-git-cmd-wrapper/blob/main/internal/readme.adoc) to help with the `descriptions.json`,
but the extracted information from the doc or the help is not useable directly,
manual editions are required.
