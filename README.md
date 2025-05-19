# R2D2 - Release Request Detection and Deployment

The R2D2 system allows tags to be created from PRs.

# Required Environment Variables

- COSMOS_CHANGELOG: The changelog file for the module being updated. This should contain the tag you want to update to.
- COSMOS_CHANGELOG_FILE: The path to the changelog file. This is used to derive the package path. For example, tagging v1.0.0 with changelog at `x/tx/changelog.md` will derive the tag `x/tx/v1.0.0`
- COSMOS_TAGS: The list of tags from git. 
- COMMIT_SHA: The sha of the commit we want to tag.
- GITHUB_TOKEN: Auth token for creating tags.
- GITHUB_REPOSITORY_OWNER: The name of the github owner/org.
- GITHUB_REPOSITORY: The name of the github repo.