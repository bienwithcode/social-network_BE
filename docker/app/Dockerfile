# Telling to use Docker's golang ready image
FROM golang:1.19.2

EXPOSE 8080 80 443

# Set pwd to the go folder
WORKDIR /var/www/src
COPY ./src /var/www/src

RUN apt-get update -y && \
	apt-get install -y \
	vim

# Install and build binaries revel for this server
# RUN go get github.com/revel/revel
# RUN go get github.com/revel/cmd/revel

# install package
RUN go mod tidy
# Copy the entrypoint file
ADD /docker/app/infinite_loop.sh /infinite_loop.sh
RUN chmod +x /infinite_loop.sh

#Start process
ENTRYPOINT ["/infinite_loop.sh"]