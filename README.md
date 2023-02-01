基于docker的NAS系统，特点是简单、免费开源、跨平台

文档地址：http://doc.dockernas.com

可以使用docker部署，运行方式如下述命令所示（将G:\nas或/nas目录替换为自己想保存数据的目录）
```
//windows
docker run -d --name dockernas --restart always -p 8080:8080 -v /var/run/docker.sock:/var/run/docker.sock -v G:\nas:/home/dockernas/data xiongzhanzhang/dockernas

//linux 
docker run -d --name dockernas --restart always --add-host=host.docker.internal:host-gateway -p 8080:8080 -v /var/run/docker.sock:/var/run/docker.sock -v /nas:/home/dockernas/data xiongzhanzhang/dockernas
```

代码构建方式如下所示
```
cd frontend 

npm run build

cd ..

CGO_ENABLED=0 go build ./dockernas.go

docker build . -t dockernas:0.1
```