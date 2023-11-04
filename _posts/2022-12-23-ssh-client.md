---
date: 2022-12-23
title: SSH client
tags:
  - ssh
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


Trouble shootings
===

When you run SSH, you get an error `Pseudo-terminal will not be allocated because stdin is not a terminal`.
---

I wanted to get one time token from 1 password and pass it to SSH for the server's 2FA.
I tried to run next for this, but failed
```
> echo (op item get --otp "server" --vault "vault") | ssh server
Pseudo-terminal will not be allocated because stdin is not a terminal.
```

So I ended up running like next
```
> op item get --otp "server" --vault "vault" && ssh server
```

See [this article](https://linuxtutorials.org/Pseudo-terminal-will-not-be-allocated-because-stdin-is-not-a-terminal/) for example.
