# Pull Requests Guidelines

See [CONTRIBUTING.md](./CONTRIBUTING.md) for more details about when to create
a GitHub [Pull Request][1] and when other kinds of contributions or
consultation might be more desirable.

When creating a new pull request, please fork the repo and work within a
development branch.

## Commit Messages

* Most changes should be accompanied by tests.
* Commit messages should explain _why_ the changes were made.

``` text
Summary of change in 50 characters or less

Background on why the change is being made with additional detail on
consequences of the changes elsewhere in the code or to the general
functionality of the library. Multiple paragraphs may be used.
```

## Reviews

* Perform a self-review.
* Make sure the CI passes.
* Assign a reviewer once both the above have been completed.

## Merging

* If a maintainer approves the change, it may be merged by the author if
  they have write access. Otherwise, the change will be merged by a maintainer.
* Multiple commits should be squashed before merging.
* Please append the line `closes #<issue-num>: description` in the merge message,
  if applicable.

[1]:  https://help.github.com/articles/about-pull-requests