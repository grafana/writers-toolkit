// Copyright 2024 Grafana Labs
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { spawn } from "node:child_process";
import * as readline from "node:readline/promises";
import { Octokit } from "octokit";
import * as fs from "node:fs";
import path from "node:path";

enum Status {
  UNMODIFIED = " ",
  MODIFIED = "M",
  ADDED = "A",
  DELETED = "D",
  RENAMED = "R",
  COPIED = "C",
  UNMERGED = "U",
  UNTRACKED = "?",
  IGNORED = "!",
}
type PathStatus = {
  indexStatus: Status;
  workingTreeStatus: Status;
  path: string;
  originalPath?: string;
};

// Git status porcelain v1 format:
// https://git-scm.com/docs/git-status#_short_format
// XY PATH
// XY ORIG_PATH -> PATH
function parsePorcelainV1(line: string): PathStatus {
  if (line.length < 4) {
    throw new Error(`Invalid porcelain v1 line: ${line}`);
  }

  const indexStatus = line[0] as Status;
  const workingTreeStatus = line[1] as Status;

  const rest = line.slice(3);
  const [originalPath, path] = rest.split(" -> ");
  if (!path) {
    return {
      indexStatus,
      workingTreeStatus,
      path: rest,
    };
  }

  return {
    indexStatus,
    workingTreeStatus,
    path,
    originalPath,
  };
}

async function rsync(source: string, destination: string): Promise<void> {
  return new Promise((resolve, reject) => {
    const rsync = spawn("rsync", [
      "-a",
      "--quiet",
      "--delete",
      source,
      destination,
    ]);

    let stderrOutput = "";
    const stderr = readline.createInterface({ input: rsync.stderr });
    stderr.on("line", (line: string) => {
      stderrOutput += line;
    });

    rsync.on("close", (code: number) => {
      if (code !== 0) {
        reject(`${stderrOutput}\nrsync process exited with code ${code}`);
      }

      resolve();
    });
  });
}

// Return a list of all unstaged Git files in the specified subdirectory.
async function getUnstagedFiles(
  repository: string,
  subdirectory: string
): Promise<Array<PathStatus>> {
  return new Promise((resolve, reject) => {
    const git = spawn("git", ["status", "--porcelain=v1", "--", subdirectory], {
      cwd: repository,
    });

    const files: Array<PathStatus> = [];

    const stdout = readline.createInterface({ input: git.stdout });
    stdout.on("line", (line: string) => {
      files.push(parsePorcelainV1(line));
    });

    let stderrOutput = "";
    const stderr = readline.createInterface({ input: git.stderr });
    stderr.on("line", (line: string) => {
      stderrOutput += line;
    });

    git.on("close", (code: number) => {
      if (code !== 0) {
        reject(`${stderrOutput}\ngit process exited with code ${code}`);
      }

      resolve(files);
    });
  });
}

async function baseTreeSha(
  octokit: Octokit,
  repo: string,
  branch: string
): Promise<string> {
  return (
    await octokit.rest.git.getTree({
      owner: "grafana",
      repo,
      tree_sha: `heads/${branch}`,
    })
  ).data.sha;
}

type sourceContext = {
  name: string;
  branch: string;
  repositoryPath: string;
  subdirectoryPath: string;
};

type websiteContext = {
  repositoryPath: string;
  subdirectoryPath: string;
};

export async function publish(
  octokit: Octokit,
  source: sourceContext,
  website: websiteContext
): Promise<string> {
  rsync(
    path.join(source.repositoryPath, source.subdirectoryPath) + "/",
    path.join(website.repositoryPath, website.subdirectoryPath)
  );
  const files = await getUnstagedFiles(
    website.repositoryPath,
    website.subdirectoryPath
  );
  const baseTree = await baseTreeSha(octokit, "website", "master");

  const newTree = (
    await octokit.rest.git.createTree({
      owner: "grafana",
      repo: "website",
      base_tree: baseTree,
      tree: files.map((file: PathStatus) => {
        if (file.workingTreeStatus === Status.DELETED) {
          return {
            mode: "100644",
            path: file.path,
            sha: null,
            type: "blob",
          };
        } else {
          const fileInfo = fs.statSync(
            path.join(website.repositoryPath, file.path)
          );
          if (fileInfo.isDirectory()) {
            return {
              mode: "040000",
              path: file.path.replace(/\/$/, ""),
              type: "tree",
              sha: null,
            };
          }
          return {
            content: fs
              .readFileSync(path.join(website.repositoryPath, file.path))
              .toString(),
            mode: "100644",
            path: file.path,
            type: "blob",
          };
        }
      }),
    })
  ).data.sha;

  const commit = (
    await octokit.rest.git.createCommit({
      owner: "grafana",
      repo: "website",
      message:
        `Publish from grafana/${source.name}:${source.branch}/${source.subdirectoryPath}\n` +
        "\n" +
        "Co-authored-by: Jack Baldry <jack.baldry@grafana.com>",
      tree: newTree,
      parents: [baseTree],
    })
  ).data.sha;

  octokit.rest.git.deleteRef({
    owner: "grafana",
    repo: "website",
    ref: "heads/jdb/2024-09-test-publish-technical-documentation",
  });

  const websiteCommit = octokit.rest.git.createRef({
    owner: "grafana",
    repo: "website",
    ref: "refs/heads/jdb/2024-09-test-publish-technical-documentation",
    sha: commit,
  });

  // const websiteCommit = (
  //   await octokit.rest.git.updateRef({
  //     owner: "grafana",
  //     repo: "website",
  //     ref: "heads/master",
  //     sha: commit,
  //   })
  // ).data;

  return (await websiteCommit).data.object.sha;
}
