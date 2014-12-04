# docker-go-hello-world

A simple example of building and running a Go application within a [Docker](https://www.docker.com/) container.


## Running

This assumes you are running OS X and have [boot2docker](http://boot2docker.io/) installed.


### Clone the repository

	git clone https://github.com/stevenwilkin/docker-go-hello-world.git
	cd docker-go-hello-world

### Start boot2docker

	boot2docker init
	boot2docker start
	$(boot2docker shellinit)

### Build the image

	docker build -t hello-world .

### Run the container

	docker run -p 8080:8080 --rm hello-world


You should now be able to point  `curl` at port 8080 on the VM spun up by `boot2docker`:

	$ curl $(boot2docker ip 2>/dev/null):8080
	Hello Go!