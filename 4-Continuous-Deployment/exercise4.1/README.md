# Exercise 4.1: Performing a Rolling Update

In this exercise, you will learn how to roll out a new version of our application using
a **rolling update** strategy. Rolling updates replace Pods incrementally. With enough healthy replicas and capacity, this can keep the application available during an update.

To reduce the risk of an availability gap, we will add a readiness probe as well as a lifecycle hook:
- The **readiness probe** checks whether our application is ready to handle the traffic.
- The lifecycle hook **preStop** runs before Kubernetes sends the termination signal to the container. A short wait gives Kubernetes and load balancers time to update their endpoints, but that propagation is asynchronous, so the wait cannot guarantee that no requests reach a terminating Pod.

## Instructions

1. To prepare your Deployment for a rolling update that can maintain application availability, specify the following fields:
    - `.spec.replicas`: Specifies the desired number of replicas. Use multiple replicas (e.g. 5) so healthy existing Pods can continue serving traffic while new Pods start.
    - `.spec.strategy.type`: Specifies the strategy used to replace old Pods by new ones. `.spec.strategy.type` can be "Recreate" or "RollingUpdate".
    - `.spec.strategy.rollingUpdate.maxUnavailable`: Specifies the maximum number of Pods that can be unavailable during the update process. With 5 replicas and a value of `1`, Kubernetes can keep at least 4 available while updating, provided the remaining Pods are healthy.
    - `.spec.strategy.rollingUpdate.maxSurge`: Specifies the maximum number of Pods that can be created over the desired number of Pods.
    - `.spec.template.spec.containers[].readinessProbe`: Specifies a readiness probe for checking whether an application container is ready to handle traffic.
    - `.spec.template.spec.containers[].lifecycle.preStop`: Specifies a lifecycle hook that must complete before Kubernetes sends the termination signal to the container.


    ```yaml
    apiVersion: apps/v1
    kind: Deployment
    metadata:
      name: demo
    spec:
      replicas: 5
      selector:
        matchLabels:
          app: demo
      strategy:
        type: RollingUpdate
        rollingUpdate:
          maxSurge: 1
          maxUnavailable: 1
      template:
        metadata:
          labels:
            app: demo
        spec:
          containers:
          - name: demo
            image: agrimmer/demo:plain
            ports:
            - containerPort: 8888
            readinessProbe:
              httpGet:
                path: /
                port: 8888
              initialDelaySeconds: 5
            lifecycle:
              preStop:
                exec:
                  command: ["sleep", "15"]
    ```

1. Save this manifest as `deployment.yaml` and apply it to update the existing `demo` Deployment:

    ```console
    kubectl apply -f deployment.yaml
    ```

1. Watch the updating of the Pods in a terminal:

    ```console
    watch kubectl get pods
    ```

1. In a second terminal, find your Service's external address:

    ```console
    kubectl get service demo
    ```

    If the Service is a LoadBalancer, replace `<EXTERNAL-ADDRESS>` below with its `EXTERNAL-IP` or hostname, then poll the application:

    ```console
    while true; do sleep 1; curl --fail http://<EXTERNAL-ADDRESS>; echo; done
    ```

1. Update the image of your deployment:

    ```console
    kubectl set image deployments/demo demo=agrimmer/demo:html
    ```

:rocket: Watch the Pods roll out and observe the application responses. With healthy replicas and sufficient capacity, requests should continue to be served during the update.
