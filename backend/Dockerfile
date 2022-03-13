FROM centos:centos7.9.2009
LABEL version="0.1"
LABEL maintainer="big_ox@163.com"

EXPOSE 9090

RUN mkdir -p /apps/DataCompare/backend
COPY ./data-compare /apps/DataCompare/backend/
RUN chmod +x /apps/DataCompare/backend/data-compare
COPY ./config-pro.yaml /apps/DataCompare/backend/
RUN export GIN_MODE=release

WORKDIR /apps/DataCompare/backend/
CMD ["nohup" ,"./data-compare" , "&"]

