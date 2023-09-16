#! /bin/sh
sudo docker stop broker-service
sudo docker run -d -it --name broker-service -p 8080:8080 broker-service 