ARG arch
FROM multiarch/debian-debootstrap:${arch}-stretch

RUN apt-get update &&\
    apt-get install -y \
    python \
    python-dev \
    python-pip \
    python-setuptools \
    python-wheel \
    picocom \
    arduino \
    --no-install-recommends

RUN pip install ino

COPY ./continuous-ino /usr/local/bin/continuous-ino

CMD /usr/local/bin/continuous-ino