import core from "@actions/core";
import github from "@actions/github";
import { Octokit } from "octokit";
import { publish } from "./lib/publish.mjs";
const sourceDirectory = core.getInput("source-directory");
const websiteDirectory = core.getInput("website-directory");
const token = core.getInput("token");
const octokit = new Octokit({ auth: token });
export async function main() {
    publish(octokit, {
        name: github.context.repo.repo,
        branch: github.context.ref,
        repositoryPath: ".",
        subdirectoryPath: sourceDirectory,
    }, {
        repositoryPath: "website",
        subdirectoryPath: websiteDirectory,
    }).catch((error) => {
        console.error(error);
        core.setFailed(error.message);
    });
}
