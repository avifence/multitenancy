package slack

import (
	"fmt"
	"strings"

	slackhook "github.com/ashwanthkumar/slack-go-webhook"
	confiv1 "github.com/configurator/multitenancy/pkg/apis/confi/v1"
	"github.com/configurator/multitenancy/pkg/eventhooks/runner"
	"k8s.io/client-go/kubernetes"
)

type slackRunner struct {
	config *confiv1.SlackConfig
}

func NewRunner(conf *confiv1.SlackConfig) runner.HookRunner {
	return &slackRunner{config: conf}
}

func (s *slackRunner) RunCreateHook(clientset *kubernetes.Clientset, mt *confiv1.MultiTenancy, tenant *confiv1.Tenant) error {
	return nil
}

func (s *slackRunner) RunUpdateHook(clientset *kubernetes.Clientset, mt *confiv1.MultiTenancy, tenant *confiv1.Tenant) error {
	return nil
}

func (s *slackRunner) RunDeleteHook(clientset *kubernetes.Clientset, mt *confiv1.MultiTenancy, tenant *confiv1.Tenant, logLines []string) error {
	attachment1 := slackhook.Attachment{}
	attachment1.AddField(slackhook.Field{Title: "Status", Value: "Deleted"}).
		AddField(slackhook.Field{Title: "Tenant", Value: tenant.GetName()}).
		AddField(slackhook.Field{Title: "Multitenancy", Value: mt.GetName()})
	text := fmt.Sprintf("%s tenant %s in namespace %s has been deleted", mt.GetName(), tenant.GetName(), tenant.GetNamespace())
	if len(logLines) > 0 {
		text = text + fmt.Sprintf("\n\n*Final Logs*\n```\n%s```\n", strings.Join(logLines, "\n"))
	}
	payload := slackhook.Payload{
		Text:        text,
		Username:    "multitenancy",
		IconEmoji:   ":monkey_face:",
		Attachments: []slackhook.Attachment{attachment1},
	}
	errs := slackhook.Send(s.config.WebhookURL, "", payload)
	if len(errs) > 0 {
		return fmt.Errorf("%+v", errs)
	}
	return nil
}
