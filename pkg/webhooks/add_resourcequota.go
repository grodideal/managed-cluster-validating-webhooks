package webhooks

import (
	"github.com/openshift/managed-cluster-validating-webhooks/pkg/webhooks/resourcequota"
)

func init() {
	Register(resourcequota.WebhookName, func() Webhook { return resourcequota.NewWebhook() })
}
