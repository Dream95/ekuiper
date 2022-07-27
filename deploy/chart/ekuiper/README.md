# eKuiper
A lightweight IoT edge analytic software

![Version: 1.3.0](https://img.shields.io/badge/Version-1.3.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.5.0](https://img.shields.io/badge/AppVersion-1.5.0-informational?style=flat-square)

## Install the Chart

- From Github
```
git clone https://github.com/lf-edge/ekuiper.git
cd deploy/chart/ekuiper
helm install my-ekuiper .
```

- From Helm repo
```
helm repo add emqx https://repos.emqx.io/charts
helm install my-ekuiper emqx/ekuiper
```


## Uninstall Chart
```
helm uninstall my-ekuiper
```

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| affinity | object | `{}` |  |
| clusterDomain | string | `"cluster.local"` | clusterDomain Kubernetes Cluster Domain |
| ekuiperEnv | object | `{"enabled":true,"key":{"mqttDefaultServer":"MQTT_SOURCE__DEFAULT__SERVER"},"value":{"mqttDefaultServer":"tcp://broker.emqx.io:1883"}}` | remove this when the mqtt_source configmap is available |
| image.pullPolicy | string | `"IfNotPresent"` |  |
| image.repository | string | `"lfedge/ekuiper"` |  |
| nodeSelector | object | `{}` |  |
| persistence.accessMode | string | `"ReadWriteOnce"` |  |
| persistence.enabled | bool | `false` |  |
| persistence.existingClaim | string | `""` | Existing PersistentVolumeClaims The value is evaluated as a template So, for example, the name can depend on .Release or .Chart |
| resources | object | `{}` |  |
| service.annotations | object | `{}` | Provide any additional annotations which may be required. Evaluated as a template |
| service.nodePorts | object | `{"ekuiper":null,"restapi":null}` | Specify the nodePort(s) value for the LoadBalancer and NodePort service types. ref: https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport |
| service.ports | object | `{"ekuiper":{"name":"ekuiper","port":20498},"restapi":{"name":"restapi","port":9081}}` | Service ports |
| service.ports.ekuiper.name | string | `"ekuiper"` | eKuiper port name |
| service.ports.ekuiper.port | int | `20498` | eKuiper port |
| service.ports.restapi.name | string | `"restapi"` | eKuiper restapi port name |
| service.ports.restapi.port | int | `9081` | eKuiper restapi port |
| service.type | string | `"ClusterIP"` | service type |
| serviceAccount.annotations | object | `{}` |  |
| serviceAccount.create | bool | `true` |  |
| serviceAccount.name | string | `""` |  |
| tls | object | `{"autoGenerated":false,"caCertificate":"","enabled":false,"existingSecret":"","existingSecretFullChain":false,"serverCertificate":"","serverKey":""}` | Enable encryption to eKuiper |
| tls.autoGenerated | bool | `false` | tls.autoGenerated Generate automatically self-signed TLS certificates |
| tls.caCertificate | string | `""` | tls.caCertificate Certificate Authority (CA) bundle content |
| tls.enabled | bool | `false` | tls.enabled Enable TLS support on eKuiper |
| tolerations | list | `[]` |  |