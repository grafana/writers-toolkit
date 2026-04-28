import * as fs from "fs";
import type * as Core from "@actions/core";
import type { GitHub } from "@actions/github/lib/utils";
import type { Context } from "@actions/github/lib/context";
import type { RestEndpointMethodTypes } from "@octokit/plugin-rest-endpoint-methods";

type ReviewComment =
  RestEndpointMethodTypes["pulls"]["listReviewComments"]["response"]["data"][number];

interface ValeAction {
  Name: "replace" | "remove" | "";
  Params: string[] | null;
}

interface ValeAlert {
  Action: ValeAction;
  Check: string;
  Line: number;
  Link: string;
  Match: string;
  Message: string;
  Severity: "error" | "warning" | "suggestion";
  Span: [number, number];
}

type ValeOutput = Record<string, ValeAlert[]>;

interface ScriptArgs {
  context: Context;
  core: typeof Core;
  github: InstanceType<typeof GitHub>;
}

const COMMENT_MARKER = "<!-- vale-action -->";

const LINK_TEXT: Record<string, string> = {
  "developers.google.com": "Google developer documentation style guide",
  "grafana.com": "Grafana Writers' Toolkit",
  "html.spec.whatwg.org": "HTML specification",
  "docs.aws.amazon.com": "AWS documentation",
};

function linkText(url: string): string {
  try {
    const { hostname } = new URL(url);
    return LINK_TEXT[hostname] ?? hostname;
  } catch {
    return "style guide";
  }
}

function buildSuggestion(alert: ValeAlert, filePath: string): string | null {
  const { Name, Params } = alert.Action;

  if (Name !== "replace" && Name !== "remove") {
    return null;
  }

  const lines = fs.readFileSync(filePath, "utf-8").split("\n");
  const line = lines[alert.Line - 1];
  if (line === undefined) {
    return null;
  }

  const [start, end] = alert.Span;
  const before = line.slice(0, start - 1);
  const after = line.slice(end);

  let corrected: string;
  if (Name === "remove") {
    corrected = (before + after).replace(/ {2,}/g, " ").trimEnd();
  } else {
    const first = Params?.[0] ?? "";
    corrected = before + first + after;
  }

  let suggestion = "```suggestion\n" + corrected + "\n```";

  if (Name === "replace" && Params && Params.length > 1) {
    const alternatives = Params.slice(1)
      .map((p) => `\`${p}\``)
      .join(", ");
    suggestion += `\n\nAlternatives: ${alternatives}`;
  }

  return suggestion;
}

function formatComment(alert: ValeAlert, filePath: string): string {
  const reference = alert.Link
    ? `\n\nFor more information, refer to [${linkText(alert.Link)}](${alert.Link}).`
    : "";

  const suggestion = buildSuggestion(alert, filePath);
  const suggestionBlock = suggestion ? `\n\n${suggestion}` : "";

  return `${COMMENT_MARKER}
**${alert.Check}** (${alert.Severity})

${alert.Message.trimEnd()}${suggestionBlock}${reference}`;
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
      const body = formatComment(alert, filePath);
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
