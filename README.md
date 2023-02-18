## 简介
基于docker的NAS系统，特点是简单、免费开源、跨平台

文档地址：http://doc.dockernas.com

可以使用docker部署，运行方式如下述命令所示（将G:\nas或/nas目录替换为自己想保存数据的目录）
```sh
#windows
docker run -d --name dockernas --restart always -p 8080:8080 -v /var/run/docker.sock:/var/run/docker.sock -v G:\nas:/home/dockernas/data xiongzhanzhang/dockernas

#linux 
docker run -d --name dockernas --restart always --add-host=host.docker.internal:host-gateway -p 8080:8080 -v /var/run/docker.sock:/var/run/docker.sock -v /nas:/home/dockernas/data xiongzhanzhang/dockernas
```

目前主要在windows上测试，Linux下问题可能相对多些，可以提issue反馈

## 编译方法
代码编译方式如下所示
```sh
cd frontend 
npm install
npm run build
cd ..
CGO_ENABLED=0 go build ./dockernas.go
```
docker镜像构建方式如下
```sh
#注意需要先在本地构建好前端代码
docker build . -t dockernas
#多平台构建，构建后直接push到dockerhub
docker buildx build --platform linux/arm,linux/arm64,linux/amd64 -t xiongzhanzhang/dockernas:latest . --push
```

## 版本记录
### [0.3.0] 发布日期: 2023-02-18
- 支持多种网络模式
- 前端样式优化
- 应用模板支持配置运行容器的user
### [0.2.0] 发布日期: 2023-02-11
- 移动端适配
- pwa配置
- 部分bug修复

