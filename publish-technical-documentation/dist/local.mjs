import { env } from "node:process";
import { Octokit } from "octokit";
import { publish } from "./lib/publish.mjs";
try {
    const token = env.GITHUB_TOKEN;
    if (!token) {
        throw new Error("Environment variable GITHUB_TOKEN is required");
    }
    const octokit = new Octokit({ auth: token });
    const sourceName = env.PUBLISH_TECHNICAL_DOCUMENTATION_SOURCE_NAME;
    if (!sourceName) {
        throw new Error("Environment variable PUBLISH_TECHNICAL_DOCUMENTATION_SOURCE_NAME is required");
    }
    const sourceBranch = env.PUBLISH_TECHNICAL_DOCUMENTATION_SOURCE_BRANCH || "main";
    const sourceRepositoryPath = env.PUBLISH_TECHNICAL_DOCUMENTATION_SOURCE_REPOSITORY_PATH;
    if (!sourceRepositoryPath) {
        throw new Error("Environment variable PUBLISH_TECHNICAL_DOCUMENTATION_SOURCE_REPOSITORY_PATH is required");
    }
    const sourceSubdirectoryPath = env.PUBLISH_TECHNICAL_DOCUMENTATION_SOURCE_SUBDIRECTORY_PATH ||
        "docs/sources";
    const websiteRepositoryPath = env.PUBLISH_TECHNICAL_DOCUMENTATION_WEBSITE_REPOSITORY;
    if (!websiteRepositoryPath) {
        throw new Error("Environment variable PUBLISH_TECHNICAL_DOCUMENTATION_WEBSITE_REPOSITORY_PATH is required");
    }
    const websiteSubdirectoryPath = env.PUBLISH_TECHNICAL_DOCUMENTATION_WEBSITE_SUBDIRECTORY_PATH;
    if (!websiteSubdirectoryPath) {
        throw new Error("Environment variable PUBLISH_TECHNICAL_DOCUMENTATION_WEBSITE_SUBDIRECTORY_PATH is required");
    }
    const sha = await publish(octokit, {
        name: sourceName,
        branch: sourceBranch,
        repositoryPath: sourceRepositoryPath,
        subdirectoryPath: sourceSubdirectoryPath,
    }, {
        repositoryPath: websiteRepositoryPath,
        subdirectoryPath: websiteSubdirectoryPath,
    });
    console.log(sha);
}
catch (error) {
    console.error(error);
}
