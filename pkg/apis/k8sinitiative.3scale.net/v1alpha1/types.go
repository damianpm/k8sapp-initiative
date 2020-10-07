package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type Product struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProductSpec `json:"spec"`
}

type ProductSpec struct {
	Description string `json:"description"`
}

type ProductList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Product `json:"items"`
}
