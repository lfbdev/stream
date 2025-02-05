## 1 `github-repo-scaffolding-golang` Plugin

This plugin bootstraps a GitHub repo with scaffolding code for a Golang web application.

_This plugin depends on the following environment variable:_

- GITHUB_TOKEN

Set it before using this plugin.

If you don't know how to create this token, check out:
- [Creating a personal access token](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token)

*Tips:*

*1. If you run `dtm delete`, the repo on GitHub will be completely removed.*

*2. If the `Update` interface is called, the repo on github will be completely removed and recreated. However, given our current implementation, this interface shall not be called, as of in v0.2.0.*

## 2 Usage

**Please note that the `owner` parameter is case-sensitive.**

```yaml
tools:
# name of the instance with github-repo-scaffolding-golang
- name: go-webapp-repo
  plugin:
    # kind of the plugin
    kind: github-repo-scaffolding-golang
    # version of the plugin
    # checkout the version from the GitHub releases
    version: 0.3.0
  # options for the plugin
  options:
    # the repo's owner. It should be case-sensitive here; strictly use your GitHub user name; please change the value below.
    owner: YOUR_GITHUB_USERNAME
    # the repo which you'd like to create; please change the value below.
    repo: YOUR_REPO_NAME
    # the branch of the repo you'd like to hold the code
    branch: main
    # the image repo you'd like to push the container image; please change the value below.
    image_repo: YOUR_DOCKERHUB_USERNAME/YOUR_DOCKERHUB_IMAGE_REPO_NAME
```

Replace the following from the config above:

- `YOUR_GITHUB_USERNAME`
- `YOUR_REPO_NAME`
- `YOUR_DOCKERHUB_USERNAME`
- `YOUR_DOCKERHUB_IMAGE_REPO_NAME`

The "branch" in the example above is "main", but you can adjust accordingly.

Currently, all the parameters in the example above are mandatory.

## 3. Outputs

This plugin has three outputs:

- `owner`
- `repo`
- `repoURL` (example: "https://github.com/IronCore864/test.git")

If, for example, you want to use the outputs as inputs for another plugin, you can refer to the following example:

```yaml
---
tools:
- name: go-webapp-repo
  plugin:
    kind: github-repo-scaffolding-golang
    version: 0.3.0
  options:
    owner: IronCore864
    repo: go-webapp-devstream-demo
    branch: main
    image_repo: ironcore864/go-webapp-devstream-demo
- name: golang-demo-actions
  plugin:
    kind: githubactions-golang
    version: 0.3.0
  dependsOn: ["go-webapp-repo.github-repo-scaffolding-golang"]
  options:
    owner: ${{go-webapp-repo.github-repo-scaffolding-golang.outputs.owner}}
    repo: ${{go-webapp-repo.github-repo-scaffolding-golang.outputs.repo}}
    language:
      name: go
      version: "1.17"
    branch: main
    build:
      enable: True
    test:
      enable: True
      coverage:
        enable: True
    docker:
      enable: False
```

Pay attention to the `${{ xxx }}` part in the example. `${{ TOOL_NAME.TOOL_KIND.outputs.var}}` is the syntax for using an output.
