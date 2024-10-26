// Uses https://github.com/actions/github-script
//
// To execute:
// https://github.com/actions/github-script#run-a-separate-file-with-an-async-function
// const script = require('./scripts/docs/vale/from-comment.js');
//
// async script({ core, context, github });
module.exports = async ({ context, core, github }) => {
  const body = context.payload.comment.body || "";

  const { data: files } = await github.rest.pulls.listFiles({
    owner: context.repo.owner,
    repo: context.repo.repo,
    pull_number: context.issue.number,
  });
  const modifiedFiles = files
    .filter((file) => file.additions > 0)
    .map((file) => file.filename);

  core.debug(`Modified files: ${modifiedFiles.join(" ")}`);

  // Expect file paths as unordered Markdown list preceded by any number of whitespace characters.
  const commentFiles = body
    .split("\n")
    .filter((line) => {
      return line.trim().match(/^\s*[-*]\s/);
    })
    .map((line) => {
      return line.trim().replace(/^\s*[-*]\s*/, "");
    });

  core.debug(`Comment files: ${commentFiles.join(" ")}`);

  if (!commentFiles.length) {
    core.setOutput("to-lint", modifiedFiles.join(" "));

    return;
  }

  const toLint = modifiedFiles.filter((file) => {
    return commentFiles.includes(file);
  });

  core.setOutput("to-lint", toLint.join(" "));
};
