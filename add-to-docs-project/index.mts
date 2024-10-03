import { GraphQlQueryResponseData } from "@octokit/graphql";
import { Octokit } from "@octokit/rest";
import process from "node:process";
import core from "@actions/core"; // Import the 'core' module

// gh api graphql -f query='
// query{
//     organization(login: "grafana"){
//       projectV2(number: 69) {
//         id
//       }
//     }
//   }'
const PROJECT_ID = "PVT_kwDOAG3Mbc027w";

async function addIssuesToProject(
  repositories: Array<string>
): Promise<Array<string>> {
  const added: Array<string> = [];
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
        if (issue.pull_request) {
          continue; // Skip pull requests
        }

        const { node }: GraphQlQueryResponseData = await octokit.graphql(
          `query issue($id: ID!) {
  node(id: $id) {
    ... on Issue{
      projectItems(first: 100) {
        nodes {
          ... on ProjectV2Item {
            project {
              id
            }
          }
        }
      }
    }
  }
}`,
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

          continue; // Skip issues that are already in the project
        }

        console.log(
          `Adding issue ${issue.html_url} to the project if it's not there already.`
        );
        added.push(issue.html_url);

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
    console.error("Error adding issues to the project:", error);
    core.setFailed(error);
  }

  return added;
}

const repositories = [
  "support-escalations",
  "technical-documentation",
  "website",
  "writers-toolkit",
];

const added = await addIssuesToProject(repositories);
core.setOutput("added", added.map((url) => `- ${url}`).join("\n"));
