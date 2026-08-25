package main

import (
	"os"
	"strings"
	"testing"
)

func TestDeployPreviewWorkflowUsesOneImmutableGARImageForDeployAndCheck(t *testing.T) {
	content, err := os.ReadFile("../../.github/workflows/deploy-preview.yml")
	if err != nil {
		t.Fatal(err)
	}
	workflow := string(content)

	required := []string{
		"Build and push preview image to GAR",
		"Authenticate to GAR for deploy previews",
		"github-docs-deploy-previews@grafanalabs-workload-identity.iam.gserviceaccount.com",
		"docker/login-action@",
		"docker/metadata-action@",
		"docker/setup-buildx-action@",
		"docker/build-push-action@",
		"us-docker.pkg.dev/grafanalabs-dev/docker-docs-previews-dev/${{ env.REPO }}",
		"push: true",
		`run: echo "image=${IMAGE}@${DIGEST}" >> "$GITHUB_OUTPUT"`,
		"image: ${{ needs.build-website.outputs.preview_image }}",
		"PREVIEW_IMAGE: ${{ needs.build-website.outputs.preview_image }}",
		"LINK_CHECKER_IMAGE: us-docker.pkg.dev/",
		"Check links in the preview image",
		"Authenticate to GAR for the preview image",
		"Pull the immutable preview image",
		"Pull the link checker image",
		"Check out deploy preview helpers",
		"repository: grafana/writers-toolkit",
		"./deploy-preview-files/deploy-preview/check-links",
		`"${LINK_CHECKER_IMAGE}"`,
		"docker run --pull never --rm",
		"changes",
		"comment",
		"-changed-files-json changed-files.json",
		`-repo "$PR_REPO"`,
		`-pr-number "$PR_NUMBER"`,
	}
	for _, fragment := range required {
		if !strings.Contains(workflow, fragment) {
			t.Errorf("deploy-preview workflow missing immutable-image contract fragment %q", fragment)
		}
	}

	if strings.Contains(workflow, "actions/download-artifact") {
		t.Error("deploy-preview workflow must not download the Hugo dist artifact")
	}
	if strings.Contains(workflow, "name: dist") || strings.Contains(workflow, "path: dist") {
		t.Error("deploy-preview workflow must not upload the Hugo dist artifact")
	}
	if strings.Contains(workflow, "actions/upload-artifact@") || strings.Contains(workflow, "name: Upload links report") {
		t.Error("deploy-preview workflow must not upload a link report artifact")
	}
	if strings.Contains(workflow, "repository: grafana/website") || strings.Contains(workflow, "scripts/docs/link-checker") {
		t.Error("deploy-preview workflow must use the published checker image because grafana/website is private")
	}
	if strings.Contains(workflow, "./link-checker") {
		t.Error("deploy-preview workflow must not build or run a checker binary from private website source")
	}
	if strings.Contains(workflow, "PREVIEW_IMAGE: ${{ needs.deploy-preview.outputs.preview-image }}") {
		t.Error("the link check must pull the immutable image produced by the build job directly")
	}
	for _, command := range []string{"docker cp", "docker create", "docker exec", "docker logs", "docker rm", "docker volume"} {
		if strings.Contains(workflow, command) {
			t.Errorf("the workflow must leave preview lifecycle management to the check-links helper; found %q", command)
		}
	}
	if strings.Contains(workflow, "git cat-file blob") || strings.Contains(workflow, "git fetch --no-tags --depth=1") {
		t.Error("the workflow must leave base commit and blob materialization to the changes command")
	}
	if strings.Contains(workflow, "gh api --paginate") || strings.Contains(workflow, "name: Determine changed files\n") {
		t.Error("the workflow must leave pull request file collection to the changes command")
	}
}

func TestCheckLinksRunsPublishedCheckerBesidePreview(t *testing.T) {
	content, err := os.ReadFile("../check-links")
	if err != nil {
		t.Fatal(err)
	}
	helper := string(content)

	required := []string{
		`"${PREVIEW_IMAGE}"`,
		`"${LINK_CHECKER_IMAGE}"`,
		"docker volume create",
		"docker run --pull never --detach",
		`--network "container:${preview_container}"`,
		"target=/usr/share/nginx,readonly",
		"check",
		"-existing-server",
		"-nginx-port 80",
		"-dist-root /usr/share/nginx/dist",
		"-changed-files-json /work/changed-files.json",
		"-output /work/links.json",
		"docker rm --force",
		"docker volume rm --force",
	}
	for _, fragment := range required {
		if !strings.Contains(helper, fragment) {
			t.Errorf("check-links helper missing sidecar contract fragment %q", fragment)
		}
	}

	if strings.Contains(helper, "docker cp") {
		t.Error("check-links must share the preview artifact through a volume, not copy it out of the image")
	}
}

func TestDeployPreviewImageContainsGeneratedRedirects(t *testing.T) {
	dockerfile, err := os.ReadFile("../Dockerfile")
	if err != nil {
		t.Fatal(err)
	}
	nginxConfig, err := os.ReadFile("../nginx.conf")
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(string(dockerfile), "COPY dist/redirects.conf /etc/nginx/redirects.conf") {
		t.Error("deploy-preview image does not copy the generated redirect configuration")
	}
	if !strings.Contains(string(dockerfile), "COPY ./dist /usr/share/nginx/dist") {
		t.Error("deploy-preview image does not preserve its existing document root")
	}
	if strings.Contains(string(dockerfile), "LINK_CHECKER_IMAGE") || strings.Contains(string(dockerfile), "broken-links") {
		t.Error("deploy-preview image must not contain the link checker")
	}
	if !strings.Contains(string(nginxConfig), "include /etc/nginx/redirects.conf;") {
		t.Error("deploy-preview nginx does not load the generated redirect configuration")
	}
	if !strings.Contains(string(nginxConfig), "alias /usr/share/nginx/dist/;") {
		t.Error("deploy-preview nginx does not preserve its existing document root")
	}
}
