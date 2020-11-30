package products

import (
	"html/template"
	"log"
	"net/http"

	"github.com/3scale/k8sapp-initiative/web-server/pkg/apis/k8sinitiative.3scale.net/v1alpha1"
)

// Index prints the list of products
func Index(w http.ResponseWriter, productList v1alpha1.ProductList) {
	const tpl = `
	<!DOCTYPE html>
	<html>
		<head>
			<meta charset="UTF-8">
			<title>Products</title>
		</head>
		<body>
			<table>
				<tr>
					<th>ID</th>
					<th>Name</th>
					<th>Description</th>
					<th>Nº of Backends</th>
					<th>Nº of Apps</th>
				</tr>
				{{range .Items}}
					<tr>
						<td>{{ .Spec.ID }}</td>
						<td>{{ .Spec.Name }}</td>
						<td>{{ .Spec.Description }}</td>
						<td>{{ .Spec.BackendsCount }}</td>
						<td>{{ .Spec.AppsCount }}</td>
					</tr>
				{{end}}
			</table>
		</body>
	</html>`

	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	t, err := template.New("webpage").Parse(tpl)
	check(err)

	err = t.Execute(w, productList)
	check(err)
}
