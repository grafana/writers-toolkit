"""Exercise preview generation without GitHub credentials or Docker.

Run: python3 deploy-preview/test_build_interactive_tutorials.py
Requires bash, git, and jq. Only external checkout and Docker execution are stubbed.
"""

import json
import os
from pathlib import Path
import subprocess
import tempfile
import unittest


SCRIPT = Path(__file__).resolve().with_name("build-interactive-tutorials")


class PreviewTest(unittest.TestCase):
    def setUp(self):
        self.temp = tempfile.TemporaryDirectory()
        self.addCleanup(self.temp.cleanup)
        self.root = Path(self.temp.name)
        self.source = self.root / "fixtures"
        self.source.mkdir()
        subprocess.run(["git", "init", "-q", str(self.source)], check=True)
        self.bin = self.root / "bin"
        self.bin.mkdir()
        self.stub("gh", """#!/usr/bin/env python3
import os, shutil, sys
if sys.argv[1:3] == ['repo', 'clone']:
    shutil.copytree(os.environ['FIXTURES'], sys.argv[-1])
""")
        self.stub("docker", """#!/usr/bin/env python3
import os, shutil, sys
from pathlib import Path
for i, arg in enumerate(sys.argv):
    if arg == '--volume':
        source, target = sys.argv[i + 1].split(':', 1)
        if target in ('/hugo/content/docs/learning-paths', '/hugo/data/docs/interactive-learning'):
            name = 'content' if '/content/' in target else 'data'
            shutil.copytree(source, Path(os.environ['CAPTURE']) / name)
""")

    def stub(self, name, content):
        path = self.bin / name
        path.write_text(content)
        path.chmod(0o755)

    def package(self, name):
        path = self.source / name
        path.mkdir(parents=True)
        (path / "content.json").write_text(json.dumps({"title": name, "blocks": []}))
        # Intentionally no trailing newline: regression covered by PR #1470.
        (path / "website.yaml").write_text("step: 1")
        (path / "manifest.json").write_text(json.dumps({"description": "A path"}))
        return path

    def run_preview(self):
        subprocess.run(["git", "-C", str(self.source), "add", "."], check=True)
        env = dict(os.environ, PATH=f"{self.bin}{os.pathsep}{os.environ['PATH']}",
                   BRANCH="test", GH_TOKEN="test-only", FIXTURES=str(self.source),
                   CAPTURE=str(self.root / "capture"), WHATS_NEW_DIR=str(self.root))
        return subprocess.run(["bash", str(SCRIPT)], cwd=self.root, env=env,
                              capture_output=True, text=True)

    def test_both_suffixes_and_inherited_settings(self):
        for name in ("legacy-lj", "renamed-lp"):
            self.package(name)
            self.package(f"{name}/first-step")
        result = self.run_preview()
        self.assertEqual(result.returncode, 0, result.stderr)
        content = self.root / "capture/content"
        for slug, package in (("legacy", "legacy-lj"), ("renamed", "renamed-lp")):
            page = (content / slug / "_index.md").read_text()
            self.assertIn(f"pathfinder_data: {package}\n", page)
            self.assertIn("step: 1\ntype: docs\n", page)
            self.assertIn(f"pathfinder_data: {package}/first-step\n",
                          (content / slug / "first-step/index.md").read_text())
            self.assertTrue((self.root / "capture/data" / package / "content.json").is_file())
        index = (content / "_index.md").read_text()
        self.assertIn("layout: single-journey", index)
        self.assertIn("learn_microsite_chrome: true", index)
        self.assertIn("outputs: [HTML, json, pathfinder, unstyled]", index)
        self.assertIn("cta:\n  defaults:", index)

    def test_conflicting_slugs_fail_instead_of_overwriting(self):
        self.package("same-lj")
        self.package("same-lp")
        result = self.run_preview()
        self.assertNotEqual(result.returncode, 0)
        self.assertIn("multiple learning path packages use the slug 'same'", result.stderr)
        self.assertFalse((self.root / "capture").exists())

    def test_packages_without_website_metadata_are_not_published(self):
        path = self.package("hidden-lp")
        (path / "website.yaml").unlink()
        result = self.run_preview()
        self.assertEqual(result.returncode, 0, result.stderr)
        self.assertFalse((self.root / "capture/content/hidden").exists())
        self.assertTrue((self.root / "capture/content/_index.md").is_file())


if __name__ == "__main__":
    unittest.main()
