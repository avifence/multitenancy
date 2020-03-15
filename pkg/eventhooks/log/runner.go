package log

import (
	confiv1 "github.com/configurator/multitenancy/pkg/apis/confi/v1"
	"github.com/configurator/multitenancy/pkg/eventhooks/runner"
	"k8s.io/client-go/kubernetes"
)

type logRunner struct {
}

func NewRunner() runner.HookRunner {
	return &logRunner{}
}

func (s *logRunner) RunCreateHook(clientset *kubernetes.Clientset, mt *confiv1.MultiTenancy, tenant *confiv1.Tenant) error {
	return nil
}

func (s *logRunner) RunUpdateHook(clientset *kubernetes.Clientset, mt *confiv1.MultiTenancy, tenant *confiv1.Tenant) error {
	return nil
}

func (s *logRunner) RunDeleteHook(clientset *kubernetes.Clientset, mt *confiv1.MultiTenancy, tenant *confiv1.Tenant, logLines []string) error {
	return nil
}
