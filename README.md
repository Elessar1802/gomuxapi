# User attendance API service using Gorilla Mux, Postgresql & JWT
This API allows an Admin user (Principal) to add new users (students/teachers). The Users can punch in/out their attendance.
The principal can view the attendance of the teachers while the teachers can view the attendance of students either by
grouping the students by their class or using their individual user ids. Students can only view their own attendances.
The logic to handle the authorizations is handled using **decorators functions pattern**.

# Steps to running it
Before deploying the helm chart you need to create a configmap and secret on your minikube cluster using
`kubectl apply -f k8/configmap.yaml -f k8/secrets.yaml`. Change the values to your own liking.
If your using minikube, deploy the chart in `cd chart/` folder by running the command `helm install .`
Then run the command `minikube service app-be-service --url` and use the returned url to access the api.

# API endpoints
```
POST: /login // need to login before accessing the below endpoints since this is the only endpoint not protected using a token
// all other endpoints require a token as an AUTHORIZATION header or as a cookie set by the /login endpoint
GET: /users // to list all the users
POST: /users // add new users (can only be done by the admin user)
GET: /attendance/user/{user_id}?start_date=2024/01/01&end_date=2024/01/02 // fetch attendance record for an user between 01/01/2024 and 02/01/2024
POST: /attendance/user/{user_id} // punch in
PUT: /attendance/user/{user_id} // punch out
GET: /attendance/class/{class}?start_date=2024/01/01&end_date=2024/01/02 // fetch attendance record for all the users in {class} between 01/01/2024 and 02/01/2024
```

# Interesting learning points
## How to handle initializing DB only once without using WIRE:
Create a Handler struct that stores a pointer to the DB. Make all the http handlers a method to this Handler struct.
Initialize a `Handler{}` with the reference to the DB created in the main func. Now the handlers always have a reference to this DB instance.
The concept of **closures** is in play here.

## How to handle authorization:
Authentication is being checked in the Authentication middleware that checks if the JWT signature is valid.
The authorization of the different endpoints is checked using the *decorator functions* defined in the `src/v1/handlers/decorators.go` 
to keep the authorization checks manageable.

## Using joins in go-pg
```
  er := db.Model().Table("users").
    ColumnExpr("users.id, users.name, users.phone").
    Join("JOIN students ON students.id = users.id and students.class = ?", name).
    Select(&users)
```

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
