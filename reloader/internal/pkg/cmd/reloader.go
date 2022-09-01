package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/monegim/Reloader/internal/pkg/metrics"
	"github.com/monegim/Reloader/internal/pkg/options"
	"github.com/monegim/Reloader/internal/pkg/util"
	"github.com/monegim/Reloader/pkg/kube"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func NewReloaderCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "reloader",
		Short:   "A watcher for your Kubernetes cluster",
		PreRunE: validateFlags,
		Run:     startReloader,
	}
	return cmd
}

func validateFlags(*cobra.Command, []string) error {
	err := fmt.Sprintf(("must be one of: "))
	return errors.New(err)
}

func configureLogging(logFormat string) error {
	switch logFormat {
	case "json":
		logrus.SetFormatter(&logrus.JSONFormatter{})
	default:
		if logFormat != "" {
			return fmt.Errorf("unsupported logging formatter: %q", logFormat)
		}
	}
	return nil
}

func startReloader(cmd *cobra.Command, args []string) {
	err := configureLogging(options.LogFormat)
	if err != nil {
		logrus.Warn(err)
	}
	logrus.Info("Strting Reloader")
	currentNamespace := os.Getenv("KUBERNETES_NAMESPACE")
	if len(currentNamespace) == 0 {
		currentNamespace = v1.NamespaceAll
		logrus.Warnf("KUBERNETES_NAMESPACE is unset, will detect changes in all namespaces.")
	}
	clientset, err := kube.GetKubernetesClient()
	if err != nil {
		logrus.Fatal(err)
	}
	ignoredResourcesList, err := getIgnoredResourcesList(cmd)
	if err != nil {
		logrus.Fatal(err)
	}
	ignoredNamespacesList, err := getIgnoredNamespacesList(cmd)
	if err != nil {
		logrus.Fatal(err)
	}

	collectors := metrics.SetupPrometheusEndpoint()

	for k := range kube.ResourceMap{
		if ignoredResourcesList.Contain(k) {
			continue
		}

		c, err := contro
	}
}

func getIgnoredNamespacesList(cmd *cobra.Command) (util.List, error) {
	return getStringSliceFromFlags(cmd, "namespaces-to-ignore")
}

func getStringSliceFromFlags(cmd *cobra.Command, flag string) ([]string, error) {
	slice, err := cmd.Flags().GetStringSlice(flag)
	if err != nil {
		return nil, err
	}

	return slice, nil
}

func getIgnoredResourcesList(cmd *cobra.Command) (util.List, error) {

	ignoredResourcesList, err := getStringSliceFromFlags(cmd, "resources-to-ignore")
	if err != nil {
		return nil, err
	}

	for _, v := range ignoredResourcesList {
		if v != "configMaps" && v != "secrets" {
			return nil, fmt.Errorf("'resources-to-ignore' only accepts 'configMaps' or 'secrets', not '%s'", v)
		}
	}

	if len(ignoredResourcesList) > 1 {
		return nil, errors.New("'resources-to-ignore' only accepts 'configMaps' or 'secrets', not both")
	}

	return ignoredResourcesList, nil
}
