import * as fs from "fs";
import type * as Core from "@actions/core";
import type { GitHub } from "@actions/github/lib/utils";
import type { Context } from "@actions/github/lib/context";
import type { RestEndpointMethodTypes } from "@octokit/plugin-rest-endpoint-methods";

type ReviewComment =
  RestEndpointMethodTypes["pulls"]["listReviewComments"]["response"]["data"][number];

interface ValeAlert {
  Check: string;
  Line: number;
  Link: string;
  Message: string;
  Severity: "error" | "warning" | "suggestion";
}

type ValeOutput = Record<string, ValeAlert[]>;

interface ScriptArgs {
  context: Context;
  core: typeof Core;
  github: InstanceType<typeof GitHub>;
}

const COMMENT_MARKER = "<!-- vale-action -->";

const SEVERITY_EMOJI: Record<ValeAlert["Severity"], string> = {
  error: "🔴",
  warning: "🟡",
  suggestion: "🔵",
};

function formatComment(alert: ValeAlert): string {
  const emoji = SEVERITY_EMOJI[alert.Severity];
  const link = alert.Link ? `\n\n[Style guide](${alert.Link})` : "";

  return `${COMMENT_MARKER}
${emoji} **${alert.Check}** (${alert.Severity})

${alert.Message.trimEnd()}${link}`;
}

module.exports = async ({ context, core, github }: ScriptArgs): Promise<void> => {
  const raw = fs.readFileSync("vale.json", "utf-8");
  if (!raw.trim()) {
    return;
  }

  const valeOutput = JSON.parse(raw) as ValeOutput;

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
    existingComments
      .filter((c: ReviewComment) => c.body.startsWith(COMMENT_MARKER))
      .map((c: ReviewComment) => `${c.path}:${c.line}:${c.body}`),
  );

  for (const [filePath, alerts] of Object.entries(valeOutput)) {
    for (const alert of alerts) {
      const body = formatComment(alert);
      const dedupeKey = `${filePath}:${alert.Line}:${body}`;

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
};
