This is a go webservice with mongoDB database

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
