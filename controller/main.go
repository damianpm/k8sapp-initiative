package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/3scale/k8sapp-initiative/pkg/apis/k8sinitiative.3scale.net/v1alpha1"

	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

var kubeconfig string

func main() {

	ns := "k8sinitiative"

	s := scheme.Scheme
	v1alpha1.AddToScheme(s)

	configuration := config.GetConfigOrDie()
	cl, err := client.New(configuration, client.Options{Scheme: s})
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create client: %v\n", err)
		os.Exit(1)
	}

	listOps := []client.ListOption{
		client.InNamespace(ns),
	}

	productList := &v1alpha1.ProductList{}
	err = cl.List(context.TODO(), productList, listOps...)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error", err)
		os.Exit(1)
	}

	jsonData, err := json.MarshalIndent(productList, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf(string(jsonData))
}
