# gf docker -t gf-demos
FROM loads/alpine:3.8
LABEL authors="yudhistiro"
LABEL maintainer="mahasiswateknologi@gmail.com"

###############################################################################
#                                INSTALLATION
###############################################################################

# 设置固定的项目路径
ENV WORKDIR /var/www/gf-demos

# 添加应用可执行文件，并设置执行权限
ADD ./bin/linux_amd64/main   $WORKDIR/main
RUN chmod +x $WORKDIR/main

# 添加I18N多语言文件、静态文件、配置文件、模板文件
ADD public   $WORKDIR/public
ADD config   $WORKDIR/config
ADD public/template $WORKDIR/template

###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR
CMD ./main

#ENTRYPOINT ["top", "-b"]