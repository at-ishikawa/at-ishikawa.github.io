---
date: "2020-03-19T21:00:00Z"
tags:
- git
- jekyll
- github-pages
- minimal-mistakes
title: Enable lunr search for non-posts on minimal mistake theme of GitHub pages
---

This page explains how to enable searching non-posts pages for [Minimal mistakes](https://github.com/mmistakes/minimal-mistakes) by [Lunr.js](https://lunrjs.com/) for someone who does not know jekyll at all.
Lunr.js is the default search engine of the theme, so if turn on search by adding `search: true` in `_config.yml`, then you all posts become searchable by lunr.js.
However, for non-post pages, you may need a different configuration, so I'll write this page for it.
The entire configuration which requires is following in this page.

```yml
search: true

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


# Add new collections for your pages
First of all, you have to understand [collections](https://jekyllrb.com/docs/collections/) in jekyll.
Collections is the grouped contents that are under a specific directory.
For example, if you have next configurations in `_config.yml`, then the directory `_docs` becomes collections.
`_` prefix of the directory is required for a collection.

```yml
collections:
  docs:
    output: true
```

If we add `output: true`, then it's gonna exported to `_site` directory when web sites is generated so that we can see the generated contents.
However, if we only create this, you notice the defaults front matter is not configured properly for those pages.


# Update default frontmatter
The default frontmatter has the `type` in `scope`, and this `type` is actually the collection name.
So if you set the type to `docs`, then the frontmatter is applied to the collection.


# Why did I write this post?
Because I had this issue.
And I did not know the details about jekyll and spent a lot of time to figure out how to configure frontmatter to a collection.
I saw this [issue](https://github.com/jekyll/jekyll/issues/2405) and figured out how to do it.
