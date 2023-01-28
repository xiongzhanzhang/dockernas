```
cd frontend 

npm run build

cd ..

CGO_ENABLED=0 go build ./dockernas.go

docker build . -t dockernas:0.1

//windows
docker run -d --name dockernas --restart always -p 8080:8080 -v /var/run/docker.sock:/var/run/docker.sock -v G:\\nas:/home/dockernas/data dockernas:0.1

//linux 
docker run -d --name dockernas --restart always --add-host=host.docker.internal:host-gateway -p 8080:8080 -v /var/run/docker.sock:/var/run/docker.sock -v G:\\nas:/home/dockernas/data dockernas:0.1
```