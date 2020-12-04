package products

import (
	"html/template"
	"log"
	"net/http"

	"github.com/3scale/k8sapp-initiative/web-server/pkg/apis/k8sinitiative.3scale.net/v1alpha1"
)

// Index prints the list of products
func Index(w http.ResponseWriter, productList v1alpha1.ProductList) {

	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}

	temp, err := template.ParseFiles("../templates/layout.html", "../templates/products.html")
	check(err)

	err = temp.ExecuteTemplate(w, "layout", productList)
	check(err)
}
