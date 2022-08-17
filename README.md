## Two Port App

**Background**: _Work around for OpenShift Routes exposing just one port._

Deploy the app in OpenShift

```bash
kubectl apply -f k8s.yaml
```

After deployed, you will see two routes that ultimately route back to the same `two-port-app` service

```bash
kubectl get routes
```

output

```bash
NAME       HOST/PORT                                                                 PATH   SERVICES       PORT     TERMINATION   WILDCARD
port-one   port-one-default.apps.tetrate-mp-cp.kni.syseng.devcluster.openshift.com          two-port-app   port-1                 None
port-two   port-two-default.apps.tetrate-mp-cp.kni.syseng.devcluster.openshift.com          two-port-app   port-2                 None
```

Test route one

```bash
curl $(kubectl get route port-one --template='{{ .spec.host }}')/healthz
```

output

```bash
OK - Port One
```

Test route two

```bash
curl $(kubectl get route port-two --template='{{ .spec.host }}')/healthz
```

output

```bash
OK - Port Two
```
