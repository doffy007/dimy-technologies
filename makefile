create-container:
	sudo docker run --name dimy_tech -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root mysql:latest

