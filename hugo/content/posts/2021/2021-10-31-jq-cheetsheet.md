---
date: "2021-10-31T00:00:00Z"
last_modified_at: "2023-03-12"
tags:
- cli
- jq
title: jq cheetsheet
---

[`jq`](https://stedolan.github.io/jq/) is used to parse JSON result, format and output on the cli.

This document shows brief overview for how to

# Basics

- Using variables: `exp as $var | exp`
    - `length as $array_length | add / $array_length`
- Boolean operations:
    - AND: `exp and exp`
    - OR: `exp or exp`
- string functions:
    - `contains(substr)` returns boolean
    - `startswith(substr)` returns boolean
    - `join(separator)` returns string
- `select` can be used to filter specific records
- `from_entries` can be used to parse

# Use cases
## Get the object of any key

```jq
exp | to_entries | .[].value | exp
```

## Output CSV/TSV result from JSON

`@csv` and `@tsv` can be used to format the array with CSV or TSV
