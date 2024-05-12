---
title: Linux command cheetsheet
date: 2024-03-30
tags:
  - linux
---

# Basic commands

## grep

`grep`'s regular expression is basic regular expression as default.

- `-$NUM`, `-C, --context $NUM`, `-B, --before-context $NUM`, `-A, --after-context $NUM`: show the `$NUM` number of lines around/before/after the selected line
- `-v, --invert-match $REGEXP` option: Select non-matching lines.
- `-E, --extended-regexp $REGEXP` option: Use extended regular expressions (EREs). See [this document](https://www.gnu.org/software/grep/manual/html_node/Basic-vs-Extended.html) for more details of the differences between basic and extended regular expressions.
- `-P, --perl-regexp $REGEXP` option: Use Perl-compatible regular expressions (PCREs).

## sed

- `-n, --quiet, --silent`: Stop printing inputs
- `-i, --in-place $SUFFIX`: Edit files in place
- `-E, -r, --regexp-extended`: Use extended regular expression.

### sed expressions

- Commands
    - `s/$BEFORE/$AFTER/`, `s/$BEFORE/$AFTER/g`: replace a pattern $BEFORE with $AFTER. The separate `/` can be replaced with anything as long as the same separator is used, like `s%/% / %`. Adding `g` replaces all patterns matching instead of only the first pattern.
    - `p`: Print selected address ranges. Use like `${LINE}p`, or `/pattern1/,/pattern2/p`
- Address ranges
    - `$LINE_NUM`
    - `$LINE_NUM_FROM,$LINE_NUM_TO`: Select lines from `$LINE_NUM_FROM` to `$LINE_NUM_TO`

## awk
## xargs
## envsubst

# Use cases
## How to output a line before other lines.

Using awk,
```bash
> echo -e "a\nb\nc" | awk 'BEGIN{ print "name" } { print $1 }'
```

## Filter a line by a regular expresesion

Using grep,

```bash
> echo -e "aa\nab\nac\nba\nca" | grep -E '^a'
aa
ab
ac
```

Using awk,
```bash
> echo -e "aa\nab\nac\nba\nca" | awk '/^a/ { print $1 }'
aa
ab
ac
> echo -e "aa\nab\nac\nba\nca" | awk '/a$/ { print $1 }'
aa
ba
ca
```

## How to select multiple lines by filtering out the first and line lines

```bash
> echo -e "aa\nab\nac\nba\nca"
aa
ab
ac
ba
ca
> echo -e "aa\nab\nac\nba\nca" | sed -n '/ab/,/ba/p'
ab
ac
ba
```

## How to add a character on each line

```bash
> echo -e "aa\nab\nac"
aa
ab
ac
> echo -e "aa\nab\nac" | sed 's/$/,/'
aa,
ab,
ac,
```

## How to remove the last character on a last line

Using `sed '$ s/.$//'`,

```bash
> echo -e "aa\nab\nac"
aa
ab
ac
> echo -e "aa\nab\nac" | sed '$ s/.$//'
aa
ab
a
```
