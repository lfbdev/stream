package argocd

import (
	"fmt"

	"github.com/mitchellh/mapstructure"

	. "github.com/merico-dev/stream/internal/pkg/plugin/common/helm"
	"github.com/merico-dev/stream/pkg/util/helm"
	"github.com/merico-dev/stream/pkg/util/k8s"
	"github.com/merico-dev/stream/pkg/util/log"
)

func Delete(options map[string]interface{}) (bool, error) {
	var opts Options
	if err := mapstructure.Decode(options, &opts); err != nil {
		return false, err
	}

	if errs := validate(&opts); len(errs) != 0 {
		for _, e := range errs {
			log.Errorf("Options error: %s.", e)
		}
		return false, fmt.Errorf("opts are illegal")
	}

	h, err := helm.NewHelm(opts.GetHelmParam())
	if err != nil {
		return false, err
	}

	log.Info("Uninstalling argocd helm chart.")
	if err = h.UninstallHelmChartRelease(); err != nil {
		return false, err
	}

	if err := dealWithNsWhenUninstall(&opts); err != nil {
		return false, err
	}

	return true, nil
}

func dealWithNsWhenUninstall(opts *Options) error {
	if !opts.CreateNamespace {
		return nil
	}

	kubeClient, err := k8s.NewClient()
	if err != nil {
		return err
	}

	return kubeClient.DeleteNamespace(opts.Chart.Namespace)
}
