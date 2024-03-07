package resourcequota

import (
	"log"
	"net/http"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	admissionctl "sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

const (
	WebhookName string = "resourcequota-validation"
	docString   string = `Webhook to prevent CU to apply ResourceQuotas on Openshift Managed Namespaces`
)

// NewWebhook creates the new webhook for ResourceQuotas
func NewWebhook() *ResourceQuotaWebhook {
	scheme := runtime.NewScheme()
	// Add the schemes for the AdmissionReview and corev1 resources, as these are likely
	// needed for handling ResourceQuota admission requests
	err := admissionv1.AddToScheme(scheme)
	if err != nil {
		log.Fatalf("Failed to add admissionv1 scheme to ResourceQuotaWebhook: %v", err)
	}
	err = corev1.AddToScheme(scheme)
	if err != nil {
		log.Fatalf("Failed to add corev1 scheme to ResourceQuotaWebhook: %v", err)
	}

	return &ResourceQuotaWebhook{
		s: *scheme,
	}
}

// Webhook logic
type ResourceQuotaPreventer struct {
	// Embed necessary structs here (Decoder, Client, etc.)
}

func (s *ResourceQuotaPreventer) Authorized(request admissionctl.Request) admissionctl.Response {
	// Assuming you have a method to check if a namespace is managed
	isManagedNamespace, err := s.isNamespaceManaged(request.Namespace)
	if err != nil {
		// Handle error, possibly logging it and returning an errored response
		return admissionctl.Errored(http.StatusInternalServerError, err)
	}

	if isManagedNamespace {
		return admissionctl.Denied("Applying ResourceQuotas to OpenShift Managed Namespaces is not allowed.")
	}

	return admissionctl.Allowed("Request is allowed")
}

// Example implementation of isNamespaceManaged
func (s *ResourceQuotaPreventer) isNamespaceManaged(namespace string) (bool, error) {
	// Logic to determine if the namespace is managed
	// This is just a placeholder. Your actual implementation will vary.
	return false, nil
}

func (s *ResourceQuotaPreventer) GetURI() string {
	// URI where the webhook will be served
	return "/prevent-resource-quotas"
}

// Implement other required methods following the framework's pattern
