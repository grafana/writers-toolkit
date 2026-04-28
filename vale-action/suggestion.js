"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.commentMarker = commentMarker;
exports.parseMarker = parseMarker;
exports.linkText = linkText;
exports.applySuggestion = applySuggestion;
exports.formatComment = formatComment;
const LINK_TEXT = {
    "developers.google.com": "Google developer documentation style guide",
    "grafana.com": "Grafana Writers' Toolkit",
    "html.spec.whatwg.org": "HTML specification",
    "docs.aws.amazon.com": "AWS documentation",
};
function commentMarker(check, match) {
    return `<!-- vale-action check="${check}" match="${match}" -->`;
}
function parseMarker(body) {
    const m = body.match(/<!-- vale-action check="([^"]*)" match="([^"]*)" -->/);
    if (!m) {
        return undefined;
    }
    return { check: m[1], match: m[2] };
}
function linkText(url) {
    try {
        const { hostname } = new URL(url);
        return LINK_TEXT[hostname] ?? hostname;
    }
    catch {
        return "style guide";
    }
}
function applySuggestion(alert, line) {
    const { Name, Params } = alert.Action;
    if (Name !== "replace" && Name !== "remove") {
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
        let rest = after;
        if (alert.Check === "Grafana.Latin" &&
            rest.startsWith(".") &&
            !rest.startsWith("..")) {
            rest = rest.slice(1);
            if (rest.startsWith(" ") && !rest.startsWith(", ")) {
                rest = "," + rest;
            }
        }
        corrected = before + first + rest;
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
function formatComment(alert, line) {
    const reference = alert.Link
        ? `\n\nFor more information, refer to [${linkText(alert.Link)}](${alert.Link}).`
        : "";
    const suggestion = applySuggestion(alert, line);
    const suggestionBlock = suggestion ? `\n\n${suggestion}` : "";
    const issueTitle = encodeURIComponent(`Vale rule: ${alert.Check}`);
    const issueUrl = `https://github.com/grafana/writers-toolkit/issues/new?title=${issueTitle}`;
    const footer = "\n\n---\n_Reported by Vale using Grafana Writers' Toolkit style." +
        ` If you believe we can improve the rule, [report an issue](${issueUrl})._`;
    return `${commentMarker(alert.Check, alert.Match)}
**${alert.Check}** (${alert.Severity})

${alert.Message.trimEnd()}${suggestionBlock}${reference}${footer}`;
}
