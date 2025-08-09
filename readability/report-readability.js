module.exports = async ({ context, core, github }) => {
  const fs = require("fs");

  const measures = [
    "AutomatedReadability",
    "ColemanLiau",
    "FleschKincaid",
    "FleschReadingEase",
    "GunningFog ",
    "LIX",
    "SMOG",
  ];

  let body = "### Readability report\n";
  body += `| File | ${measures.join(" | ")} |\n`;
  body += `| - |${" - |".repeat(measures.length)}\n`;

  let i = 0;
  const output = fs.readFileSync(".output.txt", "utf-8");
  if (output !== "") {
    output.split("\n").forEach((line) => {
      // docs/sources/_index.md	Grafana.ReadabilityAutomatedReadability	8.99	(Δ+0.84)	aim for below 8.
      const re = new RegExp(
        "^(?<file>[^\t]+?)\t(?<rule>[^\t]+?)\t(?<score>[^\t]+?)\t(?<delta>[^\t]+?)\t(?<guide>[^\t]+?)$",
      );
      const result = re.exec(line);

      if (result !== null) {
        const { file, rule, score, delta, guide } = result.groups;

        if (i === 0) {
          body += `| ${file} |`;
        }

        let color = "green";
        const fl = parseFloat(delta.replace(new RegExp("\\(Δ(.+)\\)"), "$1"));
        if (
          (rule !== "Grafana.ReadabilityFleschReadingEase" && fl > 0) ||
          (rule === "Grafana.ReadabilityFleschReadingEase" && fl < 0)
        ) {
          color = "red";
        }
        if (fl === 0) {
          color = "black";
        }

        body +=
          " $${" + score + "\\space\\color{" + color + "}" + delta + "}$$ |";

        if (i === measures.length - 1) {
          body += "\n";
        }

        i = (i + 1) % measures.length;
      }
    });

    body += `<details>
    <summary>View metric targets</summary>

| Metric | Range | Ideal score |
| --- | --- | --- |
| [AutomatedReadability](https://en.wikipedia.org/wiki/Automated_readability_index) | 6 (very easy read) to 14 (extremely difficult read) | 8 or less |
| [ColemanLiau](https://en.wikipedia.org/wiki/Coleman%E2%80%93Liau_index) | 6 (very easy read) to 17 (extremely difficult read) |  9 or less |
| [FleschKincaid](https://en.wikipedia.org/wiki/Flesch%E2%80%93Kincaid_readability_tests) | 6 (very easy read) to 17 (extremely difficult read) | 8 or less |
| [FleschReadingEase](https://en.wikipedia.org/wiki/Flesch%E2%80%93Kincaid_readability_tests) | 100 (very easy read) to 0 (extremely difficult read) |  70 or more |
| [GunningFog](https://en.wikipedia.org/wiki/Gunning_fog_index) | 6 (very easy read) to 17 (extremely difficult read) | 10 or less |
| [LIX](https://en.wikipedia.org/wiki/Lix_(readability_test)) | 20 (very easy read) to 60+ (extremely difficult read) |  35 or less |
| [SMOG](https://en.wikipedia.org/wiki/SMOG) | 6 (very easy read) to 17 (extremely difficult read) |  10 or less |
</details>`;

    const { data: comments } = await github.rest.issues.listComments({
      issue_number: context.issue.number,
      owner: context.repo.owner,
      repo: context.repo.repo,
    });

    const existing = comments.filter((comment) =>
      comment.body.match(new RegExp("^### Readability report")),
    );
    if (existing.length > 0) {
      github.rest.issues.updateComment({
        comment_id: existing[0].id,
        owner: context.repo.owner,
        repo: context.repo.repo,
        body,
      });
    } else {
      github.rest.issues.createComment({
        issue_number: context.issue.number,
        owner: context.repo.owner,
        repo: context.repo.repo,
        body,
      });
    }
  }
};
