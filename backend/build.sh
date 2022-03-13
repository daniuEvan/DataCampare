# 打包 linux amd64 程序
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./data-compare main.go

docker build -t dc-server:v1.0 .
docker run -d -p 9090:9090 --name=dc-server dc-server:v1.0




