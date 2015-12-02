This is a webservice using following
 - golang's net/http
 - golang's html/template
 - go-sqlite driver, which connects to sqlite
 - gorilla mux router

To use this app following are pre-requisites
 - go should be installed from here 'https://golang.org/dl/'
 - get gorilla mux using 'go get github.com/gorilla/mux'
 - get go-sqlite using 'go get github.com/mattn/go-sqlite3'
 - build service using 'go install pre-sales-huddle-webservice'
 - copy pre-sales-huddle.db in bin directory
 - run the service from bin directory './pre-sales-huddle-webservice'

This service is used by an ios client making REST calls
to this. Although, can also be invoked from any client making
REST calls.

Supported REST APIs
The root is http://localhost:8080/
Supported routes are
GET ->   "/prospect/all/"
          "/prospect/view/{criteria}"
POST ->  "/prospect/"
PUT  ->  "/prospect/"


GET ->  "/participant/all/"
    ->  "/participant/userid/{userid}"
    ->  "/participant/prospectid/{id:[0-9]+}"
POST->  "/participant/"
PUT ->  "/participant/"

GET ->  "/discussion/all/"
    ->  "/discussion/view/prospectid/{id:[0-9]+}"
POST->  "/discussion/"
PUT ->  "/discussion/"

 Table Schema

 CREATE TABLE "prospects" (
  `ProspectID`INTEGER PRIMARY KEY AUTOINCREMENT,
  `Name`TEXT,
  `ConfDate`TEXT,
  `TechStack`TEXT,
  `Domain`TEXT,
  `DesiredTeamSize`INT,
  `Notes`TEXT,
  `SalesID`INT,
  `CreateDate`TEXT,
  `StartDate`TEXT,
  `BUHead`TEXT,
 `TeamSize`INT);

CREATE TABLE "participants" (
`ProspectID` INTEGER,
`UserID` TEXT,
`Included` TEXT,
`Participation` TEXT,
 FOREIGN KEY(ProspectID) REFERENCES prospects(ProspectID));

CREATE TABLE "discussions" (
`DiscussionID` INTEGER PRIMARY KEY AUTOINCREMENT,
`ProspectID` INTEGER,
`UserID` TEXT,
`Query` TEXT,
`Answer` TEXT,
 FOREIGN KEY(ProspectID) REFERENCES prospects(ProspectID));
