# Static test data generated by jjaramillo on Tue May 30 20:00:33 PDT 2023

### `nri-kubernetes` commit
```
927f63f27101a7ba5840d84fa4f06ae67bcd64c2
```

`git status --short`

```
 M ../../../../charts/internal/e2e-resources/templates/cronjob.yml
?? ../1_19/
?? ../1_20/
?? ../1_21/
?? ../1_22/
?? ./
```

### `$ kubectl version`
```
Client Version: version.Info{Major:"1", Minor:"27", GitVersion:"v1.27.2", GitCommit:"7f6f68fdabc4df88cfea2dcf9a19b2b830f1e647", GitTreeState:"clean", BuildDate:"2023-05-17T14:13:27Z", GoVersion:"go1.20.4", Compiler:"gc", Platform:"darwin/arm64"}
Kustomize Version: v5.0.1
Server Version: version.Info{Major:"1", Minor:"23", GitVersion:"v1.23.17", GitCommit:"953be8927218ec8067e1af2641e540238ffd7576", GitTreeState:"clean", BuildDate:"2023-02-22T13:27:46Z", GoVersion:"go1.19.6", Compiler:"gc", Platform:"linux/arm64"}
```

### Kubernetes nodes
```
NAME           STATUS   ROLES                  AGE     VERSION    INTERNAL-IP    EXTERNAL-IP   OS-IMAGE             KERNEL-VERSION     CONTAINER-RUNTIME
datagen-1-23   Ready    control-plane,master   8m50s   v1.23.17   192.168.49.2   <none>        Ubuntu 20.04.5 LTS   5.15.49-linuxkit   containerd://1.6.20
```