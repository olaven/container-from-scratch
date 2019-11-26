docker build -t container_from_scratch ./environment
docker run -it --privileged --mount src=$(pwd),target=/shared,type=bind container_from_scratch