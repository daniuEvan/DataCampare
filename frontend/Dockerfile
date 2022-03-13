FROM nginx
LABEL version="0.1"
LABEL maintainer="big_ox@163.com"

RUN rm -rf /etc/nginx/conf.d/default.conf
RUN rm -rf /etc/nginx/conf.d/nginx.conf

COPY ./dist /apps/DataCompare/frontend/
COPY nginx.conf /etc/nginx/nginx.conf

EXPOSE 80
