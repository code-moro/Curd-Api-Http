# CURD API WITH MONGODB USING HTTP
## Description

This is a API built using HTTP. This application is used to perform curd operation on movies_collection in database.
* Go, also known as Golang, is an open-source, compiled, and statically typed programming language designed by Google. It is built to be simple, high-performing, readable, and efficient.

* MongoDB is an open source NoSQL database management program. NoSQL is used as an alternative to traditional relational databases. NoSQL databases are quite useful for working with large sets of distributed data. MongoDB is a tool that can manage document-oriented information, store or retrieve information

----

## Pre-requisites
 Install Golang
 * [Installation for Windows](https://go.dev/doc/install) 
 * [Installation for Linux](https://golangdocs.com/install-go-linux)

 Install Mux
 
 * [Installation Mux](https://github.com/gorilla/mux)
 
 Install MongoDB
 
 * [Installation Guide for Mongodb](https://www.mongodb.com/docs/manual/installation/)


 ---
 ## Getting Started

 1. Clone the repository using:
 ```
 git clone https://github.com/code-moro/Curd-Api-Http.git
 ```
 2. Install HTTP 
 ```
 go get -u github.com/gorilla/mux
 ```
 3. Start Server
```
 go run .\main.go
 localhost:8000/
```
 

---
# API Endpoints

## 1.Get Request 

### API EndPoint : /movie
 
![getrequestphoto](Images\GetMovie.png)

### Output

```
[
    {
        "_id": 1,
        "year": 2008,
        "title": "The Dark Knight",
        "director": "Christopher Nolan",
        "rating": 9
    },
    {
        "_id": 2,
        "year": 1994,
        "title": "The Shawshank Redemption",
        "director": "Frank Darabont",
        "rating": 9.4
    },
    {
        "_id": 3,
        "year": 1994,
        "title": "Forrest Gump",
        "director": "Robert Zemeckis",
        "rating": 8.8
    }
]
```
---

## 2.Post Request

### API EndPoint : /movie

![postrequest](Images\PostMovie.png)

#### Output
```
{
    "InsertedID": 4
}
```

---
## 3.Update Request

### API EndPoint : /movie/{id}

![updaterequest](Images\UpdateMovie.png)

#### Output
```
"Record Updated"
```

---
## 4. Delete Request

### API EndPoint : /movie/{id}

![deleterequest](Images\DeleteMovie.png)

#### Output
```
"User Deleted succesfully"
```
 ----