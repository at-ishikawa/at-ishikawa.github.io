---
title: Gitconfig
date: '2020-03-21'
# last_modified_at:
categories:
  - git
tags:
  - git
---
Written in March 2020.

Configuration
===
The detail for gitconfig is written in [official page](https://git-scm.com/docs/git-config).

Conditional includes
---
`includeIf` section can be used to include another git configuration under certain conditions.
* gitdir: If repositories are under a specific directory, glob pattern can be used. Note that the path of this should end with `/`.
* gitdir/i: The same as `gitdir` but case-insensitive matching,
* onbranch: If current branch matches the glob pattern of `onbranch`.

For example, there are two files `~/.gitconfig` and `~/.gitconfig-group`.

* `~/.gitconfig`
```conf
[user]
	email = "default@gmail.com"
[includeIf "gitdir:~/group/"]
	path = .gitconfig-group
```

* `~/.gitconfig-group`
```
[user]
	email = "member@group.com"
```

Then if a user work on a repository under `~/group` directory, his/her email becomes `example@group.com`, but out of the directory, it becomes `example@default.com`.

Use cases
===

Use different email for private and professional cases
---

If you use separate emails for private and jobs, then you wanna switch them automatically and wanna avoid configure emails on each repository.
You can do this using `includeIf`, and put all of private repositories or job related repositories into one specific subdirectory.
See "Conditional includes" for more details.

### Reference
- [DZone: How to Use .gitconfig's includeIf](https://dzone.com/articles/how-to-use-gitconfigs-includeif)
