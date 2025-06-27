# testapp-OTELbasic
I use this to demonstrate OTEL instrumentation

### Instructions

1. Ensure the Red Hat Build of OpenTelemetry operator is installed on your cluster

2. Create a namespace, launch deployment, create service and route.

```
$ oc apply -f otel-deploy.yaml
$ oc apply -f expose.yaml
```

3. apply the rolebinding and collector objects

```
$ oc apply -f otel-rolebinding.yaml
$ oc apply -f otel-collector.yaml
```

4. scale the testapp down and back up, inorder to deploy the sidecar

```
$ oc scale --replicas=0 deployment/otel-example-deployment
$ oc scale --replicas=1 deployment/otel-example-deployment
```

5. call the endpoint and view the trace

```
$ curl -I `oc get route | awk 'NR>1 {print $2}'`/ping
$ oc logs deploy/otel-example-deployment -c otc-container
```
