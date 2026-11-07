# Exercise 1.3: Run a Docker container with a Microservice

In this exercise, you will run a container from an image built in Exercise 1.1 or pulled in Exercise 1.2.

## Instructions

1. Run an image from your account or the classmate's account whose image you pulled. Map container port **8888** to host port **9090**.

    ```console
    docker container run -p 9090:8888 [IMAGE-OWNER-DOCKERHUB-ACCOUNT]/hello-app:0.0.1
    ```

    * *Optional*: Run the container in *detached mode* using the option `--detach` or `-d`. This runs the Docker container in the background of your terminal. It does not receive input or display output.

    ```console
    docker container run -d -p 9090:8888 [IMAGE-OWNER-DOCKERHUB-ACCOUNT]/hello-app:0.0.1
    ```

2. Open a browser and go to: http://localhost:9090

3. See your container running on your local Docker daemon:

    ```console
    docker ps
    ```
    
    ```
    CONTAINER ID   IMAGE                                     COMMAND           CREATED         STATUS         PORTS                    NAMES
    <container-id> [YOUR-DOCKERHUB-ACCOUNT]/hello-app:0.0.1   "/usr/hello-app"  21 seconds ago  Up 19 seconds  0.0.0.0:9090->8888/tcp   foxy_joliot
    ```

4. If you ran the container in the foreground, press Ctrl+C in that terminal to stop it. If you used detached mode, stop it using the container ID shown by `docker ps`:

    ```console
    docker stop <container-id>
    ```
