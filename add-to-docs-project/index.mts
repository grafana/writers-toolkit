import { GraphQlQueryResponseData } from "@octokit/graphql";
import { Octokit } from "@octokit/rest";
import process from "node:process";
import core from "@actions/core";
import fs from "fs";

// The project ID for the Docs project.
// You can find this by running the `project-id.graphql` query.
const PROJECT_ID = "PVT_kwDOAG3Mbc027w";
const ISSUE_PROJECTS_QUERY = fs.readFileSync("issue-projects.graphql", "utf8");
const ADD_TO_PROJECT_MUTATION = fs.readFileSync(
  "add-to-project.graphql",
  "utf8"
);

async function addIssuesToProject(): Promise<Array<string>> {
  const added: Array<string> = [];
  try {
    const octokit = new Octokit({
      auth: process.env.GITHUB_TOKEN,
    });

    const twoHoursAgo = new Date();
    twoHoursAgo.setHours(twoHoursAgo.getHours() - 2);

    const repositories = await octokit.paginate(octokit.repos.listForOrg, {
      org: "grafana",
      per_page: 100,
    });

    const issuePromises = repositories.map(async (repository) => {
      const issues = await octokit.paginate(octokit.issues.listForRepo, {
        owner: "grafana",
        repo: repository.name,
        filter: "all",
        labels: "type/docs",
        per_page: 100,
        since: twoHoursAgo.toISOString(),
        state: "open",
      });

      const issuePromises = issues.map(async (issue) => {
        if (issue.pull_request) {
          return; // Skip pull requests
        }

        const { node }: GraphQlQueryResponseData = await octokit.graphql(
          ISSUE_PROJECTS_QUERY,
          {
            id: issue.node_id,
          }
        );

        if (
          node.projectItems.nodes.some(
            (item: any) => item.project.id === PROJECT_ID
          )
        ) {
          console.log(
            `Skipping issue ${issue.html_url} because it's already in the project.`
          );

          return;
        }

        console.log(
          `Adding issue ${issue.html_url} to the project if it's not there already.`
        );
        added.push(issue.html_url);

        await octokit.graphql(ADD_TO_PROJECT_MUTATION, {
          projectId: PROJECT_ID,
          contentId: issue.node_id,
        });
      });

      await Promise.all(issuePromises);
    });

    await Promise.all(issuePromises);
  } catch (error: any) {
    console.error("Error adding issues to the project:", error);
    core.setFailed(error);
  }

  return added;
}

const added = await addIssuesToProject();
core.setOutput("added", added.map((url) => `- ${url}`).join("\n"));
