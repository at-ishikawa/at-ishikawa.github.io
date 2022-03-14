---
title: tmux
date: '2022-01-20'
---

Basic configuration
===

There are a few options to use set.

- `set-option`: an alias of a `set`
- `-s`: server option
- `-g`: global option
- `setw` or `set-window-option`: window option

See [StackExchange: Difference between global, server, session and window options](https://superuser.com/questions/758843/difference-between-global-server-session-and-window-options) for more details about these options and how they're effective.


### Other references
- [Stack overflow: What are the differences between set -g, set -ga and set-option -g in a .tmux.conf file?](https://stackoverflow.com/questions/45017773/what-are-the-differences-between-set-g-set-ga-and-set-option-g-in-a-tmux-co).


Plugin management
===
[TPM](https://github.com/tmux-plugins/tpm) is a plugin to install plugins easily for tmux.

However, in some cases, plugins aren't installed correctly.
As a workaround, they can be installed by running `~/.tmux/plugins/tpm/scripts/install_plugins.sh`, reported by [this issue](https://github.com/tmux-plugins/tpm/issues/193#issuecomment-775598298).


Copy and paste plugin
---

Use a plugin [tmux-yank](https://github.com/tmux-plugins/tmux-yank).
In order to set the shortcut key same as emacs, set @copy_mode_yank to "M-w", like `set -g @copy_mode_yank "M-w"`



Session and window management
===

- To rename the current session, type `Prefix + $` or `Prefix + :` and type `rename-session <new session name>`
- To rename the current window, type `Prefix + ,` or `Prefix + :` and type `rename-window <new window name>`
- To create new session, type `:new` or `:new -s <session name>`
- To search a window, type `Prefix + f`
    - In order to use fuzzy support like fzf, see [this article](https://eioki.eu/2021/01/12/tmux-and-fzf-fuzzy-tmux-session-window-pane-switcher) or [this github wiki](https://github.com/junegunn/fzf/wiki/Examples#tmux) for reference.


Start specific windows automatically
===
[tmuxinator](https://github.com/tmuxinator/tmuxinator) enables to create windows and manage them by configurations

- Install
    ```
    brew install tmuxinator
    ```

- Set up an environment variable
    ```
    set -Ux EDITOR vim
    ```

- Create new project of tmuxinator
    ```
    tmuxinator new $PROJECT_NAME
    ```

- Start a session of the project
    ```
    tmuxinator start $PROJECT_NAME
    ```
