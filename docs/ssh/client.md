---
title: SSH client
---

SSH configuration
===

This is the mapping between config file `/.ssh/config` and `ssh` cli options.

| config file field  | cli option  | Description  |
|---|---|---|
| ForwardAgent  | -A or -a  | Enable to use locale machine's keys on remote machines  |


SSH password
===

Store the key and access on remote machine without inputting a password.
---
See [stackoverflow](https://apple.stackexchange.com/questions/48502/how-can-i-permanently-add-my-ssh-private-key-to-keychain-so-it-is-automatically) for details.


How to add password to the private key without a password
---
Use: `ssh-keygen -p -f /path/to/key`. For example, `ssh-keygen -p -f ~/.ssh/id_rsa`

See [stackoverflow](https://stackoverflow.com/questions/3818886/how-do-i-add-a-password-to-an-openssh-private-key-that-was-generated-without-a-p) for more details.
