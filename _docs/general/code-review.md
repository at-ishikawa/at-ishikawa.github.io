---
title: Code review summaries
---

Last updated: March 2021

This is just a brief note for code reviews.
If you wanna see more comprehensive code review guideline, see [the links on this section](#learn-more-about-code-reviews).


Brief summary of code reviews
====

The main motivations of code reviews are

1. Ensure maintainability of code
1. Check something hard to be found on a development environment
1. Share your knowledge about design or reusable components of code or domain knowledge with a reviewee


Recommended check points
---

These are mostly on backend applications, so frontend perspectives may not be included.

### Production quality

- Performance: N+1 problems are not introduced
- Performance: SQL queries use index properly
- Performance: Throughput of external requests isn't too high
- Reliability: An error is handled correctly. For example,
    - An error is not ignored
	- An error is not caught unnecessarily
	- An error is wrapped from the previous exception and can be traced
	- An error is logged and can be found
- Security: Authorization of a user is correctly verified
- Security: Sensitive information isn't logged out
- Security: Sensitive information isn't passed as a query string

### Code quality

- Readability: Keep or improve simplicity and consistency
- Readability: Table schemas and APIs designs
    - Table schemas and APIs are harder to refactor than application codes, so we should take care more than application design
- Readability: Application design
- Productivity: Avoid introducing new toils unnecessarily
- Testing: Write table driven test cases appropriately when possible
- Testing: Check test cases are sufficient


Non-required check points
---

There are some other important things, but they are/should also be checked differently.

- Ensure no bugs are introduced: It's also checked by QA (and testings)
- Ensure coding rules are complied: It should be checked by CI
- Ensure new codes are covered by unit testings: It should be checked by CI


Review request
---

Developers also need to make code easy to review, and also tell necessary information to reviewers for their changes.
For example, these points can help reviewers to understand your changes:

1. Code itself is easy to review, like code is small enough to review in a short time
1. Reviewers can know the background of changes



Pull request review on GitHub
===

On GitHub, we have a few things to save reviewers time.

1. A PR should be created as a draft if it's not ready for review. If it's not draft, and [code owners](https://github.blog/2017-07-06-introducing-code-owners/) are set to the PR, then they might get notified and start reviewing though it's not ready. Especially, code owners enable [reminders](https://github.blog/2020-04-21-stay-on-top-of-your-code-reviews-with-scheduled-reminders/), they often get notified.



Learn more about code reviews
===
See following links for more sophisicated code review guidelines:

- [Google's Engineering Practices documentation: Code Review Developer Guide](https://google.github.io/eng-practices/review/)
- [Yelp Engineering Blog: Code Review Guidelines](https://engineeringblog.yelp.com/2017/11/code-review-guidelines.html)
