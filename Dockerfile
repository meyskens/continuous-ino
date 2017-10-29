ARG arch
FROM multiarch/debian-debootstrap:${arch}-stretch
ENV arduinoversion=1.8.5

RUN apt-get update &&\
    apt-get install -y \
    build-essential \
    wget \
    git 

RUN apt-get update && apt-get install -y xz-utils


# Install Arduino
RUN case "${ARCH}" in                                                                                 \
    amd64|x86_64)                                                                                     \
      wget -O arduino.tar.xz https://downloads.arduino.cc/arduino-${arduinoversion}-linux64.tar.xz    \
      ;;                                                                                              \
    arm64|aarch64|armv7l|armhf|arm)                                                                   \
      wget -O arduino.tar.xz https://downloads.arduino.cc/arduino-${arduinoversion}-linuxarm.tar.xz   \ 
      ;;                                                                                              \
    *)                                                                                                \
      echo "Unhandled architecture: ${ARCH}."; exit 1;                                                \
      ;;                                                                                              \
esac && tar -xJf arduino.tar.xz && rm -f arduino.tar.xz

RUN mv arduino-${arduinoversion} /usr/local/share/arduino/ && /usr/local/share/arduino/install.sh && ln -s /usr/local/share/arduino/arduino /usr/local/bin/arduino 

COPY ./continuous-ino /usr/local/bin/continuous-ino

CMD /usr/local/bin/continuous-ino