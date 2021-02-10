<h1># rh_project</h1>
<h2>Learing GO, docker , kubernetes

Start the project :

1)By clonning and running commond inside the rh_project Directory: 
</h2>
<h3>```go run startServer/main.go```</h3> 
<h3>Open LocalHost 5000 or 50001</h3>

<h2>2)By build container through Docker Image.</h2>
<h3>```docker pull partharora1010/my-golang-app:v1.0.0```</h3>
<h3>i)Running in non detached mode</h3>
<h3>```sudo docker run -it --rm --name my-runing-app partharora1010/my-golang-app:v1.0.0```</h3>
<h3>Open LocalHost 5000 or 50001</h3>
<h3>ii)Running in the detached mode</h3>
<h3>```docker run -p 5000:5000 -tid  partharora1010/my-golang-app:v1.0.0```</h3>
<h3>Open LocalHost 5000</h3>
