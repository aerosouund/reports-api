package v1alpha1

import (
	"fmt"
	"strings"

	corev1 "k8s.io/api/core/v1"
)

func (r *ReportResult) GetResource() *corev1.ObjectReference {
	if len(r.Subjects) == 0 {
		return nil
	}

	return &r.Subjects[0]
}

func (r *ReportResult) HasResource() bool {
	return len(r.Subjects) > 0
}

func (r *ReportResult) GetKind() string {
	if !r.HasResource() {
		return ""
	}

	return r.GetResource().Kind
}

func (r *ReportResult) GetID() string {
	return r.ID
}

func (r *ReportResult) ResourceString() string {
	if !r.HasResource() {
		return ""
	}

	return ToResourceString(r.GetResource())
}

func ToResourceString(res *corev1.ObjectReference) string {
	var resource string

	if res.Namespace != "" {
		resource = res.Namespace
	}

	if res.Kind != "" && resource != "" {
		resource = fmt.Sprintf("%s/%s", resource, strings.ToLower(res.Kind))
	} else if res.Kind != "" {
		resource = strings.ToLower(res.Kind)
	}

	if res.Name != "" && resource != "" {
		resource = fmt.Sprintf("%s/%s", resource, res.Name)
	} else if res.Name != "" {
		resource = res.Name
	}

	return resource
}
