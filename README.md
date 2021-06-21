# Pokemon API

<img src="assets/pokemon-header.jpg" alt="Pokemon api header" width="100%"/>

Pokemon API is a Rest API built using Golang which uses the [PokeAPI](https://pokeapi.co/) to make comparisons between two or more pokemons, providing us with useful insight.
If you want to check the hosted version of this project [click here](http://143.244.208.45:5000/).

## Table of Contents

- [Installation](#installation)
  - [Requirements](#requirements)
- [Usage](#usage)
  - [Local](#local)
  - [Docker](#docker)
  - [Kubernetes](#kubernetes)
- [Testing](#testing)
- [Deployment](#deployment)
- [Endpoints](#endpoints)

## Installation <a name="installation"/>

### Requirements <a name="requirements"/>

- This project uses Golangs so if you don't have it installed on your machine please head to the [Golang Download Page](https://golang.org/doc/install) and follow the installation steps.
- (Optional) If you want to be able to build and run a docker image, please head to the [Docker Download Page](https://docs.docker.com/get-docker/) and follow the installation steps.
- (Optional) If you want to to run a K8s cluster using this image, please head to the [Kubernetes Download Page](https://kubernetes.io/docs/tasks/tools/) and follow the installation steps. If you are running a cluster locally you might want to use [Minikube](https://minikube.sigs.k8s.io/docs/start/).
  </br>

After you have succesfully installed Golang on your machine just clone the repository using the following command.

```bash
git clone https://github.com/humbertoatondo/pokemon-api.git
```

## Usage <a name="usage"/>

Great! Now you have a copy of this repository on your machine. Next we will see how we can run the project in different environments.

#### Local <a name="local"/>

To run the project locally just go into the directory and type the following command.

```bash
go run .
```

Now you should get a welcome message if you go into your browser and go the this url: `http://localhost:5000`

#### Docker <a name="docker"/>

To spin a docker container using this repository we have to first build an image and give it a name using the following command.

```bash
docker build -t <name you want to give to the image> .
```

After we have build the image lets go and run the following command to run our new container.

```bash
docker run -p 5000:5000 -d <name of your image>
```

Now visit `http://localhost:5000` and you should see a welcome message.

#### Kubernetes <a name="kubernetes"/>

We will be running a Kubernetes cluster with the help of Minikube, so if you haven't installed it yet, please refer to the installation section under requirements.

First we need to start minikube which will work as a one node cluster. Start minikube by running:

```bash
minikube start
```

To confirm that minikube is running run

```bash
kubectl get nodes
```

and you should see that there is already a node called minikube with Ready status.

In this case you don't need to build an image because we will be using a public image of this repository hosted in [Docker Hub](https://hub.docker.com/) so you only have to apply our current kubernetes configuration located on k8s/pokemon-api.yaml using the following command.

```bash
kubectl apply -f k8s/pokemon-api.yaml
```

Once you run this command you will get a message telling you that a deployment and a service have been created.

Now if you run

```bash
minikube service pokemon-api-service
```

an url will be open in your browser directing you the running rest api.

## Testing <a name="testing"/>

There are two different types of tests in this project, unit and integration testing. Because of this we need to test the program using tags.
To run the unit test just run the following command.

```bash
go test -v ./... --tags=unit
```

Running the integration test is a little different, first we have to run the project as we will normally run it locally.

```bash
go run .
```

And now that an instance of the project is running we can now run the integration tests with the following command.

```bash
go test -v ./... --tags=integration
```

## Deployment <a name="deployment"/>

Everytime we make a commit in the main branch the following pipeline will run:

1. The unit tests will run to check that the tested functions are working as expected.
2. Then, a test image of the pokemon api will be built and deployed to Docker Hub, this will help us when running the integration test.
3. After the test image of the pokemon api is on Docker Hub it will proceed to build a container with that image and run the integration tests. Note that integration tests are being run locally and all functionality that requires third partie apis are mocke.
4. Once all the test ran successfully it will deploy a development image into Docker Hub and another one into DigitalOcean and it will proceed to deploy the clusted in the DigitalOcean cloud.

## Endpoints <a name="endpoints"/>


```bash
/comparePokemons?pokemon1=<pokemonName>&pokemon2=<pokemonName>
```

| Argument | Description                                        |
| -------- | -------------------------------------------------- |
| pokemon1 | Name of the first pokemon to compare in lowercase  |
| pokemon2 | Name of the second pokemon to compare in lowercase |

This endpoint receives 2 pokemon names and returns an object containing comparison_results which stores:

- deals_double_damage: [BOOL] Returns true if pokemon1 can deal double damage to pokemon2.
- receives_half_damage: [BOOL] Returns true if pokemon1 can receive half damage from pokemon2.
- receives_no_damage: [BOOL] Retusn true if pokemon1 can receive no damage from pokemon2

</br>
</br>

```bash
/comparePokemonsMoves?pokemon=<pokemonName>&pokemon=<pokemonName>
```

| Argument | Description                                                                                          |
| -------- | ---------------------------------------------------------------------------------------------------- |
| pokemon | Name of the first pokemon to compare in lowercase. It can receive unlimited number of pokemon names. |
| limit    | Sets the limit of common moves the endpoint should return.                                           |
| lang     | Specifies the language in which you want the pokemon moves to be return in.                          |

This endpoint receive N number of pokemon names and returns an object with an array containing the common moves between all the pokemons.
You can limit the quantity of the moves you want to get with the limit argument and specify the language of the moves with the lang argument.

List of available languages:

- ja-Hrkt
- ko
- zh-Hant
- fr
- de
- es
- it
- en
- ja
- zh-Hans
