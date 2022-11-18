# Benchmarking
automating benchmarking Kubearmor and redis 


## Steps to follow 
```
helm install release bitnam/redis
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

now do 
```
go run main.go
```

