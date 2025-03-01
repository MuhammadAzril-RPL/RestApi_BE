package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type ProvinceAPI struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Province struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type APIResponse struct {
	Status  string     `json:"status"`
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    []Province `json:"data"`
}

func main() {
	database, error := sql.Open("mysql", "Azril:Myboo5456@tcp(localhost:3307)/wilayahs")
	if error != nil {
		log.Fatal(error)
	}
	defer database.Close()

	response, error := http.Get("https://emsifa.github.io/api-wilayah-indonesia/api/provinces.json")
	if error != nil {
		log.Fatal(error)
	}
	defer response.Body.Close()

	var provincesAPIList []ProvinceAPI
	if error := json.NewDecoder(response.Body).Decode(&provincesAPIList); error != nil {
		log.Fatal(error)
	}

	for _, province := range provincesAPIList {
		_, error := database.Exec("INSERT INTO provinces (code, name) VALUES (?, ?) ON DUPLICATE KEY UPDATE name = VALUES(name)", province.ID, province.Name)
		if error != nil {
			log.Println("Error inserting:", error)
		}
	}

	router := gin.Default()

	router.GET("/province", func(context *gin.Context) {
		rows, error := database.Query("SELECT id, code, name FROM provinces")
		if error != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
			return
		}
		defer rows.Close()

		var provincesList []Province
		for rows.Next() {
			var province Province
			if error := rows.Scan(&province.ID, &province.Code, &province.Name); error != nil {
				context.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
				return
			}
			provincesList = append(provincesList, province)
		}

		apiResponse := APIResponse{
			Status:  "success",
			Code:    200,
			Message: "Successfully get data",
			Data:    provincesList,
		}
		context.JSON(http.StatusOK, apiResponse)
	})

	log.Println("Server berjalan di http://localhost:8000")
	if error := router.Run(":8000"); error != nil {
		log.Fatal(error)
	}
}
