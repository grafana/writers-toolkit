"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || (function () {
    var ownKeys = function(o) {
        ownKeys = Object.getOwnPropertyNames || function (o) {
            var ar = [];
            for (var k in o) if (Object.prototype.hasOwnProperty.call(o, k)) ar[ar.length] = k;
            return ar;
        };
        return ownKeys(o);
    };
    return function (mod) {
        if (mod && mod.__esModule) return mod;
        var result = {};
        if (mod != null) for (var k = ownKeys(mod), i = 0; i < k.length; i++) if (k[i] !== "default") __createBinding(result, mod, k[i]);
        __setModuleDefault(result, mod);
        return result;
    };
})();
Object.defineProperty(exports, "__esModule", { value: true });
const fs = __importStar(require("fs"));
const COMMENT_MARKER = "<!-- vale-action -->";
const LINK_TEXT = {
    "developers.google.com": "Google developer documentation style guide",
    "grafana.com": "Grafana Writers' Toolkit",
    "html.spec.whatwg.org": "HTML specification",
    "docs.aws.amazon.com": "AWS documentation",
};
function linkText(url) {
    try {
        const { hostname } = new URL(url);
        return LINK_TEXT[hostname] ?? hostname;
    }
    catch {
        return "style guide";
    }
}
function buildSuggestion(alert, filePath) {
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
    let corrected;
    if (Name === "remove") {
        corrected = (before + after).replace(/ {2,}/g, " ").trimEnd();
    }
    else {
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
function formatComment(alert, filePath) {
    const reference = alert.Link
        ? `\n\nFor more information, refer to [${linkText(alert.Link)}](${alert.Link}).`
        : "";
    const suggestion = buildSuggestion(alert, filePath);
    const suggestionBlock = suggestion ? `\n\n${suggestion}` : "";
    return `${COMMENT_MARKER}
**${alert.Check}** (${alert.Severity})

${alert.Message.trimEnd()}${suggestionBlock}${reference}`;
}
module.exports = async ({ context, core, github, }) => {
    const raw = fs.readFileSync("vale.json", "utf-8");
    if (!raw.trim()) {
        return;
    }
    const valeOutput = JSON.parse(raw);
    const { owner, repo } = context.repo;
    const pullNumber = context.issue.number;
    const payload = context.payload;
    const commitSha = payload.pull_request?.head?.sha;
    if (!commitSha) {
        core.setFailed("Could not determine the head commit SHA.");
        return;
    }
    const { data: existingComments } = await github.rest.pulls.listReviewComments({
        owner,
        repo,
        pull_number: pullNumber,
    });
    const existingKeys = new Set(existingComments
        .filter((c) => c.body.startsWith(COMMENT_MARKER))
        .map((c) => `${c.path}:${c.line}:${c.body}`));
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
            }
            catch (err) {
                const message = err instanceof Error ? err.message : String(err);
                core.warning(`Could not post comment on ${filePath}:${alert.Line}: ${message}`);
            }
        }
    }
};
