---
date: "2020-05-10T00:00:00Z"
tags:
- gcp
- bigquery
title: bigquery usages
---

Functions
===

String
---

### REGEXP_REPLACE
Syntax: `REGEXP_REPLACE(value, regexp, replacement)`.

Returns a string `value` of which sub string matches with `regexp` is replaced with `replacement`.
If `value` contains more than one substrings matching with `regexp`, it's only applied for the first strings.
For the syntax of regexp, we can use the syntax of [re2](https://github.com/google/re2/wiki/Syntax).
For more details, see [official document](https://cloud.google.com/bigquery/docs/reference/standard-sql/string_functions#regexp_replace).

#### Use cases
##### Remove query string from URL
```sql
SELECT REGEXP_REPLACE('https://console.cloud.google.com/bigquery?project=project', '\\?.*$', '');
```


Date
---

### DATE_TRUNC
Syntax: `DATE_TRUNC(date_expression, date_part)`.

Get the preceding date specified in `date_part`.
See [official document](https://cloud.google.com/bigquery/docs/reference/standard-sql/functions-and-operators#date_trunc) for more details.

#### Use cases
##### Get the dates of the beginning and the end of a week
```sql
SELECT
  DATE_TRUNC(CURRENT_DATE, WEEK) AS beginning_of_week,
  DATE_SUB(DATE_TRUNC(DATE_ADD(CURRENT_DATE, INTERVAL 1 WEEK), WEEK), INTERVAL 1 DAY) AS end_of_week,
;
```

The example to get them is:
```shell
> bq query --use_legacy_sql=false 'SELECT
    DATE_TRUNC(DATE \'2020-05-03\', WEEK) AS beginning_of_week_on_sunday,
    DATE_SUB(DATE_TRUNC(DATE_ADD(DATE \'2020-05-03\', INTERVAL 1 WEEK), WEEK), INTERVAL 1 DAY) AS end_of_week_on_sunday,

    DATE_TRUNC(DATE \'2020-05-09\', WEEK) AS beginning_of_week_on_saturday,
    DATE_SUB(DATE_TRUNC(DATE_ADD(DATE \'2020-05-09\', INTERVAL 1 WEEK), WEEK), INTERVAL 1 DAY) AS end_of_week_on_saturday,
  ;
  '
Waiting on bqjob_r36af9e968ad987bb_00000172009190a5_1 ... (0s) Current status: DONE
+-----------------------------+-----------------------+-------------------------------+-------------------------+
| beginning_of_week_on_sunday | end_of_week_on_sunday | beginning_of_week_on_saturday | end_of_week_on_saturday |
+-----------------------------+-----------------------+-------------------------------+-------------------------+
|                  2020-05-03 |            2020-05-09 |                    2020-05-03 |              2020-05-09 |
+-----------------------------+-----------------------+-------------------------------+-------------------------+
```
