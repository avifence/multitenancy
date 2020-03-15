package slack

import (
	"fmt"
	"strings"

	slackhook "github.com/ashwanthkumar/slack-go-webhook"
	confiv1 "github.com/configurator/multitenancy/pkg/apis/confi/v1"
	"github.com/configurator/multitenancy/pkg/eventhooks/runner"
	"github.com/configurator/multitenancy/pkg/util"
	"k8s.io/client-go/kubernetes"
)

var multitenancyEmoji = ":monkey_face:"

type slackRunner struct {
	config *confiv1.SlackConfig
}

func NewRunner(conf *confiv1.SlackConfig) runner.HookRunner {
	return &slackRunner{config: conf}
}

func (s *slackRunner) RunCreateHook(clientset *kubernetes.Clientset, mt *confiv1.MultiTenancy, tenant *confiv1.Tenant) error {
	text := fmt.Sprintf("A new tenant has been created for multitenancy _%s_", mt.GetName())
	return s.send(newPayload(text, newAttachment(mt, tenant, "Updated", "#5A8D03")))
}

func (s *slackRunner) RunUpdateHook(clientset *kubernetes.Clientset, mt *confiv1.MultiTenancy, tenant *confiv1.Tenant) error {
	text := fmt.Sprintf("The _%s_ tenant _%s_ in namespace _%s_ has been updated", mt.GetName(), tenant.GetName(), tenant.GetNamespace())
	return s.send(newPayload(text, newAttachment(mt, tenant, "Updated", "#F9A602")))
}

func (s *slackRunner) RunDeleteHook(clientset *kubernetes.Clientset, mt *confiv1.MultiTenancy, tenant *confiv1.Tenant, logLines []string) error {
	text := fmt.Sprintf("The _%s_ tenant _%s_ in namespace _%s_ has been deleted", mt.GetName(), tenant.GetName(), tenant.GetNamespace())
	if len(logLines) > 0 {
		text = text + fmt.Sprintf("\n\n*Final Logs*\n```\n%s```\n", strings.Join(logLines, "\n"))
	}
	return s.send(newPayload(text, newAttachment(mt, tenant, "Deleted", "#FF0000")))
}

func newAttachment(mt *confiv1.MultiTenancy, tenant *confiv1.Tenant, status, color string) slackhook.Attachment {
	attachment := slackhook.Attachment{}
	attachment.AddField(slackhook.Field{Title: "Status", Value: status}).
		AddField(slackhook.Field{Title: "Tenant", Value: tenant.GetName()}).
		AddField(slackhook.Field{Title: "Multitenancy", Value: mt.GetName()})
	attachment.Color = util.StringPtr(color)
	return attachment
}

func newPayload(text string, attachment slackhook.Attachment) slackhook.Payload {
	return slackhook.Payload{
		Text:        text,
		Username:    "multitenancy",
		IconEmoji:   multitenancyEmoji,
		Attachments: []slackhook.Attachment{attachment},
	}
}

func (s *slackRunner) send(payload slackhook.Payload) error {
	return checkSlackErrors(slackhook.Send(s.config.WebhookURL, "", payload))
}

func checkSlackErrors(errs []error) error {
	if len(errs) > 0 {
		return fmt.Errorf("%+v", errs)
	}
	return nil
}
