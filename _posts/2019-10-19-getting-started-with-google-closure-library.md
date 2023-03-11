---
date: 2019-10-19
title: Getting Started with Google closure library
tags:
  - javascript
  - google closure library
---

Some JavaScript library depends on Google Closure. If you need to understand the behavior of such a library, you have to know closure.
The official document of closure library is [here](https://developers.google.com/closure/library).
And using [Closure compiler](https://developers.google.com/closure/compiler), we can bundle codes with those libraries. There is [npm version](https://github.com/google/closure-compiler-npm) to use such a compiler.

# Examples
1. [./index.html](../../examples/javascript/google-closure-library/index.html)
    - Sample to use google closure library
1. [hello.js](../../examples/javascript/google-closure-library/hello.js)
    - first sample using google cloud compiler
1. [myproject/start.js](../../examples/javascript/google-closure-library/myproject/start.js), [myproject/klass.js](../../examples/javascript/google-closure-library/myproject/klass.js)
    - Sample using module
1. *TODO*
    - Sample to use gooogle closure library and compile

# Google closure compiler

The compiler can be downloaded by npm.
```
> npm install --save-dev google-closure-compiler
```
