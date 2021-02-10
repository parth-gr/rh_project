# rh_project
## Learing GO, docker , kubernetes

### Start the project :

1)By clonning and running commond inside the rh_project Directory: 

### `go run startServer/main.go` 
#### Open LocalHost 5000 or 50001

## 2)By build container through Docker Image.
### `docker pull partharora1010/my-golang-app:v1.0.0` 
### i)Running in non detached mode
### `sudo docker run -it --rm --name my-runing-app partharora1010/my-golang-app:v1.0.0`
#### Open LocalHost 5000 or 50001
### ii)Running in the detached mode
### `docker run -p 5000:5000 -tid  partharora1010/my-golang-app:v1.0.0`
#### Open LocalHost 5000
