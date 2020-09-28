# Native K8s application

**Overall Goal**: Proof of Concept. Develop a blueprint for the app architecture that runs on top of CRDs (and maybe other resources) and document do’s and don’ts along the way.

Stages:
1. Create a CRD of a product, backend or API spec. Consume CR data from another pod. (with an image that at least has curl)
2. Instead of using an image that has curl, make it have a server and present a page that consumes the CR
3. Instead of this simple page, use Portafly
4. [Optional] Authorize the calls. We may use Keycloak. We might also explore using only OpenShift RBAC and users.
