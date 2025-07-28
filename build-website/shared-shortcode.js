// Check for renamed files that contain the shared shortcode and post a comment if found
module.exports = async ({ context, core, github }) => {
  const { data: files } = await github.rest.pulls.listFiles({
    owner: context.repo.owner,
    repo: context.repo.repo,
    pull_number: context.issue.number,
  });

  const renamedFilesWithShared = [];

  for (const file of files) {
    if (file.status === "renamed" && file.additions > 0) {
      try {
        const { data: fileContent } = await github.rest.repos.getContent({
          owner: context.repo.owner,
          repo: context.repo.repo,
          path: file.filename,
          ref: context.payload.pull_request.head.sha,
        });

        const content = Buffer.from(fileContent.content, "base64").toString(
          "utf8",
        );

        if (content.includes("{{< shared")) {
          renamedFilesWithShared.push({
            oldName: file.previous_filename,
            newName: file.filename,
          });
        }
      } catch (error) {
        core.debug(`Error checking file ${file.filename}: ${error.message}`);
      }
    }
  }

  if (renamedFilesWithShared.length > 0) {
    const fileList = renamedFilesWithShared
      .map((file) => `- \`${file.oldName}\` → \`${file.newName}\``)
      .join("\n");

    const commentBody = `⚠️ **Learning Journeys update required**

This PR contains renamed files that use the \`{{< shared ... >}}\` shortcode:

${fileList}

These changes may affect learning journeys content. Review and update any affected learning journeys accordingly.
If you're not sure how to update the learning journeys, ping the @grafana/docs-tooling for assistance.`;

    await github.rest.issues.createComment({
      owner: context.repo.owner,
      repo: context.repo.repo,
      issue_number: context.issue.number,
      body: commentBody,
    });
  }
};
