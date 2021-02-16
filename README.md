# rh_project
## Learing GO, docker , kubernetes

### Start the project :
- By clonning and running commond inside the rh_project Directory: 

 ```console 
 go run startServer/main.go
 ```

> Open LocalHost port 5000 or 5001

- By build container through Docker Image.

Run Docker Daemon  
```console 
systemctl start docker 
```
Pull  docker container
```console 
docker pull partharora1010/my-golang-app:v1.0.0
```
 
- Running in non detached mode
```console 
sudo docker run -p 5000:5000 -it --rm --name my-runing-app partharora1010/my-golang-app:v1.0.0
```

>Open LocalHost port 5000 or 5001

- Running in the detached mode

```console 
docker run -p 5000:5000 -tid  partharora1010/my-golang-app:v1.0.0 
```
>Open LocalHost port 5000 or 5001

