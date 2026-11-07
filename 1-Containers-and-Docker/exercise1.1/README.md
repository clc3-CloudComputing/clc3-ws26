# Exercise 1.1: Working with Docker and Dockerfile

In this exercise, you will write a Dockerfile and you will build an image from this Dockerfile. For more information about a Dockerfile, click [here](https://docs.docker.com/engine/reference/builder/).

## Requirements

* [Docker](https://www.docker.com/) installed: 
    <details><summary>Click here for install instructions.</summary>
    <p>

    * Docker for Windows: https://docs.docker.com/docker-for-windows/install/

    * Docker for Mac: https://docs.docker.com/docker-for-mac/install

    </p>
    </details>

* [Docker Hub](https://hub.docker.com/) Account created:
    <details><summary>Click here for sign-up instructions.</summary>
    <p>

    * To sign up: https://app.docker.com/signup

    </p>
    </details>

## Instructions

1. Create a Dockerfile with the following instructions:
    * Base image: `golang:1.27.3-alpine3.24`
    * Set a maintainer label: `LABEL maintainer="<YOUR-EMAIL>"`
    * Set working directory: `/opt`
    * Copy local file `go.mod` and `main.go` to the image folder `/opt/`
    * List items in the working directory (`ls -lsa`)
    * Build the app (`CGO_ENABLED=0 go build -o /usr/hello-app`)
    * Expose the port `8888`
    * Run the app

2. Build a Docker image based on the Dockerfile:
    * Image tag: `[YOUR-DOCKERHUB-ACCOUNT]/hello-app:0.0.1`

    ```console
    docker image build -f Dockerfile -t [YOUR-DOCKERHUB-ACCOUNT]/hello-app:0.0.1 ./  --progress=plain
    ```

3. Execute the command a second time: 

    ```console
    docker image build -f Dockerfile -t [YOUR-DOCKERHUB-ACCOUNT]/hello-app:0.0.1 ./
    ```

    :mag: What is your observation? 
  
4. List the images stored in your local Docker image store:

    ```console
    docker images
    ```
