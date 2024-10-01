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
async function addIssuesToProject(repositories) {
    try {
        const octokit = new Octokit({
            auth: process.env.GITHUB_TOKEN,
        });
        const query = `{
  organization(login: "grafana") {
    projectV2(number: 69) {
      title
      items(first:100) {
        totalCount
        nodes {
          ... on ProjectV2Item {
            fieldValueByName(name: "Status") {
              ... on ProjectV2ItemFieldSingleSelectValue {
                name
                id
              }
            }
          }
        }
      }
    }
  }
}`;
        const response = await octokit.graphql(query, {
            projectId: PROJECT_ID,
        });
        console.log(response);
        const projectIssues = response.items.edges.map((edge) => edge.node.number);
        const oneWeekAgo = new Date();
        oneWeekAgo.setDate(oneWeekAgo.getDate() - 7);
        for (const repo of repositories) {
            const issues = await octokit.paginate(octokit.issues.listForRepo, {
                owner: "grafana",
                repo,
                labels: "type/docs",
                state: "open",
                per_page: 100,
                since: oneWeekAgo.toISOString(),
            });
            for (const issue of issues) {
                if (projectIssues.includes(issue.number)) {
                    console.log(`Issue ${issue.number} is already in the project`);
                    continue;
                }
                console.log(`Adding issue https://github.com/grafana/${repo}/issues/${issue.number} to the project`);
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
    }
    catch (error) {
        console.error("Error adding issues to the project:", error.message);
    }
}
const repositories = ["website"];
addIssuesToProject(repositories);
