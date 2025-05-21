---
date: "2024-03-30T00:00:00Z"
last_modified_at: "2025-05-20"
tags:
- linux
- awk
- sed
- grep
title: Linux command cheetsheet
---

# String handling

For more details, see a [shell parameter expansion](https://www.gnu.org/software/bash/manual/html_node/Shell-Parameter-Expansion.html) page.

- `$[var##pattern}`: Remove the beginning of a string matching a pattern

    > var="abc%cdg.txt"
    > echo ${var##*.}
    txt
    > echo ${var##*%}
    cdg.txt

- `${var%%pattern}`: Remove the trailing portion of a string matching a pattern

    > var="abc%cdg.txt"
    > echo ${var%%.*}
    abc%cdg
    > echo ${var%%\%*}
    abc

# Basic commands
## file names

- `basename`: Extract the basename of a file from a file path
- `dirname`: Extract the parent directory name from a file path

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
    - `i`: Insert lines BEFORE an address. Use like

        sed '/pattern/i \
        text before pattern'

    - `a`: Append lines AFTER an address. Use like

        sed '/pattern/a \
        text after pattern'

- Address ranges
    - `$LINE_NUM`
    - `$LINE_NUM_FROM,$LINE_NUM_TO`: Select lines from `$LINE_NUM_FROM` to `$LINE_NUM_TO`

## awk

- `-F` is field separator. The default value is a `0x20`, which matches a space, tab, and newlines [this article](https://stackoverflow.com/a/30406868).
- There can be some generic syntaxes like
    - Variable can be defined and used
        - Internal variables: `{ var=$1; print var }`
    - Conditional flow like `if (condition1) { statement1; }`
        - `next` can skip a line

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

## How to add a character at the end of each line

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

## How to output the 1st and 3rd column from a CSV file

Using `awk`

```bash
> echo -e 'aa,ab,ac\nba,bb,bc\nca,cb,cc' | awk -F ',' '{ print $1,":",$3 }'
aa : ac
ba : bc
ca : cc
```

Using `cut`

```bash
> echo -e 'aa,ab,ac\nba,bb,bc\nca,cb,cc' | cut -d "," -f 1,3
aa,ac
ba,bc
ca,cc
```

## How to add multiple lines before or after specific patterns

Use `sed` with an `i` or `a` expression.

```bash
> echo -e 'line 1\nline 2\nline 3' | sed '/line 2/i \
new line'
line 1
new line
line 2
line 3
> echo -e 'line 1\nline 2\nline 3' | sed '/line 2/a \
new line'
line 1
line 2
new line
line 3
```

## How to extract file names from a file path

```bash
> filepath=/var/log/mail.log.2.gz

> echo $(dirname $filepath)
/var/log
> echo $(basename $filepath)
mail.log.2.gz

## Extract extension
> echo ${filepath##*.}
gz
> echo ${filepath#*.}
log.2.gz

# Extract a filename without extensions
> echo $(basename ${filepath%%.*})
mail
> echo $(basename ${filepath%.*})
mail.log.2
```
