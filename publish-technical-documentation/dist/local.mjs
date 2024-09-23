import { env } from "node:process";
import { Octokit } from "octokit";
import { publish } from "./lib/publish.mjs";
try {
    const token = env.GITHUB_TOKEN;
    if (!token) {
        throw new Error("Environment variable GITHUB_TOKEN is required");
    }
    const octokit = new Octokit({ auth: token });
    const sourceRepository = env.PUBLISH_TECHNICAL_DOCUMENTATION_SOURCE_REPOSITORY;
    if (!sourceRepository) {
        throw new Error("Environment variable PUBLISH_TECHNICAL_DOCUMENTATION_SOURCE_REPOSITORY is required");
    }
    const sourceDirectory = env.PUBLISH_TECHNICAL_DOCUMENTATION_SOURCE_DIRECTORY || "docs/sources";
    const websiteRepository = env.PUBLISH_TECHNICAL_DOCUMENTATION_WEBSITE_REPOSITORY;
    if (!websiteRepository) {
        throw new Error("Environment variable PUBLISH_TECHNICAL_DOCUMENTATION_WEBSITE_REPOSITORY is required");
    }
    const websiteDirectory = env.PUBLISH_TECHNICAL_DOCUMENTATION_WEBSITE_DIRECTORY;
    if (!websiteDirectory) {
        throw new Error("Environment variable PUBLISH_TECHNICAL_DOCUMENTATION_WEBSITE_DIRECTORY is required");
    }
    const sha = await publish(octokit, sourceRepository, sourceDirectory, websiteRepository, websiteDirectory);
    console.log(sha);
}
catch (error) {
    console.error(error);
}
