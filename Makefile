default: docker

docker: Dockerfile
	docker build -t josepharceneaux/tomcat-server:latest .

test: tests
tests: stelligent_test.go
	go test
