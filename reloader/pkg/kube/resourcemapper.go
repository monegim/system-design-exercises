package kube

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

var ResourceMap = map[string]runtime.Object{
	"configMaps": &v1.ConfigMap{},
	"secrets":    &v1.Secret{},
}
