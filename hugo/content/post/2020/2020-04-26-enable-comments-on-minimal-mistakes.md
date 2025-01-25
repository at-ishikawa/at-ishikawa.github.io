---
date: "2020-04-26T17:00:00Z"
tags:
- git
- jekyll
- github-pages
- minimal-mistakes
title: Enable comments on minimal mistake theme of GitHub pages
---

The configuration to enable comments is described in the [official page](https://mmistakes.github.io/minimal-mistakes/docs/configuration/#comments).

Supported providers
===

There are some providers to support comments, and some of them is not free.
- Disqus: Not free
- Discourse: Not free
- Facebook: Free
- Utterance: Free with GitHub issues.
- Staticman: Free but needs to setup our own instances like on heroku (which is explain in [this issue](https://github.com/eduardoboucas/staticman/issues/343)).


Utterance
===
[utterance](https://utteranc.es/) is easy to setup, and users require GitHub authentication to comment, and it can enable you to get notifications for issues.
Furthermore, because you can manage messages on GitHub, you can manage comments on GitHub features.

Getting Started
---
Configuration is written in the [official page](https://mmistakes.github.io/minimal-mistakes/docs/configuration/#utterances-comments).

Basically, you only need to add next configurations in `_config.yml`.
```
comments:
  provider: "utterances"
  utterances:
    theme: "github-light" # "github-dark"
    issue_term: "pathname"
```

Which issue_term is supported can be checked on [utterance official page](https://utteranc.es/) and see what the generated javascript code is after you change the options.
