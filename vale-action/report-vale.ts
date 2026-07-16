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

const HUNK_HEADER =
  /^@@ -(\d+)(?:,(\d+))? \+(\d+)(?:,(\d+))? @@(?: .*)?$/;

export function addedLinesFromPatch(patch: string): Set<number> {
  const added = new Set<number>();
  let newLine = 0;
  let oldRemaining = 0;
  let newRemaining = 0;

  for (const rawLine of patch.split("\n")) {
    if (oldRemaining <= 0 && newRemaining <= 0) {
      const m = HUNK_HEADER.exec(rawLine);
      if (m) {
        oldRemaining = m[2] === undefined ? 1 : parseInt(m[2], 10);
        newLine = parseInt(m[3], 10);
        newRemaining = m[4] === undefined ? 1 : parseInt(m[4], 10);
      }
      continue;
    }

    if (rawLine.startsWith("+")) {
      added.add(newLine);
      newLine++;
      newRemaining--;
    } else if (rawLine.startsWith("-")) {
      oldRemaining--;
    } else if (rawLine.startsWith("\\")) {
      // "\ No newline at end of file" consumes nothing.
    } else {
      newLine++;
      oldRemaining--;
      newRemaining--;
    }
  }

  return added;
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

  const changedFiles = await github.paginate(github.rest.pulls.listFiles, {
    owner,
    repo,
    pull_number: pullNumber,
  });

  const addedLinesByFile = new Map<string, Set<number>>();
  for (const file of changedFiles) {
    if (typeof file.patch === "string") {
      addedLinesByFile.set(file.filename, addedLinesFromPatch(file.patch));
    }
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
    const addedLines = addedLinesByFile.get(filePath);
    for (const alert of alerts as ValeAlert[]) {
      if (!addedLines || !addedLines.has(alert.Line)) {
        continue;
      }

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
