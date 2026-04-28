import { describe, it } from "node:test";
import assert from "node:assert/strict";
import { applySuggestion, ValeAlert } from "./suggestion";

function makeAlert(
  check: string,
  actionName: "replace" | "remove" | "",
  params: string[] | null,
  match: string,
  span: [number, number],
): ValeAlert {
  return {
    Action: { Name: actionName, Params: params },
    Check: check,
    Line: 1,
    Link: "",
    Match: match,
    Message: "",
    Severity: "warning",
    Span: span,
  };
}

describe("applySuggestion", () => {
  describe("Grafana.Latin", () => {
    it("replaces e.g. mid-sentence and adds a comma", () => {
      const alert = makeAlert("Grafana.Latin", "replace", ["for example"], "e.g", [9, 11]);
      assert.equal(
        applySuggestion(alert, "foo bar e.g. baz"),
        "```suggestion\nfoo bar for example, baz\n```",
      );
    });

    it("does not add a comma when one already follows", () => {
      const alert = makeAlert("Grafana.Latin", "replace", ["for example"], "e.g", [9, 11]);
      assert.equal(
        applySuggestion(alert, "foo bar e.g., baz"),
        "```suggestion\nfoo bar for example, baz\n```",
      );
    });

    it("does not add a comma when e.g. ends the sentence", () => {
      const alert = makeAlert("Grafana.Latin", "replace", ["for example"], "e.g", [5, 7]);
      assert.equal(
        applySuggestion(alert, "See e.g."),
        "```suggestion\nSee for example\n```",
      );
    });

    it("replaces i.e. mid-sentence and adds a comma", () => {
      const alert = makeAlert("Grafana.Latin", "replace", ["that is"], "i.e", [5, 7]);
      assert.equal(
        applySuggestion(alert, "foo i.e. bar"),
        "```suggestion\nfoo that is, bar\n```",
      );
    });
  });

  describe("replace action", () => {
    it("replaces a word without adding a comma", () => {
      const alert = makeAlert("Grafana.Admin", "replace", ["administrator"], "admin", [9, 13]);
      assert.equal(
        applySuggestion(alert, "Contact admin for help."),
        "```suggestion\nContact administrator for help.\n```",
      );
    });

    it("preserves a sentence-ending period after a plain word replacement", () => {
      const alert = makeAlert("Grafana.SelfManaged", "replace", ["self-managed"], "on-prem", [5, 11]);
      assert.equal(
        applySuggestion(alert, "Use on-prem."),
        "```suggestion\nUse self-managed.\n```",
      );
    });

    it("uses the first param as the suggestion and lists alternatives", () => {
      const alert = makeAlert("Grafana.AllowsTo", "replace", ["allows you to", "makes it possible to"], "allows to", [6, 14]);
      assert.equal(
        applySuggestion(alert, "This allows to configure it."),
        "```suggestion\nThis allows you to configure it.\n```\n\nAlternatives: `makes it possible to`",
      );
    });

    it("returns null when there is no action", () => {
      const alert = makeAlert("Grafana.AltText", "", null, "", [1, 1]);
      assert.equal(applySuggestion(alert, "![](image.png)"), null);
    });
  });

  describe("remove action", () => {
    it("removes the matched word and collapses the double space", () => {
      const alert = makeAlert("Grafana.Simple", "remove", null, "simply", [4, 9]);
      assert.equal(
        applySuggestion(alert, "It simply works."),
        "```suggestion\nIt works.\n```",
      );
    });
  });
});
