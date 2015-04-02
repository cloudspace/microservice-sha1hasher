# Microservice-sha1hasher 
Encrypts data by the SHA-1 algorithm written in Go

#### Usage
Input must have 1 arguments:
-data (string)


#### Example Input
Input:
```
go run sha1hash.go "I can haz"
```
Output (Success):
```
{"result":"b64e21f3cf94ce384d579b3d128175c2"}
```

#### How to build docker image
Requirements:

1. Golang environment set up
2. Git
3. boot2docker running

```
go get https://github.com/cloudspace/microservice-sha1hasher
docker run --rm -v $(pwd):/src centurylink/golang-builder
docker build -t <username>/microservice-sha1hasher:0.1 ./

```
In order for `docker run --rm -v $(pwd):/src centurylink/golang-builder` to work you need to have the github url on the top line of main.go. It should look like this:
```
package main // import "github.com/cloudspace/microservice-md5hasher"
```
You also must push your code to github before building the docker image.
