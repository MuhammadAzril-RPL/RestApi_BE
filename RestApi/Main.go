package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

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
	// Koneksi database
	db, err := sql.Open("mysql", "Azril:Myboo5456@tcp(localhost:3307)/wilayahs")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Membuat router dengan Gin framework
	router := gin.Default()

	// Endpoint GET /province
	router.GET("/province", func(c *gin.Context) {
		// Mengambil data dari database
		rows, err := db.Query("SELECT id, code, name FROM provinces")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var province []Province
		for rows.Next() {
			var p Province
			if err := rows.Scan(&p.ID, &p.Code, &p.Name); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			province = append(province, p)
		}

		// Menyiapkan respons
		response := APIResponse{
			Status:  "sukses",
			Code:    200,
			Message: "Berhasil mendapatkan data",
			Data:    province,
		}

		// Mengirim respons dalam format JSON
		c.JSON(http.StatusOK, response)
	})

	log.Println("Server berjalan di http://localhost:8000")
	if err := router.Run(":8000"); err != nil {
		log.Fatal(err)
	}
}
