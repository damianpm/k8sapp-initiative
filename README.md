# Native K8s application

To learn more about the goals and stages of this initiative plese look at our [About file](ABOUT.md).
## Setup

### Overview
To run this project you need acces to a kubernetes cluster.
We are using an **OpenShift 4.5** cluster, but you'd be able to use another kubernetes cluster.

You can use both [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/) or [OpenShift cli](https://docs.openshift.com/container-platform/4.5/cli_reference/openshift_cli/getting-started-cli.html) tools, in this document we use the OpenShift cli, but you may use `kubectl` instead of `oc`.

### Creating a namespace

To avoid using your default or previous working [namespace](https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/), let's create a new one called *k8sinitiative* in your cluster: 

```
oc apply -f k8sinitiative_namespace.yml
```

### Creating the Custom Resource Definition of a Product

This project uses [kubernetes CRD's](https://kubernetes.io/docs/tasks/extend-kubernetes/custom-resources/custom-resource-definitions/) to create a representation of a *Product* (based on 3scale Products). To create the CRD in your cluster run:

```
oc apply -f product_crd.yml
```

### Creating some Product Custom Resources

To create a couple of *Poducts* you can use the provided file `products_examples.yml`:

```
oc apply -f products_examples.yml
```

### Setting the right permissions

We are going to create a [Sevice Account](https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/), to provide special [RBAC Authorization](https://kubernetes.io/docs/reference/access-authn-authz/rbac/) using a custom role and binding it as a means to authenticate against the apiserver.

Creating the **Service Accounts**:

There are two kind of Service Accounts: Admin (full access) and Viewer (read-only).

To create the Admin SA, run:

```
oc apply -f products_admin_sa.yml
```

To create the Viewer SA, run:

```
oc apply -f products_viewer_sa.yml
```

Creating the **Roles**:

There are two kind of roles available, corrresponding with the Admin and Viewer SA's, to create them run:

```
oc apply -f products_admin_role.yml
oc apply -f products_read_role.yml
```

To create the **Role bindings** for both Admin and Viewer roles run:

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

If you want to [modify the web server](#optional-modify-the-web-server), see section below.

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

**Important notes:**

- The [host field](https://github.com/3scale/k8sapp-initiative/blob/master/web-server_route.yaml#L7) in `web-server_route.yaml` is currently pointing to a 3scale owned domain. You should change this to fit your needs, maybe using an environment variable accessible in your cluster.

- *Routes* are specific to OpenShift, for alternatives see [alternatives to OpenShift](#optional-alternatives-to-openshift-routes) below.

### [OPTIONAL] Alternatives to OpenShift Routes

If you don't have access to OpenShift or you prefer to use a kubernetes-only solution, you shoud create your own Ingress object.

Here are some useful documentation you can check:

[Kubernetes Ingress documentation](https://kubernetes.io/docs/concepts/services-networking/ingress/)

[Kubernetes Ingress vs OpenShift Route](https://www.openshift.com/blog/kubernetes-ingress-vs-openshift-route)

### [OPTIONAL] Modify the web server

#### Configuration

The web server is designed to run in the context of a kubernetes cluster and needs access to some special configuration and environment variables such as KUBERNETES_SERVICE_HOST, KUBERNETES_SERVICE_PORT and PORT, **is your responsiblity to setup this configuration**.

We have created a [controller](https://github.com/3scale/k8sapp-initiative/tree/master/controller) using the offical **Go client** for talking to the kubernetes cluster, you can find more info at https://github.com/kubernetes/client-go and https://pkg.go.dev/k8s.io/client-go.

Once you are done with your modifications, you should build, tag and push your image and update the deployment.

**Build, tag and push container image:**

Change directory:
```
cd web-server/
```
Build and tag the image:
```
DOCKER_BUILDKIT=1 docker build -f Dockerfile -t <YOUR_USER>/<IMAGE_NAME> .
```

Then push the image to your registry:
```
docker push <YOUR_USERNAME>/<IMAGE_NAME>
```

**Update the image container in the deployment**

In the `web-server_deployment.yml` file, you should change the [container image field](https://github.com/3scale/k8sapp-initiative/blob/master/web-server_deployment.yml#L20) to use the one you built in the previous step.

