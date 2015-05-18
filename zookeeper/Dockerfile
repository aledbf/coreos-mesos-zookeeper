FROM alpine:3.1

ENV JAVA_VERSION_MAJOR 8
ENV JAVA_VERSION_MINOR 45
ENV JAVA_VERSION_BUILD 14
ENV JAVA_PACKAGE server-jre

ENV JAVA_HOME /jre
ENV PATH ${PATH}:${JAVA_HOME}/bin

EXPOSE 2181 2888 3888

COPY build.sh /tmp/build.sh

RUN DOCKER_BUILD=true /tmp/build.sh

WORKDIR /opt/zookeeper

ENTRYPOINT ["/opt/zookeeper/bin/zkServer.sh"]

CMD ["start-foreground"]

