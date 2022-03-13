docker build -t mysql:5.7 .

mkdir -p /data/mysql/data
mkdir -p /data/mysql/conf
docker run --name mysql:5.7 -p 3306:3306 -e MYSQL_ROOT_PASSWORD=Abcd@123456 -d -v /data/mysql/conf:/etc/mysql/conf.d -v /data/mysql/data/:/var/lib/mysql mysql:5.7