#!/usr/bin/env sh

# fail on any command exiting non-zero
set -eo pipefail

if [[ -z $DOCKER_BUILD ]]; then
  echo
  echo "Note: this script is intended for use by the Dockerfile and not as a way to build zoopeeper locally"
  echo
  exit 1
fi

apk add --update curl ca-certificates bash

cd /tmp

curl -sSL -o glibc-2.21-r2.apk "https://circle-artifacts.com/gh/andyshinn/alpine-pkg-glibc/6/artifacts/0/home/ubuntu/alpine-pkg-glibc/packages/x86_64/glibc-2.21-r2.apk"

apk add --allow-untrusted glibc-2.21-r2.apk

curl -sSL -o glibc-bin-2.21-r2.apk "https://circle-artifacts.com/gh/andyshinn/alpine-pkg-glibc/6/artifacts/0/home/ubuntu/alpine-pkg-glibc/packages/x86_64/glibc-bin-2.21-r2.apk"

apk add --allow-untrusted glibc-bin-2.21-r2.apk

/usr/glibc/usr/bin/ldconfig /lib /usr/glibc/usr/lib

echo "Downloading Oracle JDK..."
curl -jksSLH "Cookie: oraclelicense=accept-securebackup-cookie"\
  http://download.oracle.com/otn-pub/java/jdk/${JAVA_VERSION_MAJOR}u${JAVA_VERSION_MINOR}-b${JAVA_VERSION_BUILD}/${JAVA_PACKAGE}-${JAVA_VERSION_MAJOR}u${JAVA_VERSION_MINOR}-linux-x64.tar.gz | gunzip -c - | tar -xf -

# install confd
echo "Downloading confd..."
curl -sSL -o /sbin/confd https://s3-us-west-2.amazonaws.com/opdemand/confd-git-73f7489 \
  && chmod +x /sbin/confd

apk del curl ca-certificates

mkdir -p /tmp/zookeeper /opt

echo "Downloading zookeeper..."
curl -sSL http://apache.mirrors.pair.com/zookeeper/zookeeper-3.5.0-alpha/zookeeper-3.5.0-alpha.tar.gz | tar -xzf - -C /opt

ln -s /opt/zookeeper-3.5.0-alpha /opt/zookeeper

cp /app/zoo.cfg /opt/zookeeper-3.5.0-alpha/conf/zoo.cfg

touch /opt/zookeeper/conf/zoo_replication.cfg.dynamic

mv jdk1.${JAVA_VERSION_MAJOR}.0_${JAVA_VERSION_MINOR}/jre /jre

# cleanup

rm /jre/bin/jjs
rm /jre/bin/keytool
rm /jre/bin/orbd
rm /jre/bin/pack200
rm /jre/bin/policytool
rm /jre/bin/rmid
rm /jre/bin/rmiregistry
rm /jre/bin/servertool
rm /jre/bin/tnameserv
rm /jre/bin/unpack200
rm /jre/lib/ext/nashorn.jar
rm /jre/lib/jfr.jar

rm -rf /jre/lib/jfr
rm -rf /jre/lib/oblique-fonts

rm -rf /tmp/* /var/cache/apk/*
