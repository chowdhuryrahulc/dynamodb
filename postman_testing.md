http://localhost:8080/health    GET    {"status": 200, "result": "Service OK"}
http://localhost:8080/product   POST    
    Input:  body->raw->   {
            "createdAt": "2021-08-09T17:36:50+05:30",
            "updatedAt": "2022-03-03T10:48:07+05:30",
            "name": "shampoo"
        }
http://localhost:8080/product   GET
http://localhost:8080/product/4abfbfed-1b8d-4c48-b054-1e51583e00e1    GET 
http://localhost:8080/product/4abfbfed-1b8d-4c48-b054-1e51583e00e1    PUT
    Input: body->raw-> {
            "createdAt": "2021-08-09T17:36:50+05:30",
            "updatedAt": "2022-03-03T10:48:07+05:30",
            "name": "strawberry"
        }
http://localhost:8080/product/4abfbfed-1b8d-4c48-b054-1e51583e00e1    DELETE