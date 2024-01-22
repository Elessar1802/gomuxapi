# Some caveats to using minikube 
unless we build the docker image after executing
eval $(minikube docker-env) the kubernetes cluster won't be able
to find the local images
since we are using imagePullPolicy: Never we must do this everytime

We can run the app by `helm install app .`
We can delete all the running deployments and services by `helm delete app`

When we create a service, the env variables of the pods are updated
to reflect a few addtional variables that provide the service_pod ip addr and corresponding port
{{SERVICE_NAME}}_SERVICE_HOST # For ClusterIP
{{SERVICE_NAME}}_SERVICE_PORT # For ClusterIP
So I needed to update the main.go to relfect the above changes
