# rh_project
## Learing GO, docker , kubernetes

### Start the project :

1)By clonning and running commond inside the rh_project Directory: 

### `go run startServer/main.go` 
#### Open LocalHost port 5000 or 5001

## 2)By build container through Docker Image.
#### Run Docker Daemon is running  `systemctl start docker`
### `docker pull partharora1010/my-golang-app:v1.0.0`
 
### i)Running in non detached mode
### `sudo docker run -p 5000:5000 -it --rm --name my-runing-app partharora1010/my-golang-app:v1.0.0`
#### Open LocalHost port 5000 or 5001

### ii)Running in the detached mode
### `docker run -p 5000:5000 -tid  partharora1010/my-golang-app:v1.0.0`
#### Open LocalHost port 5000 or 5001
