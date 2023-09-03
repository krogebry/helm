# Overview

Using a starter.  Create a new helm chart with the Mount Fuji GoLang Rest Service starter:

```shell
helm-deploy create -p starters/golang-rest-service my-awesome-service-chart-name
```

# Starters

List of starter packages to help you roll your stuff!

## Baseline

All starter charts have these attributes by default.

* Auto docs using helm-doc with your own template for customization.
* Istio Virtual Service. 
* OPA rules baked in for deployment.
* Resources set; by default pretty low, but will get you started.
* Service which points to your pods.
* ServiceAccount: disabled by default, but there if you need it.
* HorizontalPodAutoscaler enabled by default.
* Basic connection test
* Output message

### Coming soon

* Default vault side car injection.

## GoLang - Rest service starter

This starter includes some nice things that we expect to have in a basic GoLang RESTful service.

* Default port 8080 - `http`.
* Metrics port 2112 - `metrics`.
* ServiceMonitor for metrics port.
* Database values
  * DB_* values wired into the ENV
