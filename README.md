# R2D2—Release Request Detection and Deployment

The R2D2 system allows tags to be created from PRs.

This action is built specifically for Cosmos SDK release practices. Please open an issue if you'd like it generalized to support your use case.

# Cosmos SDK Usage

A PR should be made to update the changelog of the package that needs the new tag. 
The changelog should include the new version you're trying to tag.

For example:

`x/tx/changelog.md`
```markdown
# Changelog

## [Unreleased]

## [v2.0.0](https://github.com/cosmos/cosmos-sdk/releases/tag/x/tx/v2.0.0) - 2025-05-20
```

A PR with this changelog update will push the tag `x/tx/v2.0.0`.

# Required Environment Variables (Using Docker Directly)

- COSMOS_CHANGELOG: The changelog file for the module being updated. This should contain the tag you want to update to.
- COSMOS_CHANGELOG_FILE: The path to the changelog file. This is used to derive the package path. For example, tagging v1.0.0 with changelog at `x/tx/changelog.md` will derive the tag `x/tx/v1.0.0`
- COSMOS_TAGS: The list of tags from git. 
- COMMIT_SHA: The sha of the commit we want to tag.
- GITHUB_TOKEN: Auth token for creating tags.
- REPO_OWNER: The name of the GitHub owner/org.
- REPO_NAME: The name of the GitHub repo.
- R2D2_COMMENT_ONLY(Boolean): If the job should make a comment on the PR. if this is true, a tag will NOT be created.
- R2D2_PR_NUMBER: The PR number.

# Required Inputs (Using Action File)
- changelog: the flattened changelog file content
- changelog-file: the path to the changelog file (i.e., x/tx/changelog.md)
- cosmos-tags: the previous tags starting from the merge commit sha.
- r2d2-comment-only: boolean indicating whether the bot should simply comment on the PR. when false, the bot will push tags. (TODO: this is dumb)