package k8s

import (
	"os"
	"strings"
)

var (
	_namespace       = ""
	_namespaceLoaded = false
)

// GetCurrentNamespace returns the namespace of the pod in which the application is running. If the namespace is not found, it returns "default".
func GetCurrentNamespace() (namespace string, loaded bool) {
	if !_namespaceLoaded {
		ns := os.Getenv("NAMESPACE")
		if len(ns) == 0 {
			ns = "default"
		} else {
			ns = strings.TrimSpace(ns)
		}
		_namespace = ns
		_namespaceLoaded = true
	}
	return _namespace, _namespaceLoaded
}
