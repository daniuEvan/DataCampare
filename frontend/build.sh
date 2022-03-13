rm -rf ./dist
npm run build

docker build -t dc-web:v1.0 .
docker run -d -p 80:80 --name=dc-web dc-web:v1.0