FROM golang
COPY . /server
WORKDIR /server
EXPOSE 5005
CMD go run . config.yml