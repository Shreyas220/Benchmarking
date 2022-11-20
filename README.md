# Benchmarking
automating benchmarking Kubearmor with redis and nginx 


## Steps to follow Before running Redis benchmark 
```
helm install release bitnami/redis
```


```
export REDIS_PASSWORD=$(kubectl get secret --namespace default release-redis -o jsonpath="{.data.redis-password}" | base64 -d)
```

```
kubectl run --namespace default redis-client --restart='Never'  --env REDIS_PASSWORD=$REDIS_PASSWORD  --image docker.io/bitnami/redis:7.0.5-debian-11-r15 --command -- sleep infinity
```

```
kubectl exec --tty -i redis-client \
   --namespace default -- bash
```

```
redis-cli config set stop-writes-on-bgsave-error no
```

```
redis-server
```

## Steps to follow Before running nginx Benchamrk 

 just have apachebench AKS setup
 
 ```
 apt-get install apache2-utils

 ```



## How to run benchmark  
```
go run main.go -b redis -n 1000000 -c 100 -i 3
```
- -n -> is the number of request to send (default 10000)
- -c -> is the number of client (default 100)
- -i -> is the number of iteration (default 2) for the first iteration n will be 10000 and for the second iteration n will be mulitplied by 10 and will be 100000 and so on 
- -b is the name of benchmark (default nginx) two options -> "redis" and "nginx"
