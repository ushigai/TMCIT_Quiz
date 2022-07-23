#!/bin/bash
git checkout shoma
git pull
docker build -t shomaigu/zissyu:latest /home/pi/go_app/hon/TMCIT_Quiz/build/
docker image push shomaigu/zissyu:latest
sleep 10

sudo -u pi kubectl delete -f /home/pi/go_app/hon/TMCIT_Quiz/k8s/deploy.yml -f /home/pi/go_app/hon/TMCIT_Quiz/k8s/nodeport.yml
sudo -u pi kubectl apply -f /home/pi/go_app/hon/TMCIT_Quiz/k8s/deploy.yml -f /home/pi/go_app/hon/TMCIT_Quiz/k8s/nodeport.yml
sudo -u pi kubectl get pod
sudo -u pi kubectl get svc