---
title: tmux
date: '2022-01-20'
---

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
