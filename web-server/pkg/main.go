package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"

	"github.com/damianpm/k8sapp-initiative/web-server/pkg/apis/k8sinitiative.3scale.net/v1alpha1"
	products "github.com/damianpm/k8sapp-initiative/web-server/pkg/pages"
)

const homepageEndPoint = "/"
const productsEndPoint = "/products"

var restClient *rest.RESTClient

func setupClient() *rest.RESTClient {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	crdConfig := *config
	crdConfig.ContentConfig.GroupVersion = &schema.GroupVersion{Group: v1alpha1.GroupName, Version: v1alpha1.GroupVersion}
	crdConfig.APIPath = "/apis"
	crdConfig.NegotiatedSerializer = serializer.NewCodecFactory(scheme.Scheme)
	crdConfig.UserAgent = rest.DefaultKubernetesUserAgent()

	restClient, err := rest.RESTClientFor(&crdConfig)
	if err != nil {
		panic(err)
	}
	return restClient
}

func init() {
	v1alpha1.AddToScheme(scheme.Scheme)
	restClient = setupClient()
}

// StartWebServer the webserver
func StartWebServer() {
	http.HandleFunc(homepageEndPoint, handleHomePage)
	http.HandleFunc(productsEndPoint, handleProductsPage)
	fs := http.FileServer(http.Dir("./templates/styles"))
	http.Handle("/templates/styles/", http.StripPrefix("/templates/styles/", fs))
	port := os.Getenv("PORT")
	if len(port) == 0 {
		panic("Environment variable PORT is not set")
	}

	log.Printf("Starting web server to listen on endpoints [%s] and port %s",
		homepageEndPoint, port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}

type IndexTemplateData struct {
	Title string
}

func handleHomePage(w http.ResponseWriter, r *http.Request) {
	urlPath := r.URL.Path
	log.Printf("Web request received on url path %s", urlPath)
	indexTemplateData := IndexTemplateData{Title: "k8s initiative"}
	temp, err := template.ParseFiles("templates/layout.html", "templates/index.html")

	if err != nil {
		fmt.Printf("Failed to write response, err: %s", err)
	}
	temp.ExecuteTemplate(w, "layout", indexTemplateData)
}

func handleProductsPage(w http.ResponseWriter, r *http.Request) {
	result := v1alpha1.ProductList{}

	getErr := restClient.
		Get().
		Namespace("damian-k8s-initiative").
		Resource("products").
		Do().
		Into(&result)

	if getErr != nil {
		panic(getErr)
	}

	fmt.Printf("%d results found: %+v\n", len(result.Items), result)
	fmt.Println(getErr)

	products.Index(w, result)
}

func main() {
	StartWebServer()
}
