---
date: 2021-07-24
title: IntelliJ plugin development
tags:
  - intellij
---

Getting started
===

Set up a project
---

See [this page](https://plugins.jetbrains.com/docs/intellij/getting-started.html) for the details.

1. Create a new repository from the [template repository](https://github.com/JetBrains/intellij-platform-plugin-template)
2. Update a repository to follow README
   - Update pluginGroup, pluginId, and plugin name
3. Update a repository for your purpose
4. Publish a plugin manually at first
   - It requires a review from JetBrains and need to wait for a few days

Development a plugin
---

### Run a plugin in an IntelliJ IDE

See an [official page](https://plugins.jetbrains.com/docs/intellij/getting-started.html) for more details.

* Download Intellij IDEA Community Edition or IntelliJ IDEA Ultimate.
* Open the project for a plugin in the IntelliJ IDEA
* Run "Run Plugin" and new window for IntelliJ will start with the plugin.


### Install a plugin in an IntelliJ IDE

* Run `./gradlew buildPlugin`
* Open **Preferences > Plugins > Install Plugin from Disk...** and choose a zip file under `build/distributions`



Troubleshootings
===


runPluginVerifier fails due to UnsupportedClassVersionError
---

Just after I created a repository from the template, `./gradlew runPluginVerifier -Pplugin.verifier.home.dir=~/.pluginVerifier` on one of GitHub Actions failed because of `Error: Exception in thread "main" java.lang.UnsupportedClassVersionError: com/jetbrains/pluginverifier/PluginVerifierMain has been compiled by a more recent version of the Java Runtime (class file version 55.0), this version of the Java Runtime only recognizes class file versions up to 52.0`.

The detail error was following.

```
2021-07-25T00:13:10.7212635Z Found online and idle hosted runner(s) in the current repository's organization account that matches the required labels: 'ubuntu-latest'
2021-07-25T00:13:10.7212690Z Waiting for a hosted runner in 'organization' to pick this job...
2021-07-25T00:13:15.4957097Z Current runner version: '2.278.0'
2021-07-25T00:13:15.4991647Z ##[group]Operating System
2021-07-25T00:13:15.4992624Z Ubuntu
2021-07-25T00:13:15.4993092Z 20.04.2
2021-07-25T00:13:15.4993625Z LTS
2021-07-25T00:13:15.4994112Z ##[endgroup]
2021-07-25T00:13:15.4994754Z ##[group]Virtual Environment
2021-07-25T00:13:15.4995415Z Environment: ubuntu-20.04
2021-07-25T00:13:15.4996039Z Version: 20210718.1
2021-07-25T00:13:15.4997087Z Included Software: https://github.com/actions/virtual-environments/blob/ubuntu20/20210718.1/images/linux/Ubuntu2004-README.md
2021-07-25T00:13:15.4998845Z Image Release: https://github.com/actions/virtual-environments/releases/tag/ubuntu20%2F20210718.1
2021-07-25T00:13:15.4999878Z ##[endgroup]
2021-07-25T00:13:15.5002065Z ##[group]GITHUB_TOKEN Permissions
2021-07-25T00:13:15.5003664Z Actions: write
2021-07-25T00:13:15.5004333Z Checks: write
2021-07-25T00:13:15.5004903Z Contents: write
2021-07-25T00:13:15.5005588Z Deployments: write
2021-07-25T00:13:15.5006685Z Discussions: write
2021-07-25T00:13:15.5007334Z Issues: write
2021-07-25T00:13:15.5007860Z Metadata: read
2021-07-25T00:13:15.5008443Z Packages: write
2021-07-25T00:13:15.5009016Z PullRequests: write
2021-07-25T00:13:15.5009933Z RepositoryProjects: write
2021-07-25T00:13:15.5010653Z SecurityEvents: write
2021-07-25T00:13:15.5011364Z Statuses: write
2021-07-25T00:13:15.5012097Z ##[endgroup]
2021-07-25T00:13:15.5015363Z Prepare workflow directory
2021-07-25T00:13:15.5740740Z Prepare all required actions
2021-07-25T00:13:15.5756036Z Getting action download info
2021-07-25T00:13:15.9051505Z Download action repository 'actions/setup-java@v2'
2021-07-25T00:13:18.0103462Z Download action repository 'actions/checkout@v2.3.4'
2021-07-25T00:13:18.1914942Z Download action repository 'actions/cache@v2.1.6'
2021-07-25T00:13:18.5806686Z ##[group]Run actions/setup-java@v2
2021-07-25T00:13:18.5807472Z with:
2021-07-25T00:13:18.5807896Z   distribution: zulu
2021-07-25T00:13:18.5808419Z   java-version: 8
2021-07-25T00:13:18.5808891Z   java-package: jdk
2021-07-25T00:13:18.5809395Z   architecture: x64
2021-07-25T00:13:18.5809879Z   check-latest: false
2021-07-25T00:13:18.5810369Z   server-id: github
2021-07-25T00:13:18.5810907Z   server-username: GITHUB_ACTOR
2021-07-25T00:13:18.5811531Z   server-password: GITHUB_TOKEN
2021-07-25T00:13:18.5812144Z   overwrite-settings: true
2021-07-25T00:13:18.5812669Z ##[endgroup]
2021-07-25T00:13:19.5112544Z Trying to resolve the latest version from remote
2021-07-25T00:13:19.5224076Z Resolved latest version as 8.0.302+8
2021-07-25T00:13:19.5224667Z Trying to download...
2021-07-25T00:13:19.5226273Z Downloading Java 8.0.302+8 (Zulu) from https://cdn.azul.com/zulu/bin/zulu8.56.0.21-ca-jdk8.0.302-linux_x64.tar.gz ...
2021-07-25T00:13:19.6799111Z Extracting Java archive...
2021-07-25T00:13:19.6910272Z [command]/usr/bin/tar xz --warning=no-unknown-keyword -C /home/runner/work/_temp/3c59a625-a806-47db-b926-ac9f1e1b7cd8 -f /home/runner/work/_temp/3f056c09-d9ab-4a0d-b97b-2ce3cfdf7612
2021-07-25T00:13:22.0822405Z Java 8.0.302+8 was downloaded
2021-07-25T00:13:22.0824351Z Setting Java 8.0.302+8 as the default
2021-07-25T00:13:22.0857126Z
2021-07-25T00:13:22.0858309Z Java configuration:
2021-07-25T00:13:22.0859034Z   Distribution: zulu
2021-07-25T00:13:22.0859721Z   Version: 8.0.302+8
2021-07-25T00:13:22.0860766Z   Path: /opt/hostedtoolcache/Java_Zulu_jdk/8.0.302-8/x64
2021-07-25T00:13:22.0861460Z
2021-07-25T00:13:22.0970186Z Creating settings.xml with server-id: github
2021-07-25T00:13:22.0970949Z Writing to /home/runner/.m2/settings.xml
2021-07-25T00:13:22.1189706Z ##[group]Run actions/checkout@v2.3.4
2021-07-25T00:13:22.1190573Z with:
2021-07-25T00:13:22.1191989Z   repository: at-ishikawa/intellij-plugin-emacs-macos-keymap
2021-07-25T00:13:22.1194405Z   token: ***
2021-07-25T00:13:22.1195134Z   ssh-strict: true
2021-07-25T00:13:22.1196112Z   persist-credentials: true
2021-07-25T00:13:22.1197032Z   clean: true
2021-07-25T00:13:22.1197762Z   fetch-depth: 1
2021-07-25T00:13:22.1198501Z   lfs: false
2021-07-25T00:13:22.1199227Z   submodules: false
2021-07-25T00:13:22.1200209Z env:
2021-07-25T00:13:22.1201165Z   JAVA_HOME: /opt/hostedtoolcache/Java_Zulu_jdk/8.0.302-8/x64
2021-07-25T00:13:22.1202192Z ##[endgroup]
2021-07-25T00:13:22.2295169Z Syncing repository: at-ishikawa/intellij-plugin-emacs-macos-keymap
2021-07-25T00:13:22.2301548Z ##[group]Getting Git version info
2021-07-25T00:13:22.2303557Z Working directory is '/home/runner/work/intellij-plugin-emacs-macos-keymap/intellij-plugin-emacs-macos-keymap'
2021-07-25T00:13:22.2360852Z [command]/usr/bin/git version
2021-07-25T00:13:22.2496497Z git version 2.32.0
2021-07-25T00:13:22.2518819Z ##[endgroup]
2021-07-25T00:13:22.2527379Z Deleting the contents of '/home/runner/work/intellij-plugin-emacs-macos-keymap/intellij-plugin-emacs-macos-keymap'
2021-07-25T00:13:22.2530364Z ##[group]Initializing the repository
2021-07-25T00:13:22.2535663Z [command]/usr/bin/git init /home/runner/work/intellij-plugin-emacs-macos-keymap/intellij-plugin-emacs-macos-keymap
2021-07-25T00:13:22.2803367Z hint: Using 'master' as the name for the initial branch. This default branch name
2021-07-25T00:13:22.2805796Z hint: is subject to change. To configure the initial branch name to use in all
2021-07-25T00:13:22.2808987Z hint: of your new repositories, which will suppress this warning, call:
2021-07-25T00:13:22.2809706Z hint:
2021-07-25T00:13:22.2810776Z hint: 	git config --global init.defaultBranch <name>
2021-07-25T00:13:22.2811443Z hint:
2021-07-25T00:13:22.2813832Z hint: Names commonly chosen instead of 'master' are 'main', 'trunk' and
2021-07-25T00:13:22.2815676Z hint: 'development'. The just-created branch can be renamed via this command:
2021-07-25T00:13:22.2816367Z hint:
2021-07-25T00:13:22.2816938Z hint: 	git branch -m <name>
2021-07-25T00:13:22.2822684Z Initialized empty Git repository in /home/runner/work/intellij-plugin-emacs-macos-keymap/intellij-plugin-emacs-macos-keymap/.git/
2021-07-25T00:13:22.2837737Z [command]/usr/bin/git remote add origin https://github.com/at-ishikawa/intellij-plugin-emacs-macos-keymap
2021-07-25T00:13:22.2871968Z ##[endgroup]
2021-07-25T00:13:22.2873020Z ##[group]Disabling automatic garbage collection
2021-07-25T00:13:22.2880921Z [command]/usr/bin/git config --local gc.auto 0
2021-07-25T00:13:22.2976801Z ##[endgroup]
2021-07-25T00:13:22.2981963Z ##[group]Setting up auth
2021-07-25T00:13:22.2989554Z [command]/usr/bin/git config --local --name-only --get-regexp core\.sshCommand
2021-07-25T00:13:22.3024798Z [command]/usr/bin/git submodule foreach --recursive git config --local --name-only --get-regexp 'core\.sshCommand' && git config --local --unset-all 'core.sshCommand' || :
2021-07-25T00:13:22.5058135Z [command]/usr/bin/git config --local --name-only --get-regexp http\.https\:\/\/github\.com\/\.extraheader
2021-07-25T00:13:22.5100810Z [command]/usr/bin/git submodule foreach --recursive git config --local --name-only --get-regexp 'http\.https\:\/\/github\.com\/\.extraheader' && git config --local --unset-all 'http.https://github.com/.extraheader' || :
2021-07-25T00:13:22.5363390Z [command]/usr/bin/git config --local http.https://github.com/.extraheader AUTHORIZATION: basic ***
2021-07-25T00:13:22.5433218Z ##[endgroup]
2021-07-25T00:13:22.5434307Z ##[group]Fetching the repository
2021-07-25T00:13:22.5436667Z [command]/usr/bin/git -c protocol.version=2 fetch --no-tags --prune --progress --no-recurse-submodules --depth=1 origin +dc0662f8078541a2f3a61cd08dfa8531b2cce474:refs/remotes/origin/main
2021-07-25T00:13:22.8876797Z remote: Enumerating objects: 32, done.
2021-07-25T00:13:22.8878128Z remote: Counting objects:   3% (1/32)
2021-07-25T00:13:22.8878868Z remote: Counting objects:   6% (2/32)
2021-07-25T00:13:22.8879596Z remote: Counting objects:   9% (3/32)
2021-07-25T00:13:22.8880450Z remote: Counting objects:  12% (4/32)
2021-07-25T00:13:22.8881211Z remote: Counting objects:  15% (5/32)
2021-07-25T00:13:22.8881880Z remote: Counting objects:  18% (6/32)
2021-07-25T00:13:22.8882624Z remote: Counting objects:  21% (7/32)
2021-07-25T00:13:22.8883464Z remote: Counting objects:  25% (8/32)
2021-07-25T00:13:22.8884144Z remote: Counting objects:  28% (9/32)
2021-07-25T00:13:22.8884819Z remote: Counting objects:  31% (10/32)
2021-07-25T00:13:22.8885449Z remote: Counting objects:  34% (11/32)
2021-07-25T00:13:22.8886106Z remote: Counting objects:  37% (12/32)
2021-07-25T00:13:22.8893630Z remote: Counting objects:  40% (13/32)
2021-07-25T00:13:22.8894340Z remote: Counting objects:  43% (14/32)
2021-07-25T00:13:22.8895163Z remote: Counting objects:  46% (15/32)
2021-07-25T00:13:22.8895787Z remote: Counting objects:  50% (16/32)
2021-07-25T00:13:22.8896502Z remote: Counting objects:  53% (17/32)
2021-07-25T00:13:22.8897096Z remote: Counting objects:  56% (18/32)
2021-07-25T00:13:22.8897681Z remote: Counting objects:  59% (19/32)
2021-07-25T00:13:22.8898247Z remote: Counting objects:  62% (20/32)
2021-07-25T00:13:22.8898822Z remote: Counting objects:  65% (21/32)
2021-07-25T00:13:22.8899429Z remote: Counting objects:  68% (22/32)
2021-07-25T00:13:22.8900017Z remote: Counting objects:  71% (23/32)
2021-07-25T00:13:22.8900578Z remote: Counting objects:  75% (24/32)
2021-07-25T00:13:22.8901120Z remote: Counting objects:  78% (25/32)
2021-07-25T00:13:22.8901677Z remote: Counting objects:  81% (26/32)
2021-07-25T00:13:22.8902223Z remote: Counting objects:  84% (27/32)
2021-07-25T00:13:22.8902788Z remote: Counting objects:  87% (28/32)
2021-07-25T00:13:22.8903336Z remote: Counting objects:  90% (29/32)
2021-07-25T00:13:22.8903896Z remote: Counting objects:  93% (30/32)
2021-07-25T00:13:22.8904453Z remote: Counting objects:  96% (31/32)
2021-07-25T00:13:22.8904994Z remote: Counting objects: 100% (32/32)
2021-07-25T00:13:22.8907273Z remote: Counting objects: 100% (32/32), done.
2021-07-25T00:13:22.8907943Z remote: Compressing objects:   3% (1/27)
2021-07-25T00:13:22.8909000Z remote: Compressing objects:   7% (2/27)
2021-07-25T00:13:22.8909639Z remote: Compressing objects:  11% (3/27)
2021-07-25T00:13:22.8910249Z remote: Compressing objects:  14% (4/27)
2021-07-25T00:13:22.8910849Z remote: Compressing objects:  18% (5/27)
2021-07-25T00:13:22.8911462Z remote: Compressing objects:  22% (6/27)
2021-07-25T00:13:22.8912074Z remote: Compressing objects:  25% (7/27)
2021-07-25T00:13:22.8913302Z remote: Compressing objects:  29% (8/27)
2021-07-25T00:13:22.8929577Z remote: Compressing objects:  33% (9/27)
2021-07-25T00:13:22.8930419Z remote: Compressing objects:  37% (10/27)
2021-07-25T00:13:22.8931225Z remote: Compressing objects:  40% (11/27)
2021-07-25T00:13:22.8932014Z remote: Compressing objects:  44% (12/27)
2021-07-25T00:13:22.8936248Z remote: Compressing objects:  48% (13/27)
2021-07-25T00:13:22.8937062Z remote: Compressing objects:  51% (14/27)
2021-07-25T00:13:22.8937871Z remote: Compressing objects:  55% (15/27)
2021-07-25T00:13:22.8938672Z remote: Compressing objects:  59% (16/27)
2021-07-25T00:13:22.8939452Z remote: Compressing objects:  62% (17/27)
2021-07-25T00:13:22.8940250Z remote: Compressing objects:  66% (18/27)
2021-07-25T00:13:22.8941038Z remote: Compressing objects:  70% (19/27)
2021-07-25T00:13:22.8941839Z remote: Compressing objects:  74% (20/27)
2021-07-25T00:13:22.8942846Z remote: Compressing objects:  77% (21/27)
2021-07-25T00:13:22.8943639Z remote: Compressing objects:  81% (22/27)
2021-07-25T00:13:22.8944433Z remote: Compressing objects:  85% (23/27)
2021-07-25T00:13:22.8945218Z remote: Compressing objects:  88% (24/27)
2021-07-25T00:13:22.8946018Z remote: Compressing objects:  92% (25/27)
2021-07-25T00:13:22.8946805Z remote: Compressing objects:  96% (26/27)
2021-07-25T00:13:22.8947605Z remote: Compressing objects: 100% (27/27)
2021-07-25T00:13:22.8948526Z remote: Compressing objects: 100% (27/27), done.
2021-07-25T00:13:22.9035809Z remote: Total 32 (delta 2), reused 17 (delta 1), pack-reused 0
2021-07-25T00:13:22.9411371Z From https://github.com/at-ishikawa/intellij-plugin-emacs-macos-keymap
2021-07-25T00:13:22.9413035Z  * [new ref]         dc0662f8078541a2f3a61cd08dfa8531b2cce474 -> origin/main
2021-07-25T00:13:22.9516541Z ##[endgroup]
2021-07-25T00:13:22.9522394Z ##[group]Determining the checkout info
2021-07-25T00:13:22.9553891Z ##[endgroup]
2021-07-25T00:13:22.9554538Z ##[group]Checking out the ref
2021-07-25T00:13:22.9555939Z [command]/usr/bin/git checkout --progress --force -B main refs/remotes/origin/main
2021-07-25T00:13:22.9601461Z Switched to a new branch 'main'
2021-07-25T00:13:22.9602362Z Branch 'main' set up to track remote branch 'main' from 'origin'.
2021-07-25T00:13:22.9618276Z ##[endgroup]
2021-07-25T00:13:22.9667643Z [command]/usr/bin/git log -1 --format='%H'
2021-07-25T00:13:22.9690884Z 'dc0662f8078541a2f3a61cd08dfa8531b2cce474'
2021-07-25T00:13:23.1040075Z ##[group]Run actions/cache@v2.1.6
2021-07-25T00:13:23.1040860Z with:
2021-07-25T00:13:23.1041284Z   path: ~/.gradle/caches
2021-07-25T00:13:23.1042963Z   key: Linux-gradle-caches-7304178c36b767d53707fc7caf65d771a8e1fd3786ce624ff9932cb11afb7619
2021-07-25T00:13:23.1044604Z env:
2021-07-25T00:13:23.1045144Z   JAVA_HOME: /opt/hostedtoolcache/Java_Zulu_jdk/8.0.302-8/x64
2021-07-25T00:13:23.1045723Z ##[endgroup]
2021-07-25T00:13:24.5004444Z Received 197132288 of 1371603264 (14.4%), 187.8 MBs/sec
2021-07-25T00:13:25.5014472Z Received 373293056 of 1371603264 (27.2%), 177.8 MBs/sec
2021-07-25T00:13:26.5051465Z Received 583008256 of 1371603264 (42.5%), 185.0 MBs/sec
2021-07-25T00:13:27.5054517Z Received 830472192 of 1371603264 (60.5%), 197.7 MBs/sec
2021-07-25T00:13:28.5093722Z Received 1090519040 of 1371603264 (79.5%), 207.6 MBs/sec
2021-07-25T00:13:29.5104184Z Received 1359020352 of 1371603264 (99.1%), 215.6 MBs/sec
2021-07-25T00:13:34.9629484Z Received 1371603264 of 1371603264 (100.0%), 114.1 MBs/sec
2021-07-25T00:13:34.9637576Z Cache Size: ~1308 MB (1371603264 B)
2021-07-25T00:13:34.9756337Z [command]/usr/bin/tar --use-compress-program zstd -d -xf /home/runner/work/_temp/4870a2d4-6da8-4359-890b-132937208416/cache.tzst -P -C /home/runner/work/intellij-plugin-emacs-macos-keymap/intellij-plugin-emacs-macos-keymap
2021-07-25T00:13:45.5792669Z Cache restored successfully
2021-07-25T00:13:46.0001682Z Cache restored from key: Linux-gradle-caches-7304178c36b767d53707fc7caf65d771a8e1fd3786ce624ff9932cb11afb7619
2021-07-25T00:13:46.5835893Z ##[group]Run actions/cache@v2.1.6
2021-07-25T00:13:46.5836443Z with:
2021-07-25T00:13:46.5836985Z   path: ~/.gradle/wrapper
2021-07-25T00:13:46.5838850Z   key: Linux-gradle-wrapper-10a37fa6138018db999d3b97bd06c81cc04dc5b55a168498f1dc6ff7fac5b8b2
2021-07-25T00:13:46.5840722Z env:
2021-07-25T00:13:46.5841373Z   JAVA_HOME: /opt/hostedtoolcache/Java_Zulu_jdk/8.0.302-8/x64
2021-07-25T00:13:46.5842003Z ##[endgroup]
2021-07-25T00:13:47.9531056Z Received 206568275 of 223345491 (92.5%), 197.0 MBs/sec
2021-07-25T00:13:48.1938412Z Received 223345491 of 223345491 (100.0%), 171.6 MBs/sec
2021-07-25T00:13:48.2052250Z Cache Size: ~213 MB (223345491 B)
2021-07-25T00:13:48.2054916Z [command]/usr/bin/tar --use-compress-program zstd -d -xf /home/runner/work/_temp/c3aa8d53-69e2-4b17-b3ae-8347db12a96c/cache.tzst -P -C /home/runner/work/intellij-plugin-emacs-macos-keymap/intellij-plugin-emacs-macos-keymap
2021-07-25T00:13:48.6421323Z Cache restored successfully
2021-07-25T00:13:48.6771954Z Cache restored from key: Linux-gradle-wrapper-10a37fa6138018db999d3b97bd06c81cc04dc5b55a168498f1dc6ff7fac5b8b2
2021-07-25T00:13:48.6994743Z ##[group]Run PROPERTIES="$(./gradlew properties --console=plain -q)"
2021-07-25T00:13:48.6995744Z [36;1mPROPERTIES="$(./gradlew properties --console=plain -q)"[0m
2021-07-25T00:13:48.6996783Z [36;1mIDE_VERSIONS="$(echo "$PROPERTIES" | grep "^pluginVerifierIdeVersions:" | base64)"[0m
2021-07-25T00:13:48.6997587Z [36;1m[0m
2021-07-25T00:13:48.6998300Z [36;1mecho "::set-output name=ideVersions::$IDE_VERSIONS"[0m
2021-07-25T00:13:48.6999256Z [36;1mecho "::set-output name=pluginVerifierHomeDir::~/.pluginVerifier"[0m
2021-07-25T00:13:48.7043537Z shell: /usr/bin/bash --noprofile --norc -e -o pipefail {0}
2021-07-25T00:13:48.7044111Z env:
2021-07-25T00:13:48.7044706Z   JAVA_HOME: /opt/hostedtoolcache/Java_Zulu_jdk/8.0.302-8/x64
2021-07-25T00:13:48.7045312Z ##[endgroup]
2021-07-25T00:13:57.3305329Z ##[group]Run actions/cache@v2.1.6
2021-07-25T00:13:57.3305811Z with:
2021-07-25T00:13:57.3306305Z   path: ~/.pluginVerifier/ides
2021-07-25T00:13:57.3309282Z   key: Linux-plugin-verifier-cGx1Z2luVmVyaWZpZXJJZGVWZXJzaW9uczogMjAyMC4yLjQsIDIwMjAuMy40LCAyMDIxLjEuMQo=
2021-07-25T00:13:57.3312100Z env:
2021-07-25T00:13:57.3312669Z   JAVA_HOME: /opt/hostedtoolcache/Java_Zulu_jdk/8.0.302-8/x64
2021-07-25T00:13:57.3313258Z ##[endgroup]
2021-07-25T00:13:57.5464507Z Cache not found for input keys: Linux-plugin-verifier-cGx1Z2luVmVyaWZpZXJJZGVWZXJzaW9uczogMjAyMC4yLjQsIDIwMjAuMy40LCAyMDIxLjEuMQo=
2021-07-25T00:13:57.5547355Z ##[group]Run ./gradlew runPluginVerifier -Pplugin.verifier.home.dir=~/.pluginVerifier
2021-07-25T00:13:57.5548782Z [36;1m./gradlew runPluginVerifier -Pplugin.verifier.home.dir=~/.pluginVerifier[0m
2021-07-25T00:13:57.5591421Z shell: /usr/bin/bash -e {0}
2021-07-25T00:13:57.5591857Z env:
2021-07-25T00:13:57.5592444Z   JAVA_HOME: /opt/hostedtoolcache/Java_Zulu_jdk/8.0.302-8/x64
2021-07-25T00:13:57.5593054Z ##[endgroup]
2021-07-25T00:13:58.2047895Z
2021-07-25T00:13:58.2049151Z Welcome to Gradle 7.0.2!
2021-07-25T00:13:58.2049510Z
2021-07-25T00:13:58.2050020Z Here are the highlights of this release:
2021-07-25T00:13:58.2051289Z  - File system watching enabled by default
2021-07-25T00:13:58.2052191Z  - Support for running with and building Java 16 projects
2021-07-25T00:13:58.2053079Z  - Native support for Apple Silicon processors
2021-07-25T00:13:58.2053905Z  - Dependency catalog feature preview
2021-07-25T00:13:58.2054289Z
2021-07-25T00:13:58.2055420Z For more details see https://docs.gradle.org/7.0.2/release-notes.html
2021-07-25T00:13:58.2056060Z
2021-07-25T00:13:59.1151550Z > Task :compileKotlin NO-SOURCE
2021-07-25T00:13:59.1163065Z > Task :compileJava NO-SOURCE
2021-07-25T00:13:59.2024485Z > Task :patchPluginXml
2021-07-25T00:13:59.3018010Z > Task :processResources
2021-07-25T00:13:59.3019032Z > Task :classes
2021-07-25T00:13:59.6020675Z > Task :instrumentCode
2021-07-25T00:13:59.6021811Z > Task :postInstrumentCode
2021-07-25T00:13:59.7024513Z > Task :inspectClassesForKotlinIC
2021-07-25T00:13:59.7025937Z > Task :jar
2021-07-25T00:13:59.7026645Z > Task :prepareSandbox
2021-07-25T00:14:15.0018474Z
2021-07-25T00:14:15.0027957Z WARNING: An illegal reflective access operation has occurred
2021-07-25T00:14:15.0028794Z > Task :buildSearchableOptions
2021-07-25T00:14:15.0030321Z WARNING: Illegal reflective access by com.intellij.util.ReflectionUtil to field java.awt.event.InvocationEvent.runnable
2021-07-25T00:14:15.0131142Z WARNING: Please consider reporting this to the maintainers of com.intellij.util.ReflectionUtil
2021-07-25T00:14:15.0133281Z WARNING: Use --illegal-access=warn to enable warnings of further illegal reflective access operations
2021-07-25T00:14:15.0134352Z WARNING: All illegal access operations will be denied in a future release
2021-07-25T00:14:16.6015894Z 2021-07-25 00:14:16,577 [   2014]   WARN - j.internal.DebugAttachDetector - Unable to start DebugAttachDetector, please add `--add-exports java.base/jdk.internal.vm=ALL-UNNAMED` to VM options
2021-07-25T00:14:17.6018752Z Starting searchable options index builder
2021-07-25T00:14:18.0016297Z 2021-07-25 00:14:17,909 [   3346]   WARN - ConfigurableExtensionPointUtil - ignore deprecated groupId: language for id: preferences.language.Kotlin.scripting
2021-07-25T00:14:38.6016111Z Jul 25, 2021 12:14:38 AM java.util.prefs.FileSystemPreferences$1 run
2021-07-25T00:14:38.6040712Z INFO: Created user preferences directory.
2021-07-25T00:14:38.6042045Z Jul 25, 2021 12:14:38 AM java.util.prefs.FileSystemPreferences$6 run
2021-07-25T00:14:38.6043875Z WARNING: Prefs file removed in background /home/runner/.java/.userPrefs/prefs.xml
2021-07-25T00:14:42.1014366Z Searchable options index builder completed
2021-07-25T00:14:42.9015376Z
2021-07-25T00:14:42.9016964Z > Task :jarSearchableOptions
2021-07-25T00:14:42.9017606Z > Task :buildPlugin
2021-07-25T00:14:43.2014173Z > Task :verifyPlugin
2021-07-25T00:15:13.2016750Z > Task :runPluginVerifier
2021-07-25T00:16:20.8014226Z Error: A JNI error has occurred, please check your installation and try again
2021-07-25T00:16:20.8024432Z ##[error]Exception in thread "main" java.lang.UnsupportedClassVersionError: com/jetbrains/pluginverifier/PluginVerifierMain has been compiled by a more recent version of the Java Runtime (class file version 55.0), this version of the Java Runtime only recognizes class file versions up to 52.0
2021-07-25T00:16:20.8039111Z
2021-07-25T00:16:20.8039993Z > Task :runPluginVerifier FAILED
2021-07-25T00:16:20.8044291Z 	at java.lang.ClassLoader.defineClass1(Native Method)
2021-07-25T00:16:20.8045580Z 	at java.lang.ClassLoader.defineClass(ClassLoader.java:757)
2021-07-25T00:16:20.8047392Z 	at java.security.SecureClassLoader.defineClass(SecureClassLoader.java:142)
2021-07-25T00:16:20.8049048Z 	at java.net.URLClassLoader.defineClass(URLClassLoader.java:468)
2021-07-25T00:16:20.8050400Z 	at java.net.URLClassLoader.access$100(URLClassLoader.java:74)
2021-07-25T00:16:20.8051583Z 	at java.net.URLClassLoader$1.run(URLClassLoader.java:369)
2021-07-25T00:16:20.8052641Z 	at java.net.URLClassLoader$1.run(URLClassLoader.java:363)
2021-07-25T00:16:20.8053946Z 	at java.security.AccessController.doPrivileged(Native Method)
2021-07-25T00:16:20.8055394Z 	at java.net.URLClassLoader.findClass(URLClassLoader.java:362)
2021-07-25T00:16:20.8056654Z 	at java.lang.ClassLoader.loadClass(ClassLoader.java:419)
2021-07-25T00:16:20.8057842Z 	at sun.misc.Launcher$AppClassLoader.loadClass(Launcher.java:352)
2021-07-25T00:16:20.8066521Z 	at java.lang.ClassLoader.loadClass(ClassLoader.java:352)
2021-07-25T00:16:20.8068112Z 	at sun.launcher.LauncherHelper.checkAndLoadMain(LauncherHelper.java:601)
2021-07-25T00:16:20.9015192Z
2021-07-25T00:16:20.9016434Z FAILURE: Build failed with an exception.
2021-07-25T00:16:20.9016981Z
2021-07-25T00:16:20.9018192Z * What went wrong:
2021-07-25T00:16:20.9019624Z Execution failed for task ':runPluginVerifier'.
2021-07-25T00:16:20.9020913Z > Process 'command '/opt/hostedtoolcache/Java_Zulu_jdk/8.0.302-8/x64/bin/java'' finished with non-zero exit value 1
2021-07-25T00:16:20.9021588Z
2021-07-25T00:16:20.9022090Z 12 actionable tasks: 12 executed
2021-07-25T00:16:20.9022606Z * Try:
2021-07-25T00:16:20.9023755Z Run with --stacktrace option to get the stack trace. Run with --info or --debug option to get more log output. Run with --scan to get full insights.
2021-07-25T00:16:20.9024502Z
2021-07-25T00:16:20.9025122Z * Get more help at https://help.gradle.org
2021-07-25T00:16:20.9025608Z
2021-07-25T00:16:20.9026025Z BUILD FAILED in 2m 23s
2021-07-25T00:16:20.9347047Z ##[error]Process completed with exit code 1.
2021-07-25T00:16:20.9448432Z Post job cleanup.
2021-07-25T00:16:21.0705168Z [command]/usr/bin/git version
2021-07-25T00:16:21.0757956Z git version 2.32.0
2021-07-25T00:16:21.0796802Z [command]/usr/bin/git config --local --name-only --get-regexp core\.sshCommand
2021-07-25T00:16:21.0832485Z [command]/usr/bin/git submodule foreach --recursive git config --local --name-only --get-regexp 'core\.sshCommand' && git config --local --unset-all 'core.sshCommand' || :
2021-07-25T00:16:21.1198262Z [command]/usr/bin/git config --local --name-only --get-regexp http\.https\:\/\/github\.com\/\.extraheader
2021-07-25T00:16:21.1239967Z http.https://github.com/.extraheader
2021-07-25T00:16:21.1241564Z [command]/usr/bin/git config --local --unset-all http.https://github.com/.extraheader
2021-07-25T00:16:21.1272678Z [command]/usr/bin/git submodule foreach --recursive git config --local --name-only --get-regexp 'http\.https\:\/\/github\.com\/\.extraheader' && git config --local --unset-all 'http.https://github.com/.extraheader' || :
2021-07-25T00:16:21.1630135Z Post job cleanup.
2021-07-25T00:16:21.2242240Z Cleaning up orphan processes
2021-07-25T00:16:21.2527423Z Terminate orphan process: pid (1722) (java)
```


### How to solve it
I haven't solved it yet.



Errors when Run "Run Plugin" on IntelliJ IDEA
---

I got following warning messages when I run
- a "Run Plugin" in IntelliJ IDE, or
- build a plugin by `./gradlew buildPlugin`

```
2021-07-24 17:37:13,157 [    526]   WARN - .intellij.util.EnvironmentUtil - can't get shell environment
java.lang.RuntimeException: command [/usr/local/bin/fish, -l, -c, '$HOME/.gradle/caches/modules-2/files-2.1/com.jetbrains.intellij.idea/ideaIC/2020.2.4/b1d028673e39dbecbf4aa2a2bcd4f2a72f8860da/ideaIC-2020.2.4/bin/printenv.py' '/var/folders/hz/7rcdbkks63n31jjb763824y00000gp/T/intellij-shell-env.5285664788779010860.tmp']
	exit code:126 text:0 out:fish: The file '$HOME/.gradle/caches/modules-2/files-2.1/com.jetbrains.intellij.idea/ideaIC/2020.2.4/b1d028673e39dbecbf4aa2a2bcd4f2a72f8860da/ideaIC-2020.2.4/bin/printenv.py' is not executable by this user
'$HOME/.gradle/caches/modules-2/files-2.1/com.jetbrains.intellij.idea/ideaIC/2020.2.4/b1d028673e39dbecbf4aa2a2bcd4f2a72f8860da/ideaIC-2020.2.4/bin/printenv.py' '/var/folders/hz/7rcdbkks63n31jjb763824y00000gp/T/intellij-shell-env.5285664788779010860.tmp'
^
	at com.intellij.util.EnvironmentUtil$ShellEnvReader.runProcessAndReadOutputAndEnvs(EnvironmentUtil.java:292)
	at com.intellij.util.EnvironmentUtil$ShellEnvReader.readShellEnv(EnvironmentUtil.java:220)
	at com.intellij.util.EnvironmentUtil.getShellEnv(EnvironmentUtil.java:180)
	at com.intellij.util.EnvironmentUtil.lambda$loadEnvironment$0(EnvironmentUtil.java:71)
	at java.base/java.util.concurrent.CompletableFuture$AsyncSupply.run(CompletableFuture.java:1700)
	at java.base/java.util.concurrent.ThreadPoolExecutor.runWorker(ThreadPoolExecutor.java:1128)
	at java.base/java.util.concurrent.ThreadPoolExecutor$Worker.run(ThreadPoolExecutor.java:628)
	at java.base/java.util.concurrent.Executors$PrivilegedThreadFactory$1$1.run(Executors.java:668)
	at java.base/java.util.concurrent.Executors$PrivilegedThreadFactory$1$1.run(Executors.java:665)
	at java.base/java.security.AccessController.doPrivileged(Native Method)
	at java.base/java.util.concurrent.Executors$PrivilegedThreadFactory$1.run(Executors.java:665)
	at java.base/java.lang.Thread.run(Thread.java:834)
2021-07-24 17:37:13,844 [   1213]   WARN - i.mac.MacOSApplicationProvider - No URL bundle (CFBundleURLTypes) is defined in the main bundle.
To be able to open external links, specify protocols in the app layout section of the build file.
Example: args.urlSchemes = ["your-protocol"] will handle following links: your-protocol://open?file=file&line=line
2021-07-24 17:37:15,575 [   2944]   WARN - j.internal.DebugAttachDetector - Unable to start DebugAttachDetector, please add `--add-exports java.base/jdk.internal.vm=ALL-UNNAMED` to VM options
2021-07-24 17:37:16,785 [   4154]   WARN - nSystem.impl.ActionManagerImpl - keymap "Xcode" not found [Plugin: com.intellij]
2021-07-24 17:37:38,083 [  25452]   WARN - api.vfs.impl.local.FileWatcher - Native file watcher is not executable: <a href="$HOME/.gradle/caches/modules-2/files-2.1/com.jetbrains.intellij.idea/ideaIC/2020.2.4/b1d028673e39dbecbf4aa2a2bcd4f2a72f8860da/ideaIC-2020.2.4/bin/mac/fsnotifier">$HOME/.gradle/caches/modules-2/files-2.1/com.jetbrains.intellij.idea/ideaIC/2020.2.4/b1d028673e39dbecbf4aa2a2bcd4f2a72f8860da/ideaIC-2020.2.4/bin/mac/fsnotifier</a>
2021-07-24 17:37:38,941 [  26310]   WARN - ion.impl.NotificationCollector - Notification group 'Heap Dump Analysis' is already registered in whitelist
2021-07-24 17:37:38,941 [  26310]   WARN - ion.impl.NotificationCollector - Notification group 'Low Memory' is already registered in whitelist
2021-07-24 17:38:28,034 [  75403]   WARN - Container.ComponentManagerImpl - Do not use constructor injection (requestorClass=org.jetbrains.android.compose.AndroidComposeAutoDocumentation)
2021-07-24 17:38:28,385 [  75754]   WARN - tartup.impl.StartupManagerImpl - Activities registered via registerPostStartupActivity must be dumb-aware: com.intellij.ui.mac.touchbar.TouchBarsManager$$Lambda$1824/0x000000080142b440@394f1533
2021-07-24 17:38:28,396 [  75765]   WARN - tartup.impl.StartupManagerImpl - Activities registered via registerPostStartupActivity must be dumb-aware: org.jetbrains.kotlin.idea.configuration.ui.KotlinConfigurationCheckerComponent$projectOpened$1@51531918
2021-07-24 17:38:34,173 [  81542]   WARN - .intellij.diagnostic.VMOptions - VM options file not configured
2021-07-24 17:38:38,878 [  86247]   WARN - ConfigurableExtensionPointUtil - ignore deprecated groupId: language for id: preferences.language.Kotlin.scripting
2021-07-24 17:38:52,395 [  99764]   WARN - com.intellij.util.xmlb.Binding - no accessors for org.jetbrains.kotlin.idea.core.script.configuration.utils.ScriptClassRootsStorage
2021-07-24 17:39:02,407 [ 109776]   WARN - com.intellij.util.xmlb.Binding - no accessors for org.jetbrains.kotlin.idea.highlighter.KotlinDefaultHighlightingSettingsProvider
2021-07-24 17:39:26,991 [ 134360]   WARN - com.intellij.util.xmlb.Binding - no accessors for org.jetbrains.kotlin.idea.scripting.gradle.GradleScriptInputsWatcher$Storage
2021-07-24 17:39:30,975 [ 138344]  ERROR - tellij.openapi.util.ObjectTree - Memory leak detected: 'com.intellij.facet.impl.ui.libraries.LibraryOptionsPanel@e76f452' of class com.intellij.facet.impl.ui.libraries.LibraryOptionsPanel
See the cause for the corresponding Disposer.register() stacktrace:

java.lang.RuntimeException: Memory leak detected: 'com.intellij.facet.impl.ui.libraries.LibraryOptionsPanel@e76f452' of class com.intellij.facet.impl.ui.libraries.LibraryOptionsPanel
See the cause for the corresponding Disposer.register() stacktrace:

	at com.intellij.openapi.util.ObjectTree.assertIsEmpty(ObjectTree.java:235)
	at com.intellij.openapi.util.Disposer.assertIsEmpty(Disposer.java:144)
	at com.intellij.openapi.util.Disposer.assertIsEmpty(Disposer.java:139)
	at com.intellij.openapi.application.impl.ApplicationImpl.disposeContainer(ApplicationImpl.java:197)
	at com.intellij.openapi.application.impl.ApplicationImpl.disposeSelf(ApplicationImpl.java:214)
	at com.intellij.openapi.application.impl.ApplicationImpl.doExit(ApplicationImpl.java:609)
	at com.intellij.openapi.application.impl.ApplicationImpl.exit(ApplicationImpl.java:579)
	at com.intellij.openapi.application.impl.ApplicationImpl.exit(ApplicationImpl.java:568)
	at com.intellij.openapi.application.ex.ApplicationEx.exit(ApplicationEx.java:101)
	at com.intellij.openapi.wm.impl.welcomeScreen.WelcomeFrame$3.windowClosing(WelcomeFrame.java:111)
	at java.desktop/java.awt.AWTEventMulticaster.windowClosing(AWTEventMulticaster.java:357)
	at java.desktop/java.awt.Window.processWindowEvent(Window.java:2079)
	at java.desktop/javax.swing.JFrame.processWindowEvent(JFrame.java:298)
	at java.desktop/java.awt.Window.processEvent(Window.java:2038)
	at java.desktop/java.awt.Component.dispatchEventImpl(Component.java:5029)
	at java.desktop/java.awt.Container.dispatchEventImpl(Container.java:2321)
	at java.desktop/java.awt.Window.dispatchEventImpl(Window.java:2773)
	at java.desktop/java.awt.Component.dispatchEvent(Component.java:4861)
	at java.desktop/java.awt.EventQueue.dispatchEventImpl(EventQueue.java:778)
	at java.desktop/java.awt.EventQueue$4.run(EventQueue.java:727)
	at java.desktop/java.awt.EventQueue$4.run(EventQueue.java:721)
	at java.base/java.security.AccessController.doPrivileged(Native Method)
	at java.base/java.security.ProtectionDomain$JavaSecurityAccessImpl.doIntersectionPrivilege(ProtectionDomain.java:85)
	at java.base/java.security.ProtectionDomain$JavaSecurityAccessImpl.doIntersectionPrivilege(ProtectionDomain.java:95)
	at java.desktop/java.awt.EventQueue$5.run(EventQueue.java:751)
	at java.desktop/java.awt.EventQueue$5.run(EventQueue.java:749)
	at java.base/java.security.AccessController.doPrivileged(Native Method)
	at java.base/java.security.ProtectionDomain$JavaSecurityAccessImpl.doIntersectionPrivilege(ProtectionDomain.java:85)
	at java.desktop/java.awt.EventQueue.dispatchEvent(EventQueue.java:748)
	at com.intellij.ide.IdeEventQueue.defaultDispatchEvent(IdeEventQueue.java:971)
	at com.intellij.ide.IdeEventQueue._dispatchEvent(IdeEventQueue.java:841)
	at com.intellij.ide.IdeEventQueue.lambda$dispatchEvent$8(IdeEventQueue.java:452)
	at com.intellij.openapi.progress.impl.CoreProgressManager.computePrioritized(CoreProgressManager.java:744)
	at com.intellij.ide.IdeEventQueue.lambda$dispatchEvent$9(IdeEventQueue.java:451)
	at com.intellij.openapi.application.impl.ApplicationImpl.runIntendedWriteActionOnCurrentThread(ApplicationImpl.java:802)
	at com.intellij.ide.IdeEventQueue.dispatchEvent(IdeEventQueue.java:505)
	at java.desktop/java.awt.EventDispatchThread.pumpOneEventForFilters(EventDispatchThread.java:203)
	at java.desktop/java.awt.EventDispatchThread.pumpEventsForFilter(EventDispatchThread.java:124)
	at java.desktop/java.awt.EventDispatchThread.pumpEventsForHierarchy(EventDispatchThread.java:113)
	at java.desktop/java.awt.EventDispatchThread.pumpEvents(EventDispatchThread.java:109)
	at java.desktop/java.awt.EventDispatchThread.pumpEvents(EventDispatchThread.java:101)
	at java.desktop/java.awt.EventDispatchThread.run(EventDispatchThread.java:90)
Caused by: java.lang.Throwable
	at com.intellij.openapi.util.ObjectNode.<init>(ObjectNode.java:35)
	at com.intellij.openapi.util.ObjectTree.createNodeFor(ObjectTree.java:113)
	at com.intellij.openapi.util.ObjectTree.register(ObjectTree.java:74)
	at com.intellij.openapi.util.Disposer.register(Disposer.java:70)
	at com.intellij.facet.impl.ui.libraries.LibraryOptionsPanel.showSettingsPanel(LibraryOptionsPanel.java:180)
	at com.intellij.facet.impl.ui.libraries.LibraryOptionsPanel.<init>(LibraryOptionsPanel.java:127)
	at com.intellij.facet.impl.ui.libraries.LibraryOptionsPanel.<init>(LibraryOptionsPanel.java:100)
	at org.jetbrains.kotlin.idea.framework.KotlinModuleSettingStep.getLibraryPanel(KotlinModuleSettingStep.java:186)
	at org.jetbrains.kotlin.idea.framework.KotlinModuleSettingStep.getComponent(KotlinModuleSettingStep.java:137)
	at org.jetbrains.kotlin.idea.framework.KotlinModuleSettingStep.<init>(KotlinModuleSettingStep.java:103)
	at org.jetbrains.kotlin.idea.framework.KotlinModuleBuilder.modifySettingsStep(KotlinModuleBuilder.kt:40)
	at com.intellij.ide.projectWizard.ProjectSettingsStep.setupPanels(ProjectSettingsStep.java:90)
	at com.intellij.ide.projectWizard.ProjectSettingsStep.updateStep(ProjectSettingsStep.java:109)
	at com.intellij.ide.util.newProjectWizard.AbstractProjectWizard.updateStep(AbstractProjectWizard.java:149)
	at com.intellij.ide.wizard.AbstractWizard.updateStep(AbstractWizard.java:384)
	at com.intellij.ide.wizard.AbstractWizard.doNextAction(AbstractWizard.java:408)
	at com.intellij.ide.util.newProjectWizard.AbstractProjectWizard.doNextAction(AbstractProjectWizard.java:253)
	at com.intellij.ide.wizard.AbstractWizard.proceedToNextStep(AbstractWizard.java:220)
	at com.intellij.ide.wizard.AbstractWizard$5.actionPerformed(AbstractWizard.java:177)
	at java.desktop/javax.swing.AbstractButton.fireActionPerformed(AbstractButton.java:1967)
	at java.desktop/javax.swing.AbstractButton$Handler.actionPerformed(AbstractButton.java:2308)
	at java.desktop/javax.swing.DefaultButtonModel.fireActionPerformed(DefaultButtonModel.java:405)
	at java.desktop/javax.swing.DefaultButtonModel.setPressed(DefaultButtonModel.java:262)
	at java.desktop/javax.swing.plaf.basic.BasicButtonListener.mouseReleased(BasicButtonListener.java:270)
	at java.desktop/java.awt.Component.processMouseEvent(Component.java:6654)
	at java.desktop/javax.swing.JComponent.processMouseEvent(JComponent.java:3345)
	at java.desktop/java.awt.Component.processEvent(Component.java:6419)
	at java.desktop/java.awt.Container.processEvent(Container.java:2263)
	at java.desktop/java.awt.Component.dispatchEventImpl(Component.java:5029)
	at java.desktop/java.awt.Container.dispatchEventImpl(Container.java:2321)
	at java.desktop/java.awt.Component.dispatchEvent(Component.java:4861)
	at java.desktop/java.awt.LightweightDispatcher.retargetMouseEvent(Container.java:4918)
	at java.desktop/java.awt.LightweightDispatcher.processMouseEvent(Container.java:4547)
	at java.desktop/java.awt.LightweightDispatcher.dispatchEvent(Container.java:4488)
	at java.desktop/java.awt.Container.dispatchEventImpl(Container.java:2307)
	at java.desktop/java.awt.Window.dispatchEventImpl(Window.java:2773)
	at java.desktop/java.awt.Component.dispatchEvent(Component.java:4861)
	at java.desktop/java.awt.EventQueue.dispatchEventImpl(EventQueue.java:778)
	at java.desktop/java.awt.EventQueue$4.run(EventQueue.java:727)
	at java.desktop/java.awt.EventQueue$4.run(EventQueue.java:721)
	at java.base/java.security.AccessController.doPrivileged(Native Method)
	at java.base/java.security.ProtectionDomain$JavaSecurityAccessImpl.doIntersectionPrivilege(ProtectionDomain.java:85)
	at java.base/java.security.ProtectionDomain$JavaSecurityAccessImpl.doIntersectionPrivilege(ProtectionDomain.java:95)
	at java.desktop/java.awt.EventQueue$5.run(EventQueue.java:751)
	at java.desktop/java.awt.EventQueue$5.run(EventQueue.java:749)
	at java.base/java.security.AccessController.doPrivileged(Native Method)
	at java.base/java.security.ProtectionDomain$JavaSecurityAccessImpl.doIntersectionPrivilege(ProtectionDomain.java:85)
	at java.desktop/java.awt.EventQueue.dispatchEvent(EventQueue.java:748)
	at com.intellij.ide.IdeEventQueue.defaultDispatchEvent(IdeEventQueue.java:971)
	at com.intellij.ide.IdeEventQueue.dispatchMouseEvent(IdeEventQueue.java:906)
	at com.intellij.ide.IdeEventQueue._dispatchEvent(IdeEventQueue.java:838)
	at com.intellij.ide.IdeEventQueue.lambda$dispatchEvent$8(IdeEventQueue.java:452)
	at com.intellij.openapi.progress.impl.CoreProgressManager.computePrioritized(CoreProgressManager.java:733)
	at com.intellij.ide.IdeEventQueue.lambda$dispatchEvent$9(IdeEventQueue.java:451)
	at com.intellij.openapi.application.impl.ApplicationImpl.runIntendedWriteActionOnCurrentThread(ApplicationImpl.java:802)
	at com.intellij.ide.IdeEventQueue.dispatchEvent(IdeEventQueue.java:505)
	at java.desktop/java.awt.EventDispatchThread.pumpOneEventForFilters(EventDispatchThread.java:203)
	at java.desktop/java.awt.EventDispatchThread.pumpEventsForFilter(EventDispatchThread.java:124)
	at java.desktop/java.awt.EventDispatchThread.pumpEventsForFilter(EventDispatchThread.java:117)
	at java.desktop/java.awt.WaitDispatchSupport$2.run(WaitDispatchSupport.java:190)
	at java.desktop/java.awt.WaitDispatchSupport$4.run(WaitDispatchSupport.java:235)
	at java.desktop/java.awt.WaitDispatchSupport$4.run(WaitDispatchSupport.java:233)
	at java.base/java.security.AccessController.doPrivileged(Native Method)
	at java.desktop/java.awt.WaitDispatchSupport.enter(WaitDispatchSupport.java:233)
	at java.desktop/java.awt.Dialog.show(Dialog.java:1063)
	at com.intellij.openapi.ui.impl.DialogWrapperPeerImpl$MyDialog.show(DialogWrapperPeerImpl.java:711)
	at com.intellij.openapi.ui.impl.DialogWrapperPeerImpl.show(DialogWrapperPeerImpl.java:438)
	at com.intellij.openapi.ui.DialogWrapper.doShow(DialogWrapper.java:1700)
	at com.intellij.openapi.ui.DialogWrapper.show(DialogWrapper.java:1659)
	at com.intellij.openapi.ui.DialogWrapper.showAndGet(DialogWrapper.java:1673)
	at com.intellij.ide.impl.NewProjectUtil.createNewProject(NewProjectUtil.java:63)
	at com.intellij.ide.actions.NewProjectAction.actionPerformed(NewProjectAction.java:24)
	at com.intellij.openapi.actionSystem.ex.ActionUtil.performActionDumbAware(ActionUtil.java:282)
	at com.intellij.openapi.actionSystem.ex.ActionUtil.invokeAction(ActionUtil.java:446)
	at com.intellij.openapi.actionSystem.ex.ActionUtil.invokeAction(ActionUtil.java:431)
	at com.intellij.ui.components.labels.ActionLink$1.linkSelected(ActionLink.java:47)
	at com.intellij.ui.components.labels.LinkLabel.doClick(LinkLabel.java:138)
	at com.intellij.ui.components.labels.ActionLink.doClick(ActionLink.java:56)
	at com.intellij.ui.components.labels.LinkLabel$MyMouseHandler.mouseReleased(LinkLabel.java:322)
	at java.desktop/java.awt.Component.processMouseEvent(Component.java:6654)
	at java.desktop/javax.swing.JComponent.processMouseEvent(JComponent.java:3345)
	at java.desktop/java.awt.Component.processEvent(Component.java:6419)
	at java.desktop/java.awt.Container.processEvent(Container.java:2263)
	at java.desktop/java.awt.Component.dispatchEventImpl(Component.java:5029)
	at java.desktop/java.awt.Container.dispatchEventImpl(Container.java:2321)
	at java.desktop/java.awt.Component.dispatchEvent(Component.java:4861)
	at java.desktop/java.awt.LightweightDispatcher.retargetMouseEvent(Container.java:4918)
	at java.desktop/java.awt.LightweightDispatcher.processMouseEvent(Container.java:4547)
	at java.desktop/java.awt.LightweightDispatcher.dispatchEvent(Container.java:4488)
	at java.desktop/java.awt.Container.dispatchEventImpl(Container.java:2307)
	at java.desktop/java.awt.Window.dispatchEventImpl(Window.java:2773)
	at java.desktop/java.awt.Component.dispatchEvent(Component.java:4861)
	at java.desktop/java.awt.EventQueue.dispatchEventImpl(EventQueue.java:778)
	at java.desktop/java.awt.EventQueue$4.run(EventQueue.java:727)
	at java.desktop/java.awt.EventQueue$4.run(EventQueue.java:721)
	at java.base/java.security.AccessController.doPrivileged(Native Method)
	at java.base/java.security.ProtectionDomain$JavaSecurityAccessImpl.doIntersectionPrivilege(ProtectionDomain.java:85)
	at java.base/java.security.ProtectionDomain$JavaSecurityAccessImpl.doIntersectionPrivilege(ProtectionDomain.java:95)
	at java.desktop/java.awt.EventQueue$5.run(EventQueue.java:751)
	at java.desktop/java.awt.EventQueue$5.run(EventQueue.java:749)
	at java.base/java.security.AccessController.doPrivileged(Native Method)
	at java.base/java.security.ProtectionDomain$JavaSecurityAccessImpl.doIntersectionPrivilege(ProtectionDomain.java:85)
	at java.desktop/java.awt.EventQueue.dispatchEvent(EventQueue.java:748)
	at com.intellij.ide.IdeEventQueue.defaultDispatchEvent(IdeEventQueue.java:971)
	at com.intellij.ide.IdeEventQueue.dispatchMouseEvent(IdeEventQueue.java:906)
	at com.intellij.ide.IdeEventQueue._dispatchEvent(IdeEventQueue.java:838)
	... 11 more
2021-07-24 17:39:30,978 [ 138347]  ERROR - tellij.openapi.util.ObjectTree - IntelliJ IDEA 2020.2.4  Build #IC-202.8194.7
2021-07-24 17:39:30,979 [ 138348]  ERROR - tellij.openapi.util.ObjectTree - JDK: 11.0.9; VM: OpenJDK 64-Bit Server VM; Vendor: JetBrains s.r.o.
2021-07-24 17:39:30,979 [ 138348]  ERROR - tellij.openapi.util.ObjectTree - OS: Mac OS X
```

### How to solve it
I haven't solved yet. Maybe this happened because the shell was fish?

Even though I saw these warnings, it didn't stop the task and I was able to do what I wanted to do, so I didn't spend time to look into it.
