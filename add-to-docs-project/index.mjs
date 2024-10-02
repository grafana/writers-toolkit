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
async function addIssuesToProject(repositories) {
    const added = [];
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
                console.log(`Adding issue ${issue.html_url} to the project if it's not there already.`);
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
    }
    catch (error) {
        console.error("Error adding issues to the project:", error.message);
        core.setFailed(error.message);
    }
    return added;
}
const repositories = ["website", "writers-toolkit"];
const added = await addIssuesToProject(repositories);
core.setOutput("added", added.map((url) => `- ${url}`).join("\n"));
