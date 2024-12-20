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
I will simply use SQLite. SQLite is sufficient for up to 100K Hits per day (According to their website). The database will be a container itself.

## The CRUD
The CRUD ops will simply handle the CRUD's for Tasks. Each container will connect to the same database container.

## The CI/CD
There will be GitHub Workflows to check if edge cases for deployment are met.

## The Test
Any function in this Projects has a test. It's developed testdriven.

## Testing my project
If you are mad enough to want to really try this then make sure the following ports are free:
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