package resourcequota

import (
	"testing"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	admissionctl "sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// TestResourceQuotaWebhookValidation tests the Validate function for different scenarios
func TestResourceQuotaWebhookValidation(t *testing.T) {
	decoder, _ := admissionctl.NewDecoder(scheme)
	webhook := NewResourceQuotaWebhook(decoder)

	// Define test cases
	tests := []struct {
		name       string
		quota      corev1.ResourceQuota
		expectPass bool
	}{
		{
			name: "Valid ResourceQuota",
			quota: corev1.ResourceQuota{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-quota",
					Namespace: "test-namespace",
				},
				// Define valid ResourceQuota spec here
			},
			expectPass: true,
		},
		{
			name: "Invalid ResourceQuota - Managed Namespace",
			quota: corev1.ResourceQuota{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-quota",
					Namespace: "openshift-managed", // Assuming this is a managed namespace
				},
				// Define ResourceQuota spec here
			},
			expectPass: false,
		},
		// Add more test cases as needed
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Encode the ResourceQuota object to simulate an admission request
			req := admissionctl.Request{
				// Populate the request object; you may need to mock or simulate an AdmissionRequest here
			}
			resp := webhook.Authorized(req)
			if resp.Allowed != tc.expectPass {
				t.Errorf("%s: expected %v, got %v", tc.name, tc.expectPass, resp.Allowed)
			}
		})
	}
}
