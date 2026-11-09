# Exercise 2.1: Setup Continuous Integration

In this exercise, you will set up a *Continuous Integration* (CI) workflow using GitHub as the version control system and GitHub Actions to run the checks. The workflow runs the unit tests and compiles the code when a change is pushed or a pull request is opened. The compiled binary is not saved after the workflow finishes; Exercise 2.2 publishes a container image to Docker Hub.

## Setup

![GitHub, GitHub Actions, DockerHub](./assets/lab_setup.png)

## Requirements

* GitHub account :octocat: 
    <details><summary>Click here for sign-up instructions.</summary>
    <p>

    To sign up:  https://github.com/join

    </p>
    </details>

## Instructions

1. Create a GitHub repository called `mini-ci-example`

2. Add and commit `main.go`, `main_test.go`, and `go.mod` to your repository.

3. In GitHub, go to **Actions** and set up a GitHub workflow by clicking `set up a workflow yourself`:

![GitHub Action activation](./assets/gh_action_activate.png)

1. Create a file `main.yml` in the folder `.github/workflows` of your repository and copy-paste the content from the provided `main.yml` file. 

```yaml
name: CI

# Controls when the workflow will run
on:
  # Triggers the workflow on push or pull request events but only for the main branch
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]

  # Allows you to run this workflow manually from the Actions tab
  workflow_dispatch:

# A workflow run is made up of one or more jobs that can run sequentially or in parallel
jobs:
  # This workflow contains a single job called "continuous_integration"
  continuous_integration:
    # The type of runner that the job will run on
    runs-on: ubuntu-latest

    # Steps represent a sequence of tasks that will be executed as part of the job
    steps:
      # Checks-out your repository under $GITHUB_WORKSPACE, so your job can access it
      - uses: actions/checkout@v7

      - name: Set up Go
        uses: actions/setup-go@v7
        with:
          go-version-file: go.mod

      # Runs a single command using the runners shell
      - name: Run a one-line script
        run: echo Starting tests and build

      # Runs a set of commands using the runners shell
      - name: Run test and build
        run: |
          go test -v ./...
          CGO_ENABLED=0 GOARCH=amd64 go build -o demo
```

1. Trigger the GitHub Action build by a code change, e.g., change "Hello" to "Hi there" in the file `./main.go`

1. Watch GitHub Actions run your tests and compile the binary.

    :mag: What is your observation? 

