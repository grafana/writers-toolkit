import core from "@actions/core";
import { Octokit } from "octokit";
import { publish } from "./lib/publish.mjs";
const sourceDirectory = core.getInput("source-directory");
const websiteDirectory = core.getInput("website-directory");
const token = core.getInput("token");
const octokit = new Octokit({ auth: token });
const sourceRepository = ".";
const websiteRepository = "website";
export async function main() {
    publish(octokit, sourceRepository, sourceDirectory, websiteRepository, websiteDirectory).catch((error) => {
        console.error(error);
        core.setFailed(error.message);
    });
}
