default: docker

docker: Dockerfile
	docker build -t josepharceneaux/tomcat-server:latest .
