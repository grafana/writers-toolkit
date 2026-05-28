import * as fs from "fs";
import type * as Core from "@actions/core";
import type { GitHub } from "@actions/github/lib/utils";
import type { Context } from "@actions/github/lib/context";
import type { RestEndpointMethodTypes } from "@octokit/plugin-rest-endpoint-methods";
import {
  commentMarker,
  formatComment,
  inhibitAlerts,
  parseMarker,
  ValeAlert,
  ValeOutput,
} from "./suggestion";

type ReviewComment =
  RestEndpointMethodTypes["pulls"]["listReviewComments"]["response"]["data"][number];

interface ScriptArgs {
  context: Context;
  core: typeof Core;
  github: InstanceType<typeof GitHub>;
}

function readLine(filePath: string, lineNumber: number): string | undefined {
  return fs.readFileSync(filePath, "utf-8").split("\n")[lineNumber - 1];
}

module.exports = async ({
  context,
  core,
  github,
}: ScriptArgs): Promise<void> => {
  const raw = fs.readFileSync("vale.json", "utf-8");
  if (!raw.trim()) {
    return;
  }

  const valeOutput = inhibitAlerts(JSON.parse(raw) as ValeOutput);

  const { owner, repo } = context.repo;
  const pullNumber = context.issue.number;
  const payload = context.payload as {
    pull_request?: { head?: { sha?: string } };
  };
  const commitSha = payload.pull_request?.head?.sha;

  if (!commitSha) {
    core.setFailed("Could not determine the head commit SHA.");
    return;
  }

  const { data: existingComments } =
    await github.rest.pulls.listReviewComments({
      owner,
      repo,
      pull_number: pullNumber,
    });

  const existingKeys = new Set(
    existingComments.flatMap((c: ReviewComment) => {
      const parsed = parseMarker(c.body);
      return parsed ? [`${c.path}:${parsed.check}:${parsed.match}`] : [];
    }),
  );

  let alertCount = 0;

  for (const [filePath, alerts] of Object.entries(valeOutput)) {
    for (const alert of alerts as ValeAlert[]) {
      alertCount++;
      const line = readLine(filePath, alert.Line) ?? "";
      const body = formatComment(alert, line);
      const dedupeKey = `${filePath}:${alert.Check}:${alert.Match}`;

      if (existingKeys.has(dedupeKey)) {
        continue;
      }

      try {
        await github.rest.pulls.createReviewComment({
          owner,
          repo,
          pull_number: pullNumber,
          commit_id: commitSha,
          path: filePath,
          line: alert.Line,
          body,
        });
      } catch (err) {
        const message = err instanceof Error ? err.message : String(err);
        core.warning(
          `Could not post comment on ${filePath}:${alert.Line}: ${message}`,
        );
      }
    }
  }

  if (process.env.FAIL_ON_ERROR === "true" && alertCount > 0) {
    core.setFailed(`Vale reported ${alertCount} linting error(s).`);
  }
};
