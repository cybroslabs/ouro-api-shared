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
	if _namespace == "" {
		if data, err := os.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace"); err == nil {
			if ns := strings.TrimSpace(string(data)); len(ns) > 0 {
				_namespace = ns
				_namespaceLoaded = true
				return ns, true
			}
		}
		if ns, ok := os.LookupEnv("NAMESPACE"); ok && len(ns) > 0 {
			_namespace = ns
			_namespaceLoaded = true
			return ns, true
		}
		_namespace = "default"
	}
	return _namespace, _namespaceLoaded
}
