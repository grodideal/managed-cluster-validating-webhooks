// pkg/webhooks/resourcequota.go

package resourcequota

import (
	admissionregv1 "k8s.io/api/admissionregistration/v1"
	corev1 "k8s.io/api/core/v1"
	admissionctl "sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// ResourceQuotaWebhookName is the name of the resource quota webhook
const ResourceQuotaWebhookName = "resourcequota-validation"

// ResourceQuotaWebhook implements the Webhook interface for resource quotas.
type ResourceQuotaWebhook struct {
	decoder *admissionctl.Decoder
}

// NewWebhook creates a new webhook for ResourceQuota validation
func NewResourceQuotaWebhook(decoder *admissionctl.Decoder) Webhook {
	return &ResourceQuotaWebhook{decoder: decoder}
}

func (w *ResourceQuotaWebhook) Authorized(request admissionctl.Request) admissionctl.Response {
	return admissionctl.Allowed("Authorized")
}

func (w *ResourceQuotaWebhook) GetURI() string {
	return "/validate-resourcequota"
}

func (w *ResourceQuotaWebhook) Validate(req admissionctl.Request) bool {
	resourceQuota := &corev1.ResourceQuota{}
	err := w.decoder.Decode(req, resourceQuota)
	if err != nil {
		return false
	}
	// Implement your validation logic here
	// For example, check if the ResourceQuota is targeting managed namespaces
	return true
}

func (w *ResourceQuotaWebhook) Name() string {
	return ResourceQuotaWebhookName
}

func (w *ResourceQuotaWebhook) FailurePolicy() admissionregv1.FailurePolicyType {
	return admissionregv1.Fail
}

// Define other required methods based on the Webhook interface...

// init function to register this webhook
func init() {
	Register(ResourceQuotaWebhookName, func() Webhook { return &ResourceQuotaWebhook{} })
}
