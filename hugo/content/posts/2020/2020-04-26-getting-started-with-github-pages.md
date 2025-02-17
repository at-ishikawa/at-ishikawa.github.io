---
date: "2020-04-26T00:00:00Z"
title: gitHub pages
---

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


Setup GitHub pages site locally
===

See [this official page](https://help.github.com/en/enterprise/2.14/user/articles/setting-up-your-github-pages-site-locally-with-jekyll) for more details for how to see GitHub pages locally.

Getting Started
---
In order to do it, you have to do followings:
1. Install bundler
1. Write `Gemfile` in the root directory of your GitHub repository.
    ```
    source 'https://rubygems.org'

    git_source(:github) {|repo_name| "https://github.com/#{repo_name}" }

    gem "jekyll", "~> 3.8"
    gem 'github-pages', group: :jekyll_plugins
    ```
1. Install the above packages by `bundle install`.
1. Run `bundle exec jekyll serve` and access `http://localhost:4000` on your browser.


Troubleshootings
---
### The warning message `GitHub Metadata: No GitHub API authentication could be found. Some fields may be missing or have incorrect data.` shows up when `bundle exec jekyll serve`.

#### Example
```
> bundle exec jekyll serve
Configuration file: /Users/at-ishikawa/src/github.com/at-ishikawa/at-ishikawa.github.io/_config.yml
            Source: /Users/at-ishikawa/src/github.com/at-ishikawa/at-ishikawa.github.io
       Destination: /Users/at-ishikawa/src/github.com/at-ishikawa/at-ishikawa.github.io/_site
 Incremental build: disabled. Enable with --incremental
      Generating...
Invalid theme folder: _sass
      Remote Theme: Using theme mmistakes/minimal-mistakes
       Jekyll Feed: Generating feed for posts
   GitHub Metadata: No GitHub API authentication could be found. Some fields may be missing or have incorrect data.
                    done in 8.254 seconds.
 Auto-regeneration: enabled for '/Users/at-ishikawa/src/github.com/at-ishikawa/at-ishikawa.github.io'
    Server address: http://127.0.0.1:4000
  Server running... press ctrl-c to stop.
```

#### Solution
See either of the ways in [this document](https://github.com/jekyll/github-metadata/blob/master/docs/authentication.md).
In my case, I created the personal token in GitHub and set the value to the environment variable `JEKYLL_GITHUB_TOKEN`.

#### Details of the issue
This warning message is from [jekyll-github-metadata](https://github.com/jekyll/github-metadata), so you should just follow how to fix this.
However, this is not a big issue as long as you do not run the jekyll with `JEKYLL_ENV=production`.

### No repo name found when `JEKYLL_ENV=production` is set

#### Example
```
> env JEKYLL_ENV=production bundle exec jekyll serve

Configuration file: /Users/at-ishikawa/src/github.com/at-ishikawa/at-ishikawa.github.io/_config.yml
            Source: /Users/at-ishikawa/src/github.com/at-ishikawa/at-ishikawa.github.io
       Destination: /Users/at-ishikawa/src/github.com/at-ishikawa/at-ishikawa.github.io/_site
 Incremental build: disabled. Enable with --incremental
      Generating...
Invalid theme folder: _sass
      Remote Theme: Using theme mmistakes/minimal-mistakes
       Jekyll Feed: Generating feed for posts
   GitHub Metadata: No GitHub API authentication could be found. Some fields may be missing or have incorrect data.
   GitHub Metadata: Error processing value 'url':
  Liquid Exception: No repo name found. Specify using PAGES_REPO_NWO environment variables, 'repository' in your configuration, or set up an 'origin' git remote pointing to your github.com repository. in /_layouts/single.html
             ERROR: YOUR SITE COULD NOT BE BUILT:
                    ------------------------------------
                    No repo name found. Specify using PAGES_REPO_NWO environment variables, 'repository' in your configuration, or set up an 'origin' git remote pointing to your github.com repository.
```

#### Solution
As error message said, set the `repository: owner/repository` in your `_config.yml`, for example.
For other solution, see [this official document](https://github.com/jekyll/github-metadata/blob/master/docs/configuration.md).

#### Details of the issue
If you do not run `jekyll` with `JEKYLL_ENV=production`, the repository configuration is read from `git remote -v` result, but with `JEKYLL_ENV_production`, it's not read.
So, that's why `repository` configuration has to be set.
It's from [github-meatadata](https://github.com/jekyll/github-metadata/blob/v2.13.0/lib/jekyll-github-metadata/repository_finder.rb#L62_L67) as of the version 2.13.


Minimal mistakes
===

Enable search on lunr
---

By following getting started on [official page](https://mmistakes.github.io/minimal-mistakes/docs/configuration/#site-search), search by lunr can be easily enabled.
The default search engine is lunr, and it searches collections of jekyll.
So in order to search pages except posts, those pages should be under a `collection`.

Enable comments
---
You can enable them by following [the official document](https://mmistakes.github.io/minimal-mistakes/docs/configuration/#comments).


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
