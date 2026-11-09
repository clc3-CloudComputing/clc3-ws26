# Exercise 2.2: Build and push a Container Image to DockerHub

In this exercise, you will integrate Docker Hub into the CI workflow to push an image to your Docker Hub account. Add a Dockerfile to the repository and extend the GitHub Actions workflow with steps to build, tag, and push the image.

To follow the steps below, start with the workflow from Exercise 2.1 and add the steps in order.

## Requirements

* Docker Hub account
    <details><summary>Click here for sign-up instructions.</summary>
    <p>

    To sign up: https://app.docker.com/signup

    </p>
    </details>

* Exercise 2.1 completed

## Instructions

1. Setup credentials for DockerHub in GitHub:

    1. In your repository, go to `Settings` > `Secrets and variables` > `Actions` > **New repository secret**

    1. Add the secret `DOCKERHUB_USERNAME` with your Docker Hub user name

    1. Create a Docker Hub personal access token with read and write permissions, then add it as the `DOCKERHUB_TOKEN` secret. Do not use your account password.

1. Add and commit the provided `Dockerfile` to your repository.

1. Extend the `main.yml` workflow with the following step. Note: This is *not required but recommended* using it to be able to build multi-platform images:

    ```yaml
    # Set up a Docker builder (the action uses the docker-container driver by default)
    - name: Set up Docker Buildx
      if: github.event_name != 'pull_request'
      uses: docker/setup-buildx-action@v4
    ```


1. Extend the `main.yml` workflow with a Docker `login` step:

    ```yaml
    # Login to DockerHub account
    - name: Login to DockerHub
      if: github.event_name != 'pull_request'
      uses: docker/login-action@v4
      with:
        username: ${{ secrets.DOCKERHUB_USERNAME }}
        password: ${{ secrets.DOCKERHUB_TOKEN }}
    ```
    
1. Extend the `main.yml` workflow with a Docker `build-push` step. Please change the [user] to your DockerHub account name: 

    ```yaml
    # Build a Docker image based on the provided Dockerfile
    - name: Build and push
      id: docker_build
      if: github.event_name != 'pull_request'
      uses: docker/build-push-action@v7
      with:
        context: .
        push: true
        tags: ${{ secrets.DOCKERHUB_USERNAME }}/app:latest
    ```

1. Finally, trigger a build by a code change. 

1. Watch GitHub Actions run your tests, build the image, and push it to Docker Hub. Docker setup, login, and push steps are skipped for pull requests so forked pull requests can run the tests without Docker Hub secrets.

:mag: Go to your DockerHub account and find your image. Can you find it?


### Use Git SHA instead of `latest`

It is not good practice to use the version `latest` for an image tag. Instead, make it as specific as possible, e.g., by using the Git commit SHA instead: 

1. Extend the `main.yml` workflow with a step that derives the short Git commit SHA and stores it in a variable: 

    ```yaml
    # Declare variables to store branch name and short Git commit SHA
    - name: Declare variables
      id: vars
      shell: bash
      run: |
        echo "sha_short=$(git rev-parse --short HEAD)" >> $GITHUB_OUTPUT
    ```

1. Change the `Build and push` step to reference the `sha_short` variable for the tag version: 

    ```yaml
        tags: ${{ secrets.DOCKERHUB_USERNAME }}/app:${{ steps.vars.outputs.sha_short }}
    ```

1. Finally, trigger a build by a code change. 
