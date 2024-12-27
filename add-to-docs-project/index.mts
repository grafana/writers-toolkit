import { GraphQlQueryResponseData } from "@octokit/graphql";
import { Octokit } from "@octokit/rest";
import process from "node:process";
import core from "@actions/core";
import fs from "fs";

// The project ID for the Docs project.
// You can find this by running the `project-id.graphql` query.
const PROJECT_ID = "PVT_kwDOAG3Mbc027w";
const ISSUES_QUERY = fs.readFileSync("issues.graphql", "utf8");
const ADD_TO_PROJECT_MUTATION = fs.readFileSync(
  "add-to-project.graphql",
  "utf8",
);

async function addIssuesToProject(): Promise<Array<string>> {
  const added: Array<string> = [];
  try {
    const octokit = new Octokit({
      auth: process.env.GITHUB_TOKEN,
    });

    const issues = (
      (await octokit.graphql(ISSUES_QUERY)) as GraphQlQueryResponseData
    ).search.nodes;

    for (const issue of issues) {
      console.log(
        `Adding issue ${issue.title} (${issue.url}) to the Docs project.`,
      );
      added.push(
        // https://api.slack.com/reference/surfaces/formatting#escaping
        `${issue.url}|${issue.title}`
          .replaceAll("&", "&amp;")
          .replaceAll("<", "&lt;")
          .replaceAll(">", "&gt;"),
      );

      await octokit.graphql(ADD_TO_PROJECT_MUTATION, {
        projectId: PROJECT_ID,
        contentId: issue.id,
      });
    }
  } catch (error: any) {
    console.error("Error adding issues to the project:", error.message);
    core.setFailed(error.message);
  }

  return added;
}

const added = await addIssuesToProject();
core.setOutput(
  "added",
  added.map((issue) => `- <${issue}>`.replaceAll('"', '\\"')).join("\\n"),
);
