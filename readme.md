## 摘要

- **项目名称:** DataCompare

- **简介:** 数据比对 工作可视化

- 功能要点: golang, gin, database/sql, cron任务调度, vue

- **数据库:** mysql, postgres, vertica, oracle

- 容器: Docker

- 主题界面:

  <img src="https://gitee.com/big_ox/my_pictrue/raw/master/Typora/image-20220316162341809.png" alt="image-20220316162341809" style="zoom:50%;" />

## 功能特点

1. 支持**mysql, postgres, vertica, oracle** 数据库创建数据库连接
2. 支持配置源表和目标表信息
3. 可自行对数据比对任务进行调度
4. docker部署, 快速解决环境依赖问题
   - nginx 镜像
   - mysql5.7镜像
   - centos7.6镜像

## 环境依赖

1. docker 环境
2. oracle-cli 环境(仅在配置oracle连接时使用)

## 开发计划

- 测试任务创建时检查 config 和result 的表是否存在
- 任务详情 写入redis
- ,自动补数 ...
- 结果页展示细化


# backend 使用说明

1. 设置golang开发环境
    ```shell
   GO111MODULE="on"
   GOPROXY="https://goproxy.cn,direct"
    ```
2. 设置开发环境标识变量
    ```shell
    export DEBUG_ENV="true" # ture为开发环境,其他值为生产环境
    ```
3. 拉取代码到本地
   ```shell
   git clone https://github.com/daniuEvan/DataCompare.git
   ```
4. 进入项目目录, 修改配置文件参数, 启动项目
   ```shell
   go run main.go
   ```

# fontend

## CD

```
cd frontend

```

## Project setup

```
npm install
```

### Compiles and hot-reloads for development

```
npm run serve
```

### Compiles and minifies for production

```
npm run build
```

### Lints and fixes files

```
npm run lint
```

### Customize configuration

See [Configuration Reference](https://cli.vuejs.org/config/).
