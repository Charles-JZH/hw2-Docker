# hw2-Docker

## Part-1 steps:  
1. Created a local docker image.  
2. Pushed the local docker image to my docker hub.  
3. Logged in the GCP and pulled the docker image from my own docker hub.  
4. Pushed the docker image to Container Registry in the GCP.  
5. Started a service using Cloud Run and deployed the docker image to it.  
(as **part-1/service-1.png** and **part-1/service-2.png** shows)  
7. The log and console showed "Hello World!".  
(as **part-1/service-2.png** and **part-1/service-3.png** shows)  
9. Created a Kubernetes cluster in GCP.  
(as **part-1/k8s-cluster-1.png** shows)  
11. Deployed the docker image to the **Kubernetes cluster**. The console of the Kubernetes cluster outputted the "Hello World!".  
(as **part-1/k8s-cluster-2.png** shows)  
The Dockerfile and source code has also been uploaded to this Github repository already as **Dockerfile** and **start.go**.


## Part-2 steps:  
1. Pulled the jupyter-notebook docker images and run it in my local computer.  
2. Went into the shell of the container and input the "jupyter-notebook" to start the service.  
3. The docker was running and I could see the web in the address http://localhost:8888.  
(as **part-2/part2.png** shows)
