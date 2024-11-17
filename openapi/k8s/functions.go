package k8s

import (
	"os"
	"strings"
)

var (
	namespace       = ""
	namespaceLoaded = false
)

// GetCurrentNamespace returns the namespace of the pod in which the application is running. If the namespace is not found, it returns "default".
func GetCurrentNamespace() string {
	if !namespaceLoaded {
		ns := strings.TrimSpace(os.Getenv("NAMESPACE"))
		if len(ns) == 0 {
			ns = "default"
		}
		namespace = ns
		namespaceLoaded = true
	}
	return namespace
}
