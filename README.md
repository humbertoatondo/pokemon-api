
# Pokemon API 

<img src="assets/pokemon-header.jpg" alt="Pokemon api header" width="100%"/>

Pokemon API is a Rest API built using Golang which uses the [PokeAPI](https://pokeapi.co/) to make comparisons between two or more pokemons, providing us with useful insight.

## Table of Contents

## Installation

### Requirements
- This project uses Golangs so if you don't have it installed on your machine please head to the [Golang Download Page](https://golang.org/doc/install) and follow the installation steps.
- (Optional) If you want to be able to build and run a docker image, please head to the [Docker Download Page](https://docs.docker.com/get-docker/) and follow the installation steps.
- (Optional) If you want to to run a K8s cluster using this image, please head to the [Kubernetes Download Page](https://kubernetes.io/docs/tasks/tools/) and follow the installation steps. If you are running a cluster locally you might want to use [Minikube](https://minikube.sigs.k8s.io/docs/start/). 
</br>


After you have succesfully installed Golang on your machine just clone the repository using the following command.
```bash
git clone https://github.com/humbertoatondo/pokemon-api.git
```

## Usage
Great! Now you have a copy of this repository on your machine. Next we will see how we can run the project in different environments.

#### Local
To run the project locally just go into the directory and type the following command.
```bash
go run .
```
Now you should get a welcome message if you go into your browser and go the this url: ```http://localhost:5000```

#### Docker
To spin a docker container using this repository we have to first build an image and give it a name using the following command.
```bash
docker build -t <name you want to give to the image> .
```

After we have build the image lets go and run the following command to run our new container.
```bash
docker run -p 5000:5000 -d <name of your image> 
```
Now visit ```http://localhost:5000``` and you should see a welcome message.


