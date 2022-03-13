# shellcheck disable=SC2164
cd ./backend/
# 打包 linux amd64 程序
rm -rf ./data-compare
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./data-compare main.go
scp ./data-compare root@aliyun:/apps/DataCompare/backend
scp ./config-pro.yaml root@aliyun:/apps/DataCompare/backend
scp ./Dockerfile root@aliyun:/apps/DataCompare/backend

cd ../frontend
rm -rf ./dist
npm run build
scp -r ./dist root@aliyun:/apps/DataCompare/frontend
scp ./Dockerfile root@aliyun:/apps/DataCompare/frontend
scp ./nginx.conf root@aliyun:/apps/DataCompare/frontend

cd ../
scp -r backendDatabase  root@aliyun:/apps/DataCompare




