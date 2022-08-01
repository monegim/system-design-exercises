package controller

import (
	"github.com/monegim/Reloader/internal/pkg/metrics"
	"github.com/monegim/Reloader/internal/pkg/util"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

type Controller struct {
	client  kubernetes.Interface
	indexer cache.Indexer
	queue workqueue.RateLimitingInterface
	informer cache.Controller
	namespace string
	ignoredNamespaces util.List
	collectors metrics.Collectors
}
