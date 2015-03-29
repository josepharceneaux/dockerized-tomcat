default: docker

docker: Dockerfile
	docker build -t tomcat-server:latest .

clean:
	rm server
