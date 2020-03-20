---
title: GitHub Pages
---
Written in March 2020.

Getting Started
===
See [Official tutorial](https://help.github.com/en/github/working-with-github-pages/getting-started-with-github-pages) for detail steps.

Basically, what you have to do is
1. Create <username>.github.io repository
1. Choose theme on settings tab in GitHub


Customize pages
===
Before updating your repository, please see each repository [here](https://github.com/pages-themes) for your theme to check how to customize for your purpose.

The examples of customization are
1. Twitter share button: go to [publish twitter button page](https://publish.twitter.com/?buttonText=Share&buttonType=TweetButton&widget=Button).


Choose good theme
===
There are some pages to collect the themes of GitHub pages.
These are the pages I checked.
- [https://github.com/planetjekyll/awesome-jekyll-themes](https://github.com/planetjekyll/awesome-jekyll-themes)
- [#jekyll-theme on GitHub](https://github.com/topics/jekyll-theme)
- [My favorite jekyll themes on GitHub community forum](https://github.community/t5/GitHub-Pages/My-favorite-jekyll-themes-How-bout-you/td-p/24680)

For me, these are the themes I thought they look good.
- https://github.com/Gaohaoyang/gaohaoyang.github.io
    - No instruction on setup
- https://github.com/pmarsceill/just-the-docs
    - Navigations can be configured on front matter in each page.
    - Search data can be generated by `just-the-docs`
- https://github.com/mmistakes/minimal-mistakes
    - Highly customizable
    - Navigations including sidebar have to be configured properly by yaml files.
	- Searching entire pages is enabled by default

The important things for me was sidebar and search features, and I did not have to update every pages for sidebar, so I chose minimal-mistakes theme.


Minimal mistakes
===

Enable search on lunr
---

By following getting started on [official page](https://mmistakes.github.io/minimal-mistakes/docs/configuration/#site-search), search by lunr can be easily enabled.
The default search engine is lunr, and it searches collections of jekyll.
So in order to search pages except posts, those pages should be under a `collection`.


Jekyll
===
[Jekyll](https://jekyllrb.com/) is the site generator developed in ruby, and GitHub pages use it.

Collections
---
[Collections](https://jekyllrb.com/docs/collections/) are the grouped contents, and some common configurations among them can be applied easily, like default front matters.
The directory name of collections must have the `_` prefix like `_docs`, and if it has `output: true` configuration, then it's exported to websites.
In default front matters, the `type` of `scope` can have collection name, for example, in next configuration, `docs` collection has the default layout `single`.

```yml
collections:
  docs:
    output: true

defaults:
  # _docs
  - scope:
      path: ""
      type: docs
    values:
      layout: single
```