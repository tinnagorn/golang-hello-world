# Golang Hello World

docker step 1: build

```
docker build -t tintinktb20/hello-docker:dev .
```


and check image status 
```
docker image ls | grep hello-docker
```

docker step 2: run 
```
docker run -d -p 8081:1323 tintinktb20/hello-docker:dev 
```

docker step 3: test this service use curl

```
$ curl --location --request GET 'http://localhost:8081/hello-world'
```

How to see the log
```
$ docker container ls
$ docker exec -it $CONTAINER_ID /app/goapp bash
``

docker step 4: push 

```
docker push tintinktb20/hello-docker:dev 
```


check result in https://hub.docker.com/repository/docker/tintinktb20/hello-docker














step 4: run docker-compose up for create mysql / redis

```
$ docker-compose up
```

step 5: run docker-compose up for create mysql / redis

```
$ docker-compose up
```