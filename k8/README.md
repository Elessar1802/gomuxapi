Since I am using minikube I need to additionally
run the command `minikube service app-be-service`
to allot it an external ip address.
I will continue to get <pending> status under External-IP since
I am using Minikube.

configMap stores key, value pairs in plain text.
Secrets contains key, value pairs in base64 encoded text. 
To encode something in base64 i need to run `echo -n 'sometext' | base64`


