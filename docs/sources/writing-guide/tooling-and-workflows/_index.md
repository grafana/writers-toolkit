---
title: Documentation tooling and workflows
menuTitle: Documentation tooling and workflows
description: How to use documentation tools and understand our workflows.
aliases:
  - /docs/writers-toolkit/latest/writing-guide/tooling-and-workflows/
weight: 800
keywords:
  - git
  - repository
  - github
  - review
  - approve
---

# Documentation tooling and workflows

This section provides an overview of the documentation tools we use at Grafana and highlights some of the key processes we use to create documentation.

## Using git

This document explains how to use the `git` command line tool to contribute changes to a Grafana Labs repository.
If you are unfamiliar with `git`, refer to [The Git & Github Bootcamp](https://grafanalabs.udemy.com/course/git-and-github-bootcamp/learn/lecture/24877520%23overview) or [Git and GitHub for Writers](https://grafanalabs.udemy.com/course/git-and-github-for-writers/#overview).

Although processes for contributing changes to a repository differ for each repository, at Grafana there is a generally consistent workflow defined below.

### Creating a local repository

Creating a local repository is only necessary when first contributing to a new repository.
To create a local repository from a remote repository, use `git clone` with the URL of the repository.

There are two types of URL used for cloning: SSH and HTTPS.
Using SSH URLs means that you don't have to provide a username and personal access token when pushing commits.
Instead, authentication with for URLs uses an SSH key.
To set up SSH key authentication in GitHub, refer to [Adding a new SSH key to your GitHub account](https://docs.github.com/en/authentication/connecting-to-github-with-ssh/adding-a-new-ssh-key-to-your-github-account).

The SSH URL for a Grafana repository is `git@github.com:grafana/<REPOSITORY>.git`.
For example, the SSH URL for the Grafana Mimir repository is `git@github.com:grafana/mimir.git`.

Cloning a repository creates a directory containing a git repository and a configured git remote named `origin` that refers to the remote repository.

To clone the Grafana Mimir repository, for example:

```bash
git clone git@github.com:grafana/mimir.git
```

The output is similar to the following:

```console
Cloning into 'mimir'...
remote: Enumerating objects: 111551, done.
remote: Counting objects: 100% (12447/12447), done.
remote: Compressing objects: 100% (3830/3830), done.
remote: Total 111551 (delta 9389), reused 10865 (delta 8511), pack-reused 99104
Receiving objects: 100% (111551/111551), 101.82 MiB | 3.50 MiB/s, done.
Resolving deltas: 100% (73091/73091), done.
```

You can then enter the repository by changing to the newly created directory.

```
cd mimir
```

For a full list of Grafana repositories, refer to [Grafana Labs](https://github.com/orgs/grafana/repositories).

### Downloading updated references from the remote repository

Before contributing changes to a new repository, it's **critical** to have an up-to-date copy of the remote references so that your changes aren't out of date.
To fetch the updated references from the remote repository, use `git fetch`.

```bash
git fetch
```

If there are no updates, the command produces no output.
If there are updates, the output is similar to the following:

```console
comremote: Enumerating objects: 547, done.
remote: Counting objects: 100% (488/488), done.
remote: Compressing objects: 100% (128/128), done.
remote: Total 373 (delta 280), reused 328 (delta 242), pack-reused 0
Receiving objects: 100% (373/373), 156.38 KiB | 1.38 MiB/s, done.
Resolving deltas: 100% (280/280), completed with 79 local objects.
From github.com:grafana/mimir
   659af75c3..321d1ae89  main                             -> origin/main
 * [new branch]          threaded-reader                  -> origin/threaded-reader
```

### Creating a new branch from the main remote branch

By convention, the remote repository in GitHub is the source of truth for a repository's history.
The main branch of a repository is typically called `main` and occasionally called `master`.
Note that we prefer to use inclusive language, so `main` is the preferred name.

After fetching the latest changes to your local repository from the remote repository in GitHub, create a local branch to commit your changes.
Working on your own branch separates and isolates your changes so that they can be later reviewed before incorporation into the main branch.

A branch name should be unique.

To create a new branch called `my-branch` from the remote branch called `main`:

```bash
git checkout -b my-branch origin/main
```

The output is the following:

```console
branch 'my-branch' set up to track 'origin/main'.
Switched to a new branch 'my-branch'
```

You are now on a new local branch and can begin to commit changes.
This means that you are now working on a branch you've created to reflect the changes you're planning to make and can use this branch to develop your content and test different layouts/approaches/structures freely.

To check which branch you currently working on, use `git branch`.
The command outputs a list of local branches with your current branch marked with an asterisk `*`.
For example:

```bash
git branch
```

The output is similar to the following:

```
* my-branch
main
```

Indicating that you are on the branch `my-branch` and there is another local branch named `main`.

Alternatively, you can use `git status` to check your current branch and understand the status of the branch.

> **Note:** `git status` relies on your local repository having the up-to-date references from the remote repository.
> Run `git fetch` before `git status` for the most accurate status.

To understand the output of `git status`, refer to [Git - git-status Documentation](https://git-scm.com/docs/git-status#_output).

### Committing changes to your branch

Committing changes has two steps: staging and committing.
You can stage whole files and directories, or you can stage individual areas within files.

### Staging files and directories

To stage whole files or directories, use `git add <PATH...>`.
To stage the `updated.md` file in your current working directory, for example:

```bash
git add updated.md
```

Stage all files relevant to your current change.

### Staging individual areas within files

Diff hunks are individual areas of difference in files.
Each diff hunk shows one area where a file differs between versions.
You can interactively stage hunks using `git add -p`.
`git` presents you with a diff hunk for each change that you have made and asks if you would like to stage it.
Answering the prompt with `y` stages the hunk.
Answering the prompt with `n` skips staging the hunk.
Answering the prompt with `q` skips staging the hunk and any remaining hunks.
Answering the prompt with `a` stages the hunk and any remaining hunks.

Using the prompt, stage all hunks relevant to your current change.

Once you have staged your changes, you can commit them with `git commit -s`.
`git` opens your text editor where you can type a commit message.

> **Note:** The `-s` flag adds a `Signed-off-by` message to your commits that states you agree to the terms published at https://developercertificate.org/ for that particular commit.
> This is a requirement for a number of repositories.

The first line of a message is the subject.
Commit subjects should be descriptive and concise and are typically written in the imperative, present tense.
To provide additional information, leave a blank line after the subject and write a commit body.
For example:

```
Use American English spellings

American English is preferred by our technical documentation style-guide.
For more information, refer to https://github.com/grafana/technical-documentation/tree/main/docs/sources/style-guide.
```

Finally, save and close the file opened by `git` to finish the commit.

For small changes where you only need write a subject, use the `-m` flag to provide the message without invoking your editor.
For example:

```bash
git commit -s -m "Use American English spellings"
```

### Pushing changes to the remote repository

Pushing changes to the remote repository allows other people to look at your commits and review them.
It's also the first step in getting your changes incorporated into the main branch.

Push your changes using `git push`.

To push your local branch called `my-branch` to a remote branch of the same name:

```bash
git push -u origin my-branch
```

The output is similar to the following:

```console
Enumerating objects: 14, done.
Counting objects: 100% (14/14), done.
Delta compression using up to 8 threads
Compressing objects: 100% (8/8), done.
Writing objects: 100% (10/10), 3.34 KiB | 3.34 MiB/s, done.
Total 10 (delta 5), reused 0 (delta 0), pack-reused 0
remote: Resolving deltas: 100% (5/5), completed with 3 local objects.
remote:
remote: Create a pull request for 'my-branch' on GitHub by visiting:
remote:      https://github.com/grafana/technical-documentation/pull/new/my-branch
remote:
To github.com:grafana/technical-documentation.git
 * [new branch]      HEAD -> my-branch
branch 'my-branch' set up to track 'origin/my-branch'.
```

The response from GitHub includes a link used to open a pull request for your remote branch.
Click the link to open GitHub, then click the green button on the upper left of the screen to open the pull request for reviewers.
Here, you can also edit the title and further detail in the larger text box as well as add reviewers by clicking the Settings icon and entering reviewer GitHub usernames.

### Merging changes from the main branch

Because git and GitHub are naturally collaborative, others are also making and merging pull requests at the same time as you.
From the point where you created your branch and when you're ready to merge the pull request, the main branch might have changed in significant ways, or even developed a conflict with your branch.

If the changes don't conflict, you can use git to incorporate the changes made to the main branch with yours without making any additional changes.

First, fetch changes to all branches and remotes:

```
git fetch --all
```

Update your local copy of the main branch by merging the remote changes into it:

```
git switch main
git merge --ff-only
```

After switching branches, git also lets you know if your local copy needs updating ("Your branch is behind 'origin/main'..."). The `--ff-only` flag means "fast-forwarding", which simply applies the changes to the end of your local copy because your local copy has no conflicting changes.

```
Switched to branch 'main'
Your branch is behind 'origin/main' by 30 commits, and can be fast-forwarded.
  (use "git pull" to update your local branch)
```

Now switch back to your branch and merge the main branch into it, applying its changes into yours:

```
git switch my-branch
git merge main
```

If there are no changes to apply, git confirms this:

```
$ git merge main
Already up to date.
```

Otherwise, if successful, git outputs the changes being applied to your branch:

```
$ git merge main
Removing public/app/plugins/panel/geomap/utils/view.ts
...
Auto-merging .github/CODEOWNERS
Merge made by the 'recursive' strategy.
 .betterer.results                                                           |  142 +--
 .drone.yml                                                                  |  351 ++++++-
 .github/CODEOWNERS                                                          |    3 +
...
```

### Resolving conflicts

GitHub and git both warn you when the source branch and your branch directly conflict.

GitHub notifies you in the pull request if your branch conflicts with the source branch.
If the conflict is easy to resolve, GitHub offers to help you resolve it using the web editor.
See [its documentation](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/addressing-merge-conflicts/about-merge-conflicts) for details.

If the conflict is too complex to resolve in the web editor, GitHub directs you to resolve it on the command line.

If there's a conflict when attempting to merge the main branch into yours, git also tells you which files are in conflict:

```
$ git merge main
Auto-merging .github/CODEOWNERS
CONFLICT (content): Merge conflict in .github/CODEOWNERS
Automatic merge failed; fix conflicts and then commit the result.
```

GitHub also has detailed, cross-platform instructions for resolving a merge conflict using git on the command line.
See [its documentation](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/addressing-merge-conflicts/resolving-a-merge-conflict-using-the-command-line) for details.

