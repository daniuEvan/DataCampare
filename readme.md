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
cd SweetCake/fontend

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
