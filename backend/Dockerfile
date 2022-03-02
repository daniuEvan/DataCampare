FROM centos

RUN ["mkdir","-p","/opt/sweetcake"]
WORKDIR "/opt/sweetcake"
CMD ["pwd"]
COPY ./main /opt/sweetcake
CMD ["/bin/bash ./main"]
CMD ["yum install vim -y"]
CMD ["ping", "127.0.0.1"]

