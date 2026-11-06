# Exercise 1.3: Run a Docker container with a Microservice

In this exercise, you will run a container from the before built image.

## Instructions

1. Run the container from the image and expose the container port: **8888** to the host port: **9090**.

    ```console
    docker container run -p 9090:8888 [YOUR-DOCKERHUB-ACCOUNT]/hello-app:0.0.1
    ```

    * *Optional*: Run the container in *detached mode* using the option `--detach` or `-d`. This runs the Docker container in the background of your terminal. It does not receive input or display output.

    ```console
    docker container run -d -p 9090:8888 [YOUR-DOCKERHUB-ACCOUNT]/hello-app:0.0.1
    ```

2. Open a browser and go to: http://localhost:9090

3. See your container running on your local Docker daemon:

    ```console
    docker ps
    ```
    
    ```
    CONTAINER ID        IMAGE                    COMMAND             CREATED             STATUS              PORTS                    NAMES
    789d08da1704        xyz/hello-app:0.0.1        "/usr/hello-app"        21 seconds ago      Up 19 seconds       0.0.0.0:9999->8888/tcp   foxy_joliot
    ```

4. Stop your container:

    ```console
    docker stop 789d08da1704
    ```
