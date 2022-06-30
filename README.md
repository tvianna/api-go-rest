# api-go-rest
Study project: API rest with Go.

## Description

This project was developed with GoLang and PostgreSQL.

## Install
## Starting the containers
Access the root directory and execute:
```bash
# development
$ docker-compose up
```

Two services will be created:
  - postgres:
    - this service will create the postgres container;
    - the 'personalidades' table will be created automatically.
  - pgadmin-compose:
    - this service will create the postgres admin to access the database.
    
## Download to local cache all dependent modules
Access the root directory and execute:
```bash
$  go mod download
```
    
## Starting the application
Access the root directory and execute:
```bash
$  go run main.go
```

## Created Paths
## Personalidades
- Create 'personalidade'

  POST
  /api/personalidades
  
```bash
# body request example
  {
    "id": 3,
    "nome": "Anitta",
    "historia": "A maior do Brasil."
  }
```
- Edit 'personalidade'

  PUT
  /api/personalidades/{id}
  
```bash
# body request example
  {
    "nome": "Anitta 2.0",
    "historia": "A maior do Brasil. 2.0"
  }
```

- Delete 'personalidade'

DELETE
  /api/personalidades/{id}
  
  
- List all 'personalidade'

GET
  /api/personalidades/
  
- List 'personalidade' by id 

GET
  /api/personalidades/{id}

#
Study project developed by [Taís Vianna](https://www.linkedin.com/in/taiisvianna/).
