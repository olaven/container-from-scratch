from ubuntu:latest 

# RUN add-apt-repository -y ppa:gophers/archive
RUN apt-get -y update
RUN apt-get -y install golang

# Download and extract rootfs 
RUN apt-get -y install wget 
RUN wget http://cdimage.ubuntu.com/ubuntu-base/releases/18.04/release/ubuntu-base-18.04-base-amd64.tar.gz
RUN mkdir -p /home/rootfs 
RUN tar -xzf ubuntu-base-18.04-base-amd64.tar.gz -C /home/rootfs