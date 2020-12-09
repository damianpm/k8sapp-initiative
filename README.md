# Native K8s application

**Overall Goal**: Proof of Concept. Develop a blueprint for the app architecture that runs on top of CRDs (and maybe other resources) and document do’s and don’ts along the way.

Stages:
1. Create a CRD of a product, backend or API spec. Consume CR data from another pod. (with an image that at least has curl)
2. Instead of using an image that has curl, make it have a server and present a page that consumes the CR
3. Instead of this simple page, use Portafly
4. [Optional] Authorize the calls. We may use Keycloak. We might also explore using only OpenShift RBAC and users.


## Setup

### Overview
To run this project you need acces to a kubernetes cluster.
We are using an **OpenShift 4.5** cluster, but you'd be able to use another kubernetes cluster.

You can use both [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/) or [OpenShift cli](https://docs.openshift.com/container-platform/4.5/cli_reference/openshift_cli/getting-started-cli.html) tools, in this document we use the OpenShift cli, but you may use `kubectl` instead of `oc`.

### Creating a namespace

To avoid namespace pollution, create a [kubernetes namespace](https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/) called *k8sinitiative* in your cluster: 

```
oc apply -f k8sinitiative_namespace.yml
```

### Creating the Custom Resource Definition of a Product

This project uses [kubernetes CRD's](https://kubernetes.io/docs/tasks/extend-kubernetes/custom-resources/custom-resource-definitions/), to create a representation of a *Product* in your cluster run:

```
oc apply -f product_crd.yml
```

### Creating some Product Custom Resources

To create a couple of *Poducts* you can use the provided file `products_examples.yml`:

```
oc apply -f products_examples.yml
```

### Setting the right permissions

This project makes use of kubernetes [Sevice Account](https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/), and [RBAC Authorization](https://kubernetes.io/docs/reference/access-authn-authz/rbac/) using Role and Role Bindings as a means to authenticate against the apiserver.

Creating the **Service Account**:

```
oc apply -f products_admin_sa.yml
```

Creating the **Roles**:

There are two kind of roles available: Admin and Read-only, to create them run:

```
oc apply -f products_admin_role.yml
oc apply -f products_read_role.yml
```

To create the **Role bindings** for both Admin and Read-only roles run:

```
oc apply -f products_admin_rb.yml
oc apply -f products_read_rb.yml
```

#### Deploy the web Server

We have crafted a [Go Web server](https://github.com/3scale/k8sapp-initiative/tree/master/web-server) to allow you access the *Products* created in the previous steps and display them with a nice UI using [Patternfly](https://www.patternfly.org/v4/).

You don't need to run the web server locally, you can just create a [kubernetes deployment](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/) using the `web-server_deployment.yml` manifest file and the web server will be deployed to your cluster:

```
oc apply -f web-server_deployment.yml
```

### Create a Service

To expose the application you have to create a [kubernetes service](https://kubernetes.io/docs/concepts/services-networking/service/):

```
oc apply -f web-server_service.yml
```

### Create an OpenShift Route
Finally, to create a public URL to access the application you should create an OpenShift route:

```
oc apply -f web-server_route.yaml
```

**Congratulations!**, now you can visit the application at: http://k8sinitiative.apps.dev-eng-ocp4-5.dev.3sca.net/

### [OPTIONAL] Alternatives to OpenShift Routes

If you don't have access to OpenShift or you prefer to use a kubernetes-only solution, you shoud create your own Ingress object.

Here are some useful documentation you can check:

[Kubernetes Ingress documentation](https://kubernetes.io/docs/concepts/services-networking/ingress/)

[Kubernetes Ingress vs OpenShift Route](https://www.openshift.com/blog/kubernetes-ingress-vs-openshift-route)

### [OPTIONAL] Running the web server locally

If you want to run the web server locally, for testing, debugging purposes or to collaborate with the project, you can:

**Build and tag container image:**

Change directory:
```
cd web-server/
```
Build and tag the image:
```
DOCKER_BUILDKIT=1 docker build -f Dockerfile -t <YOUR_USER>/<IMAGE_NAME> .
```
**Run Go web server:**

Change directory:
```
cd pkg/
```
Run server:

```
go run main.go
```
**Note**: You have to setup the needed enviroment variables in your cluster, such as KUBERNETES_SERVICE_HOST and KUBERNETES_SERVICE_PORT.
