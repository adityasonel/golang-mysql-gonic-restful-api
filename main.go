package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Note struct {
	Id         int
	Content    string
	CreateTime int
	ModifyTime int
	Deleted    int
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "groot@123"
	dbName := "note_db"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@(localhost:3306)/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func Notes(c *gin.Context) {
	db := dbConn()
	selectQuery, err := db.Query("SELECT * FROM notes")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"data":    "",
			"success": "true",
		})
		panic(err.Error())
	}
	note := Note{}
	result := []Note{}
	for selectQuery.Next() {
		var id int
		var content string
		var createTime int
		var modifyTime int
		var deleted int

		err = selectQuery.Scan(&id, &content, &createTime, &modifyTime, &deleted)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"data":    "",
				"success": "false",
			})
			panic(err.Error())
		}

		note.Id = id
		note.Content = content
		note.CreateTime = createTime
		note.ModifyTime = modifyTime
		note.Deleted = deleted

		result = append(result, note)
	}
	c.JSON(http.StatusOK, gin.H{
		"error":   "",
		"data":    result,
		"success": "true",
	})

	defer db.Close()
}

func Add(c *gin.Context) {
	db := dbConn()
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"data":    "",
			"success": "false",
		})
		panic(err.Error())
	}
	bodyMap := make(map[string]interface{})
	err1 := json.Unmarshal([]byte(string(body)), &bodyMap)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"data":    "",
			"success": "false",
		})
		panic(err1.Error())
	}
	insertQuery, err2 := db.Prepare("INSERT INTO notes(content, create_time, modify_time) VALUES(?,?,?)")
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err2.Error(),
			"data":    "",
			"success": "false",
		})
		panic(err.Error())
	}
	_, err3 := insertQuery.Exec(bodyMap["content"], bodyMap["create_time"], bodyMap["modify_time"])
	if err3 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err3.Error(),
			"data":    "",
			"success": "false",
		})
		panic(err3.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"error":   "",
		"data":    "",
		"success": "true",
	})

	defer db.Close()
}

func Modify(c *gin.Context) {
	db := dbConn()

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error":   err.Error(),
			"data":    "",
			"success": "false",
		})
		panic(err.Error())
	}
	bodyMap := make(map[string]interface{})
	err1 := json.Unmarshal([]byte(string(body)), &bodyMap)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err1.Error(),
			"data":    "",
			"success": "false",
		})
		panic(err1.Error())
	}
	updateQuery, err2 := db.Prepare("UPDATE notes SET content = ?, modify_time = ? WHERE id = ?")
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err2.Error(),
			"data":    "",
			"success": "false",
		})
		panic(err2.Error())
	}
	_, err3 := updateQuery.Exec(bodyMap["content"], bodyMap["modify_time"], bodyMap["id"])
	if err3 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err3.Error(),
			"data":    "",
			"success": "false",
		})
		panic(err3.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"error":   "",
		"data":    "",
		"success": "true",
	})

	defer db.Close()
}

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/notes", Notes)
		v1.POST("/add", Add)
		v1.POST("/modify", Modify)
	}

	router.Run()
}
