---
title: Helm Client
---

Troubleshootings
===

How to downgrade helm client version using Homebrew
---

If helm versions are incompatible with servers, then helm cli can't be used.
See [this issue](https://github.com/helm/helm/issues/4547#issuecomment-423312200) for the details of issues and solutions.
To sum up, following commands should be ran to downgrade the version to 2.11.0
In order to downgrade the helm cli version, we have to use
```
brew unlink kubernetes-helm
brew install https://raw.githubusercontent.com/Homebrew/homebrew-core/9a698c885051f99b629513cf38684675df2109f9/Formula/kubernetes-helm.rb
brew switch kubernetes-helm 2.11.0
```

This is the example for outputs
```
> helm version
Client: &version.Version{SemVer:"v2.14.3", GitCommit:"0e7f3b6637f7af8fcfddb3d2941fcc7cbebb0085", GitTreeState:"clean"}
Server: &version.Version{SemVer:"v2.11.0", GitCommit:"2e55dbe1fdb5fdb96b75ff144a339489417b146b", GitTreeState:"clean"}
> brew unlink kubernetes-helm
Unlinking /usr/local/Cellar/kubernetes-helm/2.14.3... 47 symlinks removed
> brew install https://raw.githubusercontent.com/Homebrew/homebrew-core/ee94af74778e48ae103a9fb080e26a6a2f62d32c/Formula/kubernetes-helm.rb
==> Consider using `brew extract kubernetes-helm ...`!
This will extract your desired kubernetes-helm version to a stable tap instead of
installing from an unstable URL!

######################################################################## 100.0%
Warning: kubernetes-helm 2.14.3 is available and more recent than version 2.11.0.
==> Downloading https://homebrew.bintray.com/bottles/kubernetes-helm-2.11.0.mojave.bottle.1.tar.gz
==> Downloading from https://akamai.bintray.com/74/74bfbc75ed551ba51124d1b088f45df642a55d9d9fdef45d796d690f70c1f10e?
######################################################################## 100.0%
==> Pouring kubernetes-helm-2.11.0.mojave.bottle.1.tar.gz
==> Caveats
Bash completion has been installed to:
  /usr/local/etc/bash_completion.d

zsh completions have been installed to:
  /usr/local/share/zsh/site-functions
==> Summary
🍺  /usr/local/Cellar/kubernetes-helm/2.11.0: 51 files, 78.5MB
Removing: /Users/at-ishikawa/Library/Caches/Homebrew/kubernetes-helm--2.11.0.mojave.bottle.1.tar.gz... (22.5MB)
Removing: /Users/at-ishikawa/Library/Caches/Homebrew/kubernetes-helm--2.11.0.high_sierra.bottle.tar.gz... (22.5MB)
> brew switch kubernetes-helm 2.11.0
Cleaning /usr/local/Cellar/kubernetes-helm/2.14.3
Cleaning /usr/local/Cellar/kubernetes-helm/2.11.0
47 links created for /usr/local/Cellar/kubernetes-helm/2.11.0
> helm ls
NAME                            REVISION        UPDATED                         STATUS          CHART                                       APP VERSION     NAMESPACE
cert-manager                    29              Wed Oct  2 18:39:34 2019        DEPLOYED        cert-manager-v0.10.1                        v0.10.1         cert-manager
```
