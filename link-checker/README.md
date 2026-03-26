How the link checker works:

1. Download the built website artifact.
2. A Go script walks the dist dir to build a list of pages.
3. Nginx is configured, tested, and spun up.
4. Each page of the build is crawled and every link is checked if it hasn't already been checked.
5. A comment is added to the PR describing the broken links and where they occur (see screenshot). The comment is updated on subsequent runs.
