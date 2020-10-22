FROM centos:centos8

ADD . /opt/

EXPOSE 8090

CMD [ "/opt/httpIdle" ]
