# 编译镜像
FROM golang:1.11 AS builderImage
WORKDIR /go/src/apiproject
COPY . .
RUN go get -v && go install && mv /go/bin/apiproject /root

# 产物运行镜像
FROM scratch
WORKDIR /root
COPY --from=builderImage /root .
EXPOSE 8089
CMD ["/root/apiproject"]