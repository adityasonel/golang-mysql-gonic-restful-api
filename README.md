# golang-mysql-gonic-restful-api
golang mySQL gin-gonic restful api example

## API ENDPOINTS

### All Notes
- Path : `v1/notes`
- Method: `GET`
- Response: `200`

### Add note
- Path : `v1/add`
- Method: `POST`
- Fields: `id, content, create_time, modify_time, deleted`
- Response: `201`

### Update Note
- Path : `v1/note/{id}`
- Method: `PUT`
- Fields: `content, modify_time`
- Response: `200`

### Delete note
- Path : `v1/note/{id}`
- Method: `DELETE`
- Response: `204`

## Required Packages
- Dependency management
    * [dep](https://github.com/golang/dep)
- Database
    * [MySql](https://github.com/go-sql-driver/mysql)
- Routing
    * [Gin-Gonic](https://github.com/gin-gonic)
