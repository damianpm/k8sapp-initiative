package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

//go:generate controller-gen object paths=$GOFILE

type Product struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProductSpec `json:"spec"`
}

type ProductSpec struct {
	ID            *int    `json:"id"`
	Description   *string `json:"description"`
	AppsCount     *int    `json:"apps_count"`
	BackendsCount *int    `json:"backends_count"`
	Name          *string `json:"name"`
	SystemName    *string `json:"system_name"`
	UnreadAlerts  *int    `json:"unread_alerts_count"`
	UpdatedAt     *string `json:"updated_at"`
	Link          *string `json:"link"`
}

type ProductList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Product `json:"items"`
}
