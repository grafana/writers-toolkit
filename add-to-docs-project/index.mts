import type { GraphQlQueryResponseData } from "@octokit/graphql";
import { Octokit } from "@octokit/rest";
import process from "node:process";

// gh api graphql -f query='
// query{
//     organization(login: "grafana"){
//       projectV2(number: 69) {
//         id
//       }
//     }
//   }'
const PROJECT_ID = "PVT_kwDOAG3Mbc027w";

async function addIssuesToProject(repositories: Array<string>) {
  try {
    const octokit = new Octokit({
      auth: process.env.GITHUB_TOKEN,
    });

    const twoHoursAgo = new Date();
    twoHoursAgo.setHours(twoHoursAgo.getHours() - 2);

    for (const repo of repositories) {
      const issues = await octokit.paginate(octokit.issues.listForRepo, {
        owner: "grafana",
        repo,
        labels: "type/docs",
        state: "open",
        per_page: 100,
        since: twoHoursAgo.toISOString(),
      });

      for (const issue of issues) {
        console.log(
          `Adding issue https://github.com/grafana/${repo}/issues/${issue.number} to the project if it's not there already.`
        );

        const mutation = `mutation AddProjectItem($projectId: ID!, $contentId: ID!) {
                        addProjectV2ItemById(input: { projectId: $projectId, contentId: $contentId }) {
                                item {
                                        id
                                }
                        }
                }`;

        await octokit.graphql(mutation, {
          projectId: PROJECT_ID,
          contentId: issue.node_id,
        });
      }
    }
  } catch (error: any) {
    console.error("Error adding issues to the project:", error.message);
  }
}

const repositories = ["website"];
addIssuesToProject(repositories);
