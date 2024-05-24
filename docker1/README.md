# docker

Demo on :
1) How to containerize your go code.
2) How to create DockerFile

module name : `docker1` (go mod init docker1)

NOTE : install docker extension on vscode (IDE)

# Docker file

Capital `D`ockerfile

synatx : all caps

FROM
WORKDIR
COPY
RUN
EXPOSE
ENTRYPOINT

1) `FROM golang:version-alpine`  
- choose from hub.docker.com >> search golang >> click first golang image >> choose one from list 
- `docker pull golang` --> `golang` is image name 
- `version` specified in go.mod (recommanded `alpine`)

2) `WORKDIR /name` 
- a working directory is created within the Docker image where the application's source code will be copied

3) `COPY . /name`
- The Dockerfile copies the golang code from current dir `.` into the Docker image's working directory `/name`.
Ensuring the application code is available inside the container.

4) `RUN go build -o exect /name` 
- builds `/name` current folder
- `-o exect` name of executable
- if any dependencies use ` RUN go mod download` which will download dependencies mentioned in `go.mod` file

5) `EXPOSE 7011` 
- localhost:7011 , same as mentioned in `main.go`
- in case of multiple ports, list them all

6) `ENTRYPOINT ["./exect"]`
- ENTRYPOINT or CMD ,either can be used
- command to run golang code/app when docker container starts.
- `./exect` from RUN cmd


Docker cmd : (in terminal)

1) `docker build -t name .`

docker build --> building docker image
-t name --> tag; name of docker 
. --> dockerfile dir ; here it is in same dir so `.`

2) `docker run -t name -p 0000:0000`

docker run --> run docker image
-t name --> which image
-p --> app port : docker port (both can be different bt app port from main.go)

`docker run -t name -d -p 0000:0000`
-d --> runs app on background & doesnt take up terminal window

NOTE : sequence of flags are important.

3) `docker ps` 
lists all running images 