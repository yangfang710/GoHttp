FROM golang:1.15.5

MAINTAINER fangfang.yang <413621484@qq.com>

LABEL name="Go-http" \
      description="Project HTTP Structure Meow" \
      owner="413621484@qq.com"

RUN ln -s -f /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

WORKDIR /bin
COPY entrypoint.sh /entrypoint.sh
COPY bin/server /bin/server
COPY conf/app.test.json conf/app.test.json

RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
