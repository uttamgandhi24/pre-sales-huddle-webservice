This is a go webservice with mongoDB database

Packages and their use
- gopkg.in/mgo.v2

  MongoDB database driver. More information available at
  https://godoc.org/gopkg.in/mgo.v2
  Get this package by 'go get gopkg.in/mgo.v2'

- github.com/gorilla/mux

  Gorilla mux router, for handling parameterised routes and restricting
  http methods among other functions. More information is available at
  http://www.gorillatoolkit.org/pkg/mux
  Get this package by 'go get github.com/gorilla/mux'

Supported REST APIs
The root is http://localhost:8080/
Supported routes are

Prospect
- GET  "/prospect/all/"
- POST "/prospect/"
- PUT  "/prospect/"

Participant
- GET  "/participant/all/"
- GET  "/participant/userid/{userid}"
- GET  "/participant/prospectid/{id}"
- POST "/participant/"
- PUT  "/participant/"

Discussion
- GET  "/discussion/all/"
- GET  "/discussion/prospectid/{id}"
- POST "/discussion/"
- PUT  "/discussion/"

User
- GET  "/user/all/"
- GET  "/user/email/{email}"
- POST "/user/"

Run the unit tests
- go test -test.v pre-sales-huddle-webservice