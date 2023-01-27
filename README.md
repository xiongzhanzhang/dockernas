```
cd frontend 

npm run build

cd ..

CGO_ENABLED=0 go build ./dockernas.go

docker build . -t dockernas:0.1

docker run -d -p 8080:8080 -v /var/run/docker.sock:/var/run/docker.sock -v G:\\nas:/home/dockernas/data dockernas:0.1
```