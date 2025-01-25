---
date: "2020-09-03T00:00:00Z"
tags:
- php
title: PHP Benchmarks
---

PHPBench framework
===

There is an [phpbench](https://phpbench.readthedocs.io/en/latest/index.html) for php benchmarks.

Set up
---
1. Install: `composer require phpbench/phpbench --dev`
1. Add `phpbench.json` under a project root directory
```
{
    "bootstrap": "vendor/autoload.php"
}
```

Example of benchmark code
---
The example to run a benchmark code.
There are a few annotations for methods of phpbench. For example,
    - `Revs`: How many times code is executed within a single measurement.
    - `Iterations`: How many times measure execute the benchmark

The example file is [`benchmarks/TimerBench.php`](/examples/php/benchmark/benchmarks/TimerBench.php).

In order to run the benchmark with a default report format,
```
> ./vendor/bin/phpbench run benchmarks/TimerBench.php --report=default -vvv
PhpBench @git_tag@. Running benchmarks.
Using configuration file: /Users/at-ishikawa/src/github.com/at-ishikawa/at-ishikawa.github.io/examples/php/benchmark/phpbench.json
\TimerBench
    benchConsume............................I4 [μ Mo]/r: 130.338 130.652 (μs) [μSD μRSD]/r: 0.505μs 0.39%
1 subjects, 5 iterations, 1,000 revs, 0 rejects, 0 failures, 0 warnings
(best [mean mode] worst) = 129.514 [130.338 130.652] 130.824 (μs)
⅀T: 651.689μs μSD/r 0.505μs μRSD/r: 0.388%
suite: 1343dc77b361e3f5b474c1cf1e36eabef7a92edb, date: 2020-09-03, stime: 05:28:16
+------------+--------------+-----+------+------+----------+-----------+--------------+----------------+
| benchmark  | subject      | set | revs | iter | mem_peak | time_rev  | comp_z_value | comp_deviation |
+------------+--------------+-----+------+------+----------+-----------+--------------+----------------+
| TimerBench | benchConsume | 0   | 1000 | 0    | 987,712b | 130.753μs | +0.82σ       | +0.32%         |
| TimerBench | benchConsume | 0   | 1000 | 1    | 987,712b | 130.824μs | +0.96σ       | +0.37%         |
| TimerBench | benchConsume | 0   | 1000 | 2    | 987,712b | 130.604μs | +0.53σ       | +0.20%         |
| TimerBench | benchConsume | 0   | 1000 | 3    | 987,712b | 129.514μs | -1.63σ       | -0.63%         |
| TimerBench | benchConsume | 0   | 1000 | 4    | 987,712b | 129.994μs | -0.68σ       | -0.26%         |
+------------+--------------+-----+------+------+----------+-----------+--------------+----------------+
```

With an aggregate report format,
```
> ./vendor/bin/phpbench run benchmarks/TimerBench.php --report=aggregate -vvv
PhpBench @git_tag@. Running benchmarks.
Using configuration file: /Users/at-ishikawa/src/github.com/at-ishikawa/at-ishikawa.github.io/examples/php/benchmark/phpbench.json
\TimerBench
    benchConsume............................I4 [μ Mo]/r: 129.945 129.961 (μs) [μSD μRSD]/r: 0.209μs 0.16%
1 subjects, 5 iterations, 1,000 revs, 0 rejects, 0 failures, 0 warnings
(best [mean mode] worst) = 129.621 [129.945 129.961] 130.196 (μs)
⅀T: 649.725μs μSD/r 0.209μs μRSD/r: 0.161%
suite: 1343dc74f4317ec67c2c151bea7d01e97e447eed, date: 2020-09-03, stime: 05:28:20
+------------+--------------+-----+------+-----+----------+-----------+-----------+-----------+-----------+---------+--------+-------+
| benchmark  | subject      | set | revs | its | mem_peak | best      | mean      | mode      | worst     | stdev   | rstdev | diff  |
+------------+--------------+-----+------+-----+----------+-----------+-----------+-----------+-----------+---------+--------+-------+
| TimerBench | benchConsume | 0   | 1000 | 5   | 987,712b | 129.621μs | 129.945μs | 129.961μs | 130.196μs | 0.209μs | 0.16%  | 1.00x |
+------------+--------------+-----+------+-----+----------+-----------+-----------+-----------+-----------+---------+--------+-------+
```
