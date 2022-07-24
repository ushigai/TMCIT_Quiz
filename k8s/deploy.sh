#!/bin/bash
git checkout shoma
git pull
docker build -t shomaigu/ubuntu:latest /home/pi/go_app/hon/TMCIT_Quiz/k8s/
docker image push shomaigu/ubuntu:latest
sleep 10

sleep 10
sudo -u pi kubectl delete -f /home/pi/go_app/hon/TMCIT_Quiz/k8s/ubuntu.yml
sleep 10
sudo -u pi kubectl apply -f /home/pi/go_app/hon/TMCIT_Quiz/k8s/ubuntu.yml
sleep 10
sudo -u pi kubectl get pod -n kobayakawake
sudo -u pi kubectl get svc -n kobayakawake

