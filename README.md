# hw2-Docker

## Part-1
Note that:
- **Dockerfile** and **start.go** refer to the Dockerfile and source code.  
- I have finished all basic tasks of the part-1. The cluster name where I deployed the container using Cloud Run is "hw2-14848" as the following stops show.
- I have finished the extra-credit task. I have deployed the container to the Kubernetes cluster as the following step 7 shows.  

Steps:  
1. Created a local docker image.  
2. Pushed the local docker image to my docker hub.  
3. Logged in the GCP and pulled the docker image from my own docker hub.  
4. Pushed the docker image to Container Registry in the GCP.  
5. Started a service using Cloud Run and deployed the docker image to it. When the docker started, the console showed "Hello World!".  
> as **part-1/service-1.png** and **part-1/service-2.png** shows
6. The cluster information of the running container was taken as the screenshot. The log showed "Hello World!".  
> as **part-1/service-3.png** shows. The cluster name "hw2-14848" is included in the image.  
7. Deployed the docker image to the **Kubernetes cluster**. The console of the Kubernetes cluster outputted the "Hello World!".  
> as **part-1/k8s-cluster-1.png** and **part-1/k8s-cluster-2.png** shows  


## Part-2
Steps:  
1. Pulled the jupyter-notebook docker images and run it in my local computer.  
2. Went into the shell of the container and input the "jupyter-notebook" to start the service.  
3. The docker was running and I could see the web in the address http://localhost:8888.  
> as **part-2/part2.png** shows
