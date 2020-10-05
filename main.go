package main

import (
	"os"
	"path/filepath"

	"context"
	"fmt"
	"flag"

	"k8s.io/apimachinery/pkg/runtime/schema"
 	"k8s.io/apimachinery/pkg/runtime/serializer"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/3scale/k8sapp-initiative/pkg/apis/k8sinitiative.3scale.net/v1alpha1"
)

var kubeconfig string

func main() {
	flag.StringVar(&kubeconfig, "kubeconfig", kubeconfigPath(), "")
	flag.Parse()

	config, err := rest.InClusterConfig()
	if err != nil {
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			panic(err)
		}
	}


	v1alpha1.AddToScheme(scheme.Scheme)
	crdConfig := *config
	crdConfig.ContentConfig.GroupVersion = &schema.GroupVersion{Group: v1alpha1.GroupName, Version: v1alpha1.GroupVersion}
	crdConfig.APIPath = "/apis"
	crdConfig.NegotiatedSerializer = serializer.NewCodecFactory(scheme.Scheme)
	crdConfig.UserAgent = rest.DefaultKubernetesUserAgent()
	threescaleRestClient, err := rest.RESTClientFor(&crdConfig)
	if err != nil {
		panic(err)
	}

	result := v1alpha1.ProductList{}
	err = threescaleRestClient.Get().Namespace("k8sinitiative").Resource("products").Do(context.TODO()).Into(&result)


	fmt.Printf("%d results found: %+v\n", len(result.Items), result)
	fmt.Println(err)
}

func kubeconfigPath() string {
	fname := os.Getenv("KUBECONFIG")
	if fname != "" {
		return fname
	}
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("The kubeconfig cannot be loaded: %v\n", err)
		panic(err)
	}
	return filepath.Join(home, ".kube", "config")
}
