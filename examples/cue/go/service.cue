import "k8s.io/api/core/v1"

services: [string]: v1.#Service

services: {
    http: {
        // Next is shows an error
        // unknown: "name"
        apiVersion: "v1"
        kind: "Service"
        metadata: {
            name: "http-service"
            namespace: "app"
        }
        spec: {
            selector: {
                "app.kubernetes.io/name": "app"
            }
            ports: [
                {
                    protocol: "TCP"
                    port: 80
                    targetPort: 80
                }
            ]
        }
    }
}
