# Deploy to GitHub Pages

Both project shapes deploy the same way: build, then let GitHub Pages serve
the output directory from a branch.

## 1. Choose the hosting shape

GitHub Pages serves one of two sources:

- **From `main`** — for a landing page whose build output lives at the
  repository root (the `cmd/site` shape with `runvil build --output .`).
- **From `gh-pages`** — for a documentation site whose `site/` output is
  copied to a clean commit on the `gh-pages` branch.

## 2. Deployment from gh-pages

Rebuild the orphan branch command by command:

```sh
# build
runvil build

# publish site/ as a clean gh-pages commit
git branch -D gh-pages
git checkout --orphan gh-pages
git rm --cached -r .
git clean -fdx -e site
cp -r site/* .
git add .
git commit -m "Publish"
git push --force origin gh-pages
```

Set the Pages source to `gh-pages` in the repository settings.

## 3. Verify

Curl the live URL to confirm the chapters, assets, and breadcrumbs resolve:

```sh
curl -sfI https://runvil.github.io/docs/ | head -1
curl -sf  https://runvil.github.io/docs/chapters/tutorials/ | grep -o '<title>[^<]*</title>'
```

Every generated link carries the configured base path (`/docs/...`), so the
site is fully functional at its deployed location.