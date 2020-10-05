package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type Product struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Spec   ProductSpec
}

type ProductSpec struct {
	Description string
}

type ProductList struct {
	metav1.TypeMeta
	metav1.ListMeta

	Items []Product
}
