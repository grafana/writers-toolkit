---
aliases:
  - /docs/writers-toolkit/writing-guide/tooling-and-workflows/
  - /docs/writers-toolkit/write/tooling-and-workflows/
date: "2022-06-27T12:25:38-05:00"
description: Build and review your content locally.
  Learn how to use documentation tools and understand our workflows.
keywords:
  - git
  - repository
  - github
  - review
  - approve
menuTitle: Tooling and workflows
review_date: "2024-09-03"
title: Documentation tooling and workflows
weight: 200
---

# Documentation tooling and workflows

Grafana manages documentation as code and stores it in version control using Git.
To learn more about how to use Git, refer to [Use Git](#use-git).

## Use Git

Although processes for contributing changes differ for each repository, at Grafana there is a generally consistent workflow.

### Create a local repository

Creating a local repository is only necessary when first contributing to a GitHub repository.
To create a local repository from a remote repository, use either:

- `git clone <URL>`: where _`<URL>`_ is the URL of the repository.

  For more information, refer to [Clone with Git](#clone-with-git).

- `gh repo clone <REPOSITORY>`: where _`<REPOSITORY>`_ is the name of the repository.
  For example, `grafana/writers-toolkit`

  For more information, refer to [Clone with GitHub CLI](#clone-with-github-cli).

#### Clone with GitHub CLI

You must first have installed GitHub CLI and authenticated with GitHub.
To install the GitHub CLI tool, refer to [Installation](https://github.com/cli/cli#installation).
To authenticate with GitHub, refer to [Authenticate with GitHub](#authenticate-with-github).

To clone the `grafana/writers-toolkit` repository, change to the directory containing your repositories and run:

```bash
gh repo clone grafana/writers-toolkit
```

The output is similar to the following:

```console
Cloning into 'writers-toolkit'...
remote: Enumerating objects: 5599, done.
remote: Counting objects: 100% (1493/1493), done.
remote: Compressing objects: 100% (682/682), done.
remote: Total 5599 (delta 1065), reused 1086 (delta 750), pack-reused 4106
Receiving objects: 100% (5599/5599), 8.11 MiB | 4.68 MiB/s, done.
Resolving deltas: 100% (3314/3314), done.
```

You can then enter the repository by changing to the newly created directory.

```bash
cd mimir
```

For a full list of Grafana repositories, refer to [the GitHub repository list for the Grafana organization](https://github.com/orgs/grafana/repositories).

#### Clone with Git

<!-- vale Grafana.Timeless = NO -->
<!-- Usage comes from the external page title -->

There are two types of URL used for cloning: SSH and HTTPS.
Using SSH URLs means that you don't have to provide a username and personal access token when pushing commits.
Instead, authentication with for URLs uses an SSH key.
To set up SSH key authentication in GitHub, refer to [Adding a new SSH key to your GitHub account](https://docs.github.com/en/authentication/connecting-to-github-with-ssh/adding-a-new-ssh-key-to-your-github-account).

<!-- vale Grafana.Timeless = YES -->

The SSH URL for a Grafana repository is `git@github.com:grafana/<REPOSITORY>.git`.
For example, the SSH URL for the Grafana Mimir repository is `git@github.com:grafana/mimir.git`.

Cloning a repository creates a directory containing a Git repository and a configured Git remote named `origin` that refers to the remote repository.

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

```bash
cd mimir
```

For a full list of Grafana repositories, refer to [the GitHub repository list for the Grafana organization](https://github.com/orgs/grafana/repositories).

### Download updated references from the remote repository

Before contributing changes to a repository, it's important to have an up-to-date copy of the remote references so that your changes aren't out of date.
To fetch the updated references from the remote repository, use `git fetch`.

```bash
git fetch
```

If there are no updates, the command produces no output.
If there are updates, the output is similar to the following:

```console
remote: Enumerating objects: 547, done.
remote: Counting objects: 100% (488/488), done.
remote: Compressing objects: 100% (128/128), done.
remote: Total 373 (delta 280), reused 328 (delta 242), pack-reused 0
Receiving objects: 100% (373/373), 156.38 KiB | 1.38 MiB/s, done.
Resolving deltas: 100% (280/280), completed with 79 local objects.
From github.com:grafana/mimir
   659af75c3..321d1ae89  main                             -> origin/main
 * [new branch]          threaded-reader                  -> origin/threaded-reader
```

### Create a branch from the default remote branch

By convention, the remote repository in GitHub is the source of truth for a repository's history.
The default branch of a repository is typically called `main` and occasionally called `master`.
Grafana prefers to use inclusive language, so `main` is the preferred name.

After fetching the changes from the remote repository in GitHub, create a local branch to commit your changes.
Working on your own branch separates and isolates your changes so that they can be later reviewed before incorporation into the default branch.

A branch name should be unique.

To create a branch called `my-branch` from the remote branch called `main`:

```bash
git checkout -b my-branch origin/main
```

The output is the following:

```console
branch 'my-branch' set up to track 'origin/main'.
Switched to a new branch 'my-branch'
```

<!-- vale Grafana.Timeless = NO -->
<!-- This isn't discussing a new feature but instead a new result of a command -->

You are now on a new local branch and can begin to commit changes.
This means that you are now working on a branch you've created to reflect the changes you're planning to make and can use this branch to develop your content and test different layouts, approaches, or structures freely.

<!-- vale Grafana.Timeless = YES -->

To check which branch you are working on, use `git branch`.
The command outputs a list of local branches with your current branch marked with an asterisk `*`.
For example:

```bash
git branch
```

The output is similar to the following:

```console
* my-branch
main
```

The output indicates that you are on the branch `my-branch` and there is another local branch named `main`.

Alternatively, you can use `git status` to check your current branch and understand the status of the branch.

{{< admonition type="note" >}}
`git status` relies on your local repository having the up-to-date references from the remote repository.

Run `git fetch` before `git status` for the most accurate status.
{{< /admonition >}}

To understand the output of `git status`, refer to [Git - `git-status` Documentation](https://git-scm.com/docs/git-status#_output).

### Check out a PR branch

To check out a PR branch, use the GitHub CLI (`gh`) tool.
To install the GitHub CLI tool, refer to [Installation](https://github.com/cli/cli#installation).

It fetches from the remote repository and configures a local branch in your repository to track that remote branch.

Run the command from a directory within your local checkout of the repository to check out a contributor's PR.
Replace _`<PR NUMBER>`_ with the number of the pull request.

```shell
gh pr checkout <PR NUMBER>
```

### Check out a PR branch from a fork

A fork is a repository that shares code and visibility settings with the original _upstream_ repository.

For more information, refer to [About forks](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/working-with-forks/about-forks).

To check out a PR branch from a fork, use the GitHub CLI (`gh`) tool.
To install the GitHub CLI tool, refer to [Installation](https://github.com/cli/cli#installation).

It fetches from the fork remote repository and configures a local branch in your repository to track that remote branch in the fork.

Run the command from a directory within your local checkout of the upstream repository to check out a contributor's PR.
Replace _`<PR NUMBER>`_ with the number of the pull request.

```shell
gh pr checkout <PR NUMBER>
```

The output is similar to the following:

```console
remote: Enumerating objects: 14, done.
remote: Counting objects: 100% (9/9), done.
remote: Compressing objects: 100% (2/2), done.
remote: Total 14 (delta 7), reused 7 (delta 7), pack-reused 5
Unpacking objects: 100% (14/14), 7.68 KiB | 1.10 MiB/s, done.
From github.com:grafana/grafana
 * [new ref]               refs/pull/76925/head -> patch-2
Switched to branch 'patch-2'
```

The GitHub CLI tool automatically configures the remote tracking branch so you can push and pull from the fork branch.

<!-- vale Grafana.Timeless = NO -->
<!-- Usage comes from the external page title -->

{{< admonition type="note" >}}
You can only push to a fork if the PR author has enabled **Allow edits and access to secrets by maintainers** or **Allow edits by maintainers**.

To enable maintainer edits, refer to [Enabling repository maintainer permissions on existing pull requests](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/working-with-forks/allowing-changes-to-a-pull-request-branch-created-from-a-fork#enabling-repository-maintainer-permissions-on-existing-pull-requests).
{{< /admonition >}}

<!-- vale Grafana.Timeless = YES -->

### Commit changes to your branch

Committing changes has two steps: staging and committing.
You can stage whole files and directories, or you can stage individual areas within files.

### Stage files and directories

To stage whole files or directories, use `git add <PATH...>`.
To stage the `updated.md` file in your current working directory, for example:

```bash
git add updated.md
```

Stage all files relevant to your current change.

### Stage individual areas within files

Diff hunks are individual areas of difference in files.
Each diff hunk shows one area where a file differs between versions.
You can interactively stage hunks using `git add -p`.
`git` presents you with a diff hunk for each change that you have made and asks if you would like to stage it.
Answering the prompt with `y` stages the hunk.
Answering the prompt with `n` skips staging the hunk.
Answering the prompt with `q` skips staging the hunk and any remaining hunks.
Answering the prompt with `a` stages the hunk and any remaining hunks.

Using the prompt, stage all hunks relevant to your current change.

After you have staged your changes, you can commit them with `git commit -s`.
`git` opens your text editor where you can type a commit message.

{{< admonition type="note" >}}
The `-s` flag adds a `Signed-off-by` message to your commits that states you agree to the terms published at https://developercertificate.org/ for that particular commit.

This is a requirement for a number of repositories.
{{< /admonition >}}

The first line of a message is the subject.
Commit subjects should be descriptive and concise and are typically written in the imperative, present tense.
To provide additional information, leave a blank line after the subject and write a commit body.
For example:

```
Use US English spellings

US English is preferred by our technical documentation style-guide.
For more information, refer to https://github.com/grafana/technical-documentation/tree/main/docs/sources/style-guide.
```

Save and close the file opened by Git to finish the commit.

For small changes where you only need write a subject, use the `-m` option to provide the message without invoking your editor.
For example:

```bash
git commit -s -m "Use US English spellings"
```

You can provide the`-m` option multiple times.
Git uses the argument to the first option as the commit subject.
The arguments to the other options become the commit body paragraphs.

### Push changes to the remote repository

Pushing changes to the remote repository allows other people to look at your commits and review them.
It's also the first step in getting your changes into the default branch.

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
Here, you can also edit the title and further detail in the larger text box as well as add reviewers by clicking **Reviewers** and entering reviewer GitHub usernames.

### Force push changes

If you rewrite local history, or your local branch diverges from the one in the remote for other reasons, you might need to force the remote to accept your changes.

In such cases, prefer `--force-with-lease` over `--force`, which overwrites the branch in the remote only if your local branch is up to date.
That way, you won't accidentally overwrite commits pushed by others that you didn't know about.
For more information, refer to the documentation of the option in `man git-push`.

### Merge changes from the main branch

Git and GitHub allow many authors to make and merge pull requests at the same time as you.
From the point where you created your branch and when you're ready to merge the pull request, the default branch might have changed in significant ways, or even developed a conflict with your branch.

If the changes don't conflict, you can use Git to incorporate the changes made to the default branch with yours without making any additional changes.

First, fetch changes to all branches and remotes:

```bash
git fetch --all
```

Update your local copy of the main branch by merging the remote changes into it:

```bash
git switch main
git merge --ff-only
```

After switching branches, Git also lets you know if your local copy needs updating with the message _Your branch is behind 'origin/main'_.
The `--ff-only` flag means _fast-forwarding_, which applies the changes to the end of your local copy when your local copy has no conflicting changes.

```console
Switched to branch 'main'
Your branch is behind 'origin/main' by 30 commits, and can be fast-forwarded.
  (use "git pull" to update your local branch)
```

Now switch back to your branch and merge the default branch into it, applying its changes into yours:

```bash
git switch my-branch
git merge main
```

If there are no changes to apply, Git confirms this.
The output is similar to the following:

```console
Already up to date.
```

Otherwise, if successful, Git outputs the changes it applied to your branch:
The output is similar to the following:

```console
Removing public/app/plugins/panel/geomap/utils/view.ts
...
Auto-merging .github/CODEOWNERS
Merge made by the 'recursive' strategy.
 .betterer.results                                                           |  142 +--
 .drone.yml                                                                  |  351 ++++++-
 .github/CODEOWNERS                                                          |    3 +
...
```

### Resolve conflicts

GitHub and Git both warn you when the source branch and your branch directly conflict.

GitHub notifies you in the pull request if your branch conflicts with the target branch.
If the conflict is easy to resolve, GitHub offers to help you resolve it using the web editor.
For more information refer to [About merge conflicts](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/addressing-merge-conflicts/about-merge-conflicts).

If the conflict is too complex to resolve in the web editor, GitHub directs you to resolve it on the command line.

If there's a conflict when attempting to merge the target branch into yours, Git also tells you which files are in conflict:

```console
Auto-merging .github/CODEOWNERS
CONFLICT (content): Merge conflict in .github/CODEOWNERS
Automatic merge failed; fix conflicts and then commit the result.
```

GitHub has detailed, cross-platform instructions for resolving a merge conflict using Git on the command line.
For more information refer to [Resolving a merge conflict using the command line](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/addressing-merge-conflicts/resolving-a-merge-conflict-using-the-command-line).

## Sync a fork with its upstream

To sync a fork with its upstream repository, change to the directory of your repository checkout and run:

```bash
gh repo sync <FORK REPOSITORY>
```

For more information about the command's options, refer to the [`gh repo sync` documentation](https://cli.github.com/manual/gh_repo_sync).

## Use GitHub CLI

GitHub CLI simplifies some Git workflows when working with GitHub repositories.

### Install GitHub CLI

To install the GitHub CLI tool, refer to [Installation](https://github.com/cli/cli#installation).

### Authenticate with GitHub

After installing GitHub CLI you must authenticate with GitHub.
To authenticate with GitHub, run `gh auth login` and follow the interactive setup.

You can also use `gh` as a credential helper for Git.
To use `gh` as a credential helper for Git, run `gh auth setup-git`.
For more information about Git credential helpers, refer to [Git - `gitcredentials` Documentation](https://git-scm.com/docs/gitcredentials).
