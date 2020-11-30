module github.com/3scale/k8sapp-initiative/web-server

go 1.13

require (
	k8s.io/apimachinery v0.17.2
	k8s.io/client-go v0.17.2

)

replace (
	github.com/3scale/k8sapp-initiative/web-server/pkg/apis/k8sinitiative.3scale.net/v1alpha1 => ./pkg/apis/k8sinitiative.3scale.net/v1alpha1
	github.com/3scale/k8sapp-initiative/web-server/pkg/pages => ./pkg/pages
)
