ARG arch
FROM multiarch/debian-debootstrap:${arch}-stretch

RUN apt-get update &&\
    apt-get install -y python python-dev python-pip picocom arduino

RUN pip install ino
