#!/usr/bin/env bash
set -euo pipefail

TAG="v1.21.3"
MESSAGE="OpenShift Pipelines 1.21.3 release"
DRY_RUN="${DRY_RUN:-true}"

declare -A REPOS=(
  ["openshift-pipelines/p12n-console-plugin"]="c6c8eb76ab9b2cafc5752f1446338d6fd8944548"
  ["openshift-pipelines/p12n-console-plugin-pf5"]="ba85d3c7247038b4e3a0d0a013eb029b300e2ba7"
  ["openshift-pipelines/p12n-manual-approval-gate"]="e25d64a2cfc4a941af0111f4e02096783370f02b"
  ["openshift-pipelines/p12n-opc"]="938f533cf53c62b99879ea7579d2c0658077e96b"
  ["openshift-pipelines/p12n-tekton-caches"]="a582ebcda617b1e812ff99c90ca99ea9734077c2"
  ["openshift-pipelines/pac-downstream"]="cadb77a6ee6ec3f134d3fbf6ee17bad39c777cd6"
  ["openshift-pipelines/serve-tkn-cli"]="b01eedb194a8ccb0ed1f5d2b3d6e3bf8d5772d19"
  ["openshift-pipelines/tektoncd-chains"]="6149030b735e2e0574706eb1e480f37932eb0f05"
  ["openshift-pipelines/tektoncd-cli"]="dd1b37cfe2d2ebca8a624e0485b83d4616f3d20f"
  ["openshift-pipelines/tektoncd-git-clone"]="cca25ea35394a0493dca1fc895334a5912308e70"
  ["openshift-pipelines/tektoncd-hub"]="1109073ab8ce4c5348fd2ec7f5dca017e59cda08"
  ["openshift-pipelines/tektoncd-pipeline"]="caa67fad818a02244b22a94ecfef3fc02d4170ea"
  ["openshift-pipelines/tektoncd-pruner"]="5d09825b776b870d9bcb4f54f286cede781968bb"
  ["openshift-pipelines/tektoncd-results"]="989e3e4614721a5b1b0ee4d7b0cee14f69f9e8ac"
  ["openshift-pipelines/tektoncd-triggers"]="a4018bc83f716c553764a4c22b24fe67007a3409"
  ["openshift-pipelines/operator"]="070d2518d5afe7a9ce8e49eb294c254e0a37d971"
)

for repo in "${!REPOS[@]}"; do
  sha="${REPOS[$repo]}"
  echo "--- ${repo} @ ${sha:0:12}"

  if [[ "$DRY_RUN" == "true" ]]; then
    echo "  [dry-run] would create tag ${TAG} at ${sha:0:12}"
    continue
  fi

  # Create annotated tag object
  tag_sha=$(gh api "repos/${repo}/git/tags" \
    -f tag="${TAG}" \
    -f message="${MESSAGE}" \
    -f object="${sha}" \
    -f type="commit" \
    --jq '.sha')

  # Create the ref pointing to the tag object
  gh api "repos/${repo}/git/refs" \
    -f ref="refs/tags/${TAG}" \
    -f sha="${tag_sha}" \
    --silent

  echo "  ✓ tagged ${TAG}"
done

echo ""
if [[ "$DRY_RUN" == "true" ]]; then
  echo "Dry run complete. Run with DRY_RUN=false to create tags."
else
  echo "All repos tagged with ${TAG}."
fi
