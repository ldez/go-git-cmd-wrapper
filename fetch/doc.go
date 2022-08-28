/*
Package fetch git-fetch - Download objects and refs from another repository.

# SYNOPSIS

Reference: https://git-scm.com/docs/git-fetch

	git fetch [<options>] [<repository> [<refspec>...]]
	git fetch [<options>] <group>
	git fetch --multiple [<options>] [(<repository> | <group>)...]
	git fetch --all [<options>]

# DESCRIPTION

Fetch branches and/or tags (collectively, "refs") from one or more other repositories, along with the objects necessary to complete their histories.
Remote-tracking branches are updated (see the description of <refspec> below for ways to control this behavior).

By default, any tag that points into the histories being fetched is also fetched; the effect is to fetch tags that point at branches that you are interested in.
This default behavior can be changed by using the --tags or --no-tags options or by configuring remote.<name>.tagOpt.
By using a refspec that fetches tags explicitly, you can fetch tags that do not point into branches you are interested in as well.

git fetch can fetch from either a single named repository or URL, or from several repositories at once if <group> is given and there is a remotes.<group> entry in the configuration file.

When no remote is specified, by default the origin remote will be used, unless there’s an upstream branch configured for the current branch.

The names of refs that are fetched, together with the object names they point at, are written to .git/FETCH_HEAD. This information may be used by scripts or other git commands, such as git-pull(1).
*/
package fetch
