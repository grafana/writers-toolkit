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
import * as fs from "node:fs";
import path from "node:path";
var Status;
(function (Status) {
    Status["UNMODIFIED"] = " ";
    Status["MODIFIED"] = "M";
    Status["ADDED"] = "A";
    Status["DELETED"] = "D";
    Status["RENAMED"] = "R";
    Status["COPIED"] = "C";
    Status["UNMERGED"] = "U";
    Status["UNTRACKED"] = "?";
    Status["IGNORED"] = "!";
})(Status || (Status = {}));
// Git status porcelain v1 format:
// https://git-scm.com/docs/git-status#_short_format
// XY PATH
// XY ORIG_PATH -> PATH
function parsePorcelainV1(line) {
    if (line.length < 4) {
        throw new Error(`Invalid porcelain v1 line: ${line}`);
    }
    const indexStatus = line[0];
    const workingTreeStatus = line[1];
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
async function rsync(source, destination) {
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
        stderr.on("line", (line) => {
            stderrOutput += line;
        });
        rsync.on("close", (code) => {
            if (code !== 0) {
                reject(`${stderrOutput}\nrsync process exited with code ${code}`);
            }
            resolve();
        });
    });
}
// Return a list of all unstaged Git files in the specified subdirectory.
async function getUnstagedFiles(repository, subdirectory) {
    return new Promise((resolve, reject) => {
        const git = spawn("git", ["status", "--porcelain=v1", "--", subdirectory], {
            cwd: repository,
        });
        const files = [];
        const stdout = readline.createInterface({ input: git.stdout });
        stdout.on("line", (line) => {
            files.push(parsePorcelainV1(line));
        });
        let stderrOutput = "";
        const stderr = readline.createInterface({ input: git.stderr });
        stderr.on("line", (line) => {
            stderrOutput += line;
        });
        git.on("close", (code) => {
            if (code !== 0) {
                reject(`${stderrOutput}\ngit process exited with code ${code}`);
            }
            resolve(files);
        });
    });
}
async function baseTreeSha(octokit, repo, branch) {
    return (await octokit.rest.git.getTree({
        owner: "grafana",
        repo,
        tree_sha: `heads/${branch}`,
    })).data.sha;
}
export async function publish(octokit, sourceRepository, sourceDirectory, websiteRepository, websiteDirectory) {
    rsync(path.join(sourceRepository, sourceDirectory) + "/", path.join(websiteRepository, websiteDirectory));
    const files = await getUnstagedFiles(websiteRepository, websiteDirectory);
    const baseTree = await baseTreeSha(octokit, "website", "master");
    const newTree = (await octokit.rest.git.createTree({
        owner: "grafana",
        repo: "website",
        base_tree: baseTree,
        tree: files.map((file) => {
            if (file.workingTreeStatus === Status.DELETED) {
                return {
                    mode: "100644",
                    path: file.path,
                    sha: null,
                    type: "blob",
                };
            }
            else {
                const fileInfo = fs.statSync(path.join(websiteRepository, file.path));
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
                        .readFileSync(path.join(websiteRepository, file.path))
                        .toString(),
                    mode: "100644",
                    path: file.path,
                    type: "blob",
                };
            }
        }),
    })).data.sha;
    const commit = (await octokit.rest.git.createCommit({
        owner: "grafana",
        repo: "website",
        message: `Publish from grafana/writers-toolkit:main/${sourceDirectory}\n` +
            "\n" +
            "Co-authored-by: Jack Baldry <jack.baldry@grafana.com>",
        tree: newTree,
        parents: [baseTree],
    })).data.sha;
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
