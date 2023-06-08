docker-build:
	docker build -t my-go-app .

docker-start:
	docker start movie-api

docker-stop:
	docker stop movie-api

docker-run:
	docker run -p 8000:5555 -it --name movie-api  my-go-app	
