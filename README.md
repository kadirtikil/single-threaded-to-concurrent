# Concurrent CRUD API using Docker and Kubernetes.
(Yes I know Go has build in concurrency features, yes I know its far easier to use than this.)
This Project is to illustrate my understanding of Software Engineering/Testing and Devlopment Operations, 
and not to develop a blazingly fast concurrent backend system. 

## The Architecture
Each Container gets a special port for unique identification.

## The Cluster
The Cluster contains docker images. Each image contains a CRUD operation (might add more operations for better illustration).
I then use Kubernetes to handle the requests and the load balancing, such that they are taken care of concurrently and therefore faster.

## The Database
MySQL is the DB of choice for this, due to it's ability for concurrent read/write operations.
Just pull the default mysql image from docker for this by using 
```
$ docker pull mysql
```
And then start an instance
```
$ docker run --name crudDB -e MYSQL_ROOT_PASSWORD=admin123 -d mysql:9.1.0
```

To connect to it, you should create a seperate docker network for the whole application with:
```
$ docker network create crudnetwork
```

Now connect to the container with:

```
$ docker run -it --network crudnetwork --rm mysql mysql -h crudDB -u admin -p admin123
```
(Note that MySQL could ask for the password seperately by creating a new line. If it doesn't work dont add anything after **-p** and see if the CLI asks for a password.)



If you want to use the CLI inside the container to check the DB for yourself then run:
```
$ docker exec -it crudDB bash
```

Logs are visible by running:
```
$ docker logs some-mysql

```

Normaly I wanted to use a special directory for the data storage so the other containers could find it easier. But for that there would have to be extra steps 
like making sure that the directory exists and so on. Therefore I'm sticking with the default and let the container manage the data. Since the application in itself
is not that complex, this is sufficient.

For additional options you can check out the [MySQL Docker Documentation](https://hub.docker.com/_/mysql).


## The CRUD
The CRUD ops will simply handle the CRUD's for Tasks. Each container will connect to the same database container.

## The CI/CD
There will be GitHub Workflows to check if edge cases for deployment are met.

## The Test
Any function in this Projects has a test. It's developed testdriven.

## Testing my project
If you want to really try this then make sure the following ports are free:
- 8001
- 8002
- 8003
- 8004
- 8005

These ports handle the 5 containers. 
You can also just pull the database container and on other container for an operation like create and read to check if it really works. 
Other than that the project should show my skill level quite well.

## TL;DR;
- Porfolio Project
- Showing DevOps & Software Engineering/Testing skills.
- Not for commercial use (at least not intended).
- Techstack:
    - Go
    - Docker
    - Kubernetes

(I realized that this architecture could be used for single threaded systems like NodeJS to make them concurrent.)

Thanks for reading!