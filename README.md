# A Kubernetes client tool written in Go 

A simple client tool for Kubernetes written in Go that lists 8 common resources (pods, deployments, services etc.) in a namespace.  


## Technologies Used :fireworks:

- Golang :heavy_check_mark:
- Kubernetes :heavy_check_mark:
- Docker :heavy_check_mark:
- Make utility :heavy_check_mark:


## Table of Contents :round_pushpin:

- [Key Features](#key-features)
- [Build and run the application outside of a Kubernetes cluster](#build-and-run-the-application-outside-of-a-Kubernetes-cluster)
- [Run the application inside a Kubernetes cluster](#run-the-application-inside-a-Kubernetes-cluster)


## Key Features :point_up:

- Lists the following resources in a namespace: 
    - pods
    - services
    - daemonsets
    - deployments
    - replicasets
    - statefulsets
    - jobs 
    - cronjobs
- Can run both as a container in a Kubernetes Cluster and outside of a Kubernetes Cluster
- If no namespace is initially provided, it shows the avaliable namespaces and waits for user input
- Provides further info about the state, creation time, and container image of each pod


## Build and run the application outside of a cluster :outbox_tray:  
1. Clone the repo
2. You will need an up and running Kubernetes cluster. You can deploy a Kubernetes cluster locally using [minikube](https://minikube.sigs.k8s.io/docs/start/)
3. Build and run the client tool outside of a Kubernetes Cluster


    - Define namespace as variable (NS=\<namespace>)

        ```
        make runoutside NS=kube-system
        ```

    - OR, Select one of the avaliable namespaces
        ```
        make runoutside 
        ```
        For example
        ```
        Available Namespaces =>
        [0] default
        [1] kube-node-lease
        [2] kube-public
        [3] kube-system

        -> Please select an id from the namespaces above.
        3
        ```

4. The output should be similar to the following

        Selected namespace: kube-system
        Current time:  2022-09-30 19:06:47.4787865 +0200 CEST m=+4.061903001


        *** Pods *** 
        In total there are 7 Pods in namespace kube-system : 

                                ---- Running Pods :)  ----
                Name   |        CreationTimeStamp      |        ContainerImage
        [1]     coredns-6d4b75cb6d-rtlkq        2022-09-29 11:00:00 +0200 CEST          k8s.gcr.io/coredns/coredns:v1.8.6
        [2]     etcd-minikube   2022-09-29 10:59:46 +0200 CEST          k8s.gcr.io/etcd:3.5.3-0
        [3]     kube-apiserver-minikube         2022-09-29 10:59:47 +0200 CEST          k8s.gcr.io/kube-apiserver:v1.24.1
        [4]     kube-controller-manager-minikube        2022-09-29 10:59:42 +0200 CEST          k8s.gcr.io/kube-controller-manager:v1.24.1
        [5]     kube-proxy-x465q        2022-09-29 11:00:00 +0200 CEST          k8s.gcr.io/kube-proxy:v1.24.1
        [6]     kube-scheduler-minikube         2022-09-29 10:59:46 +0200 CEST          k8s.gcr.io/kube-scheduler:v1.24.1
        [7]     storage-provisioner     2022-09-29 10:59:50 +0200 CEST          gcr.io/k8s-minikube/storage-provisioner:v5

                                ---- NON Running Pods  :(  ----
                                        none


        *** Services ***
        kube-dns


        *** Deployments ***
        coredns


        *** DaemonSets ***
        kube-proxy


        *** ReplicaSets ***
        coredns-6d4b75cb6d


        *** StatefulSets ***
        none


        *** Jobs ***
        none 


        *** CronJobs ***
        none




## Run the application inside a cluster :inbox_tray:
1. Clone the repo
2. You will need an up and running Kubernetes cluster. You can deploy a Kubernetes cluster locally using [minikube](https://minikube.sigs.k8s.io/docs/start/)
3. Run this tool as a container/pod in a Kuberentes cluster. 

   ```
    make runinside 
    ```

    The above command will:
    1. create the associated Kuberentes resources (namespace, service account, cluster role, cluster role binding and job) 
    2. run the client-tool in a container/pod in a seperate namespace and present the output 
    3. delete all the associated resources after the Kuberentes job completes


 
4. The output should be similar to the following

        * Deploy the associated Kubernetes resources and run the application as job. *
        make[1]: Entering directory `C:/.../go-k8s-client'
        job.batch/k8s-report condition met

        Available Namespaces =>
        [0] default
        [1] kube-node-lease
        [2] kube-public
        [3] kube-system
        [4] report

        Selected namespace: kube-system
        Current time:  2022-09-30 19:36:46.647571595 +0000 UTC m=+0.021822113


        *** Pods ***
        In total there are 7 Pods in namespace kube-system :

                                ---- Running Pods :)  ----
                Name   |        CreationTimeStamp      |        ContainerImage
        [1]     coredns-6d4b75cb6d-rtlkq        2022-09-29 09:00:00 +0000 UTC   k8s.gcr.io/coredns/coredns:v1.8.6
        [2]     etcd-minikube   2022-09-29 08:59:46 +0000 UTC   k8s.gcr.io/etcd:3.5.3-0
        [3]     kube-apiserver-minikube         2022-09-29 08:59:47 +0000 UTC   k8s.gcr.io/kube-apiserver:v1.24.1
        [4]     kube-controller-manager-minikube        2022-09-29 08:59:42 +0000 UTC   k8s.gcr.io/kube-controller-manager:v1.24.1
        [5]     kube-proxy-x465q        2022-09-29 09:00:00 +0000 UTC   k8s.gcr.io/kube-proxy:v1.24.1
        [6]     kube-scheduler-minikube         2022-09-29 08:59:46 +0000 UTC   k8s.gcr.io/kube-scheduler:v1.24.1
        [7]     storage-provisioner     2022-09-29 08:59:50 +0000 UTC   gcr.io/k8s-minikube/storage-provisioner:v5

                                ---- NON Running Pods  :(  ----
                                        none


        *** Services ***
        kube-dns


        *** Deployments ***
        coredns


        *** DaemonSets ***
        kube-proxy


        *** ReplicaSets ***
        coredns-6d4b75cb6d


        *** StatefulSets ***
        none


        *** Jobs ***
        none


        *** CronJobs ***
        none

        make[1]: Leaving directory `C:/Users/imavridis/Desktop/go-k8s-client'

        * Report produced, associated Kubernetes resources cleaned up. *
       

- Note! You can investigate different namespaces by changing the K8S_NAMESPACE environment variable
