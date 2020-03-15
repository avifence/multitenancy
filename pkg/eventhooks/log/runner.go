package log

import (
	confiv1 "github.com/configurator/multitenancy/pkg/apis/confi/v1"
	"github.com/configurator/multitenancy/pkg/eventhooks/runner"
	"k8s.io/client-go/kubernetes"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

var log = logf.Log.WithName("log_hook")

type logRunner struct {
}

func NewRunner() runner.HookRunner {
	return &logRunner{}
}

func (s *logRunner) RunCreateHook(clientset *kubernetes.Clientset, mt *confiv1.MultiTenancy, tenant *confiv1.Tenant) error {
	log.Info("New tenant created for multitenancy", "Tenant", tenant.GetName(), "MultiTenancy", mt.GetName(), "Namespace", tenant.GetNamespace())
	return nil
}

func (s *logRunner) RunUpdateHook(clientset *kubernetes.Clientset, mt *confiv1.MultiTenancy, tenant *confiv1.Tenant) error {
	log.Info("Tenant updated for multitenancy", "Tenant", tenant.GetName(), "MultiTenancy", mt.GetName(), "Namespace", tenant.GetNamespace())
	return nil
}

func (s *logRunner) RunDeleteHook(clientset *kubernetes.Clientset, mt *confiv1.MultiTenancy, tenant *confiv1.Tenant, logLines []string) error {
	log.Info("Tenant deleted for multitenancy", "Tenant", tenant.GetName(), "MultiTenancy", mt.GetName(), "Namespace", tenant.GetNamespace())
	return nil
}
