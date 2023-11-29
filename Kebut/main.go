package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"kebut/database"
	"kebut/entity"
	"log"
	"strings"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Failed to load env")
	}
	db := database.InitDb()
	if err := database.Migrate(db); err != nil {
		log.Fatalln("Failed to migrate")
	}
	//handler.PDummy()
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		if c.Request.Method == "OPTIONS" {
			c.Writer.Header().Set("Content-Type", "application/json")
			c.AbortWithStatus(204)
		} else {
			c.Next()
		}
	})
	r.GET("/test", func(c *gin.Context) {
		var data entity.Penyakit
		if err := database.DB.First(&data).Error; err != nil {
			c.JSON(500, gin.H{
				"message": "failed",
			})
			return
		}
		c.JSON(200, data)
	})

	type Link struct {
		Link string `json:"link"`
	}
	type Result struct {
		Nama       string   `json:"nama"`
		Gejala     []string `json:"gejala"`
		Penanganan []string `json:"penanganan"`
		Links      []Link   `json:"link"`
	}

	type Input struct {
		Value []int `json:"value" binding:"required"`
	}

	r.POST("/get", func(c *gin.Context) {
		var req Input
		if err := c.BindJSON(&req); err != nil {
			c.JSON(400, gin.H{
				"message": "Invalid",
			})
			return
		}
		penyakit := make(map[string]int)

		for _, point := range []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"} {
			penyakit[point] = 0
		}

		reqArr := req.Value
		for i, value := range reqArr {
			switch i {
			case 0:
				if value == 1 {
					penyakit["1"] += 1
					penyakit["8"] += 1
				}
			case 1:
				if value == 1 {
					penyakit["1"] += 1
					penyakit["2"] += 1
					penyakit["3"] += 1
					penyakit["6"] += 1
					penyakit["7"] += 1
				}
			case 2:
				if value == 1 {
					penyakit["1"] += 1
					penyakit["2"] += 1
					penyakit["6"] += 1
					penyakit["7"] += 1
					penyakit["8"] += 1
				}
			case 3:
				if value == 1 {
					penyakit["2"] += 3
				}
			case 4:
				if value == 1 {
					penyakit["2"] += 3
					penyakit["9"] += 3
				}
			case 5:
				if value == 1 {
					penyakit["3"] += 1
					penyakit["2"] += 1
					penyakit["8"] += 1
					penyakit["9"] += 1
				}
			case 6:
				if value == 1 {
					penyakit["5"] += 5
				}
			case 7:
				if value == 1 {
					penyakit["5"] += 5
				}
			case 8:
				if value == 1 {
					penyakit["5"] += 5
				}
			case 9:
				if value == 1 {
					penyakit["4"] += 3
				}
			case 10:
				if value == 1 {
					penyakit["4"] += 1
					penyakit["8"] += 1
				}
			case 11:
				if value == 1 {
					penyakit["4"] += 1
					penyakit["6"] += 1
					penyakit["9"] += 1
				}
			case 12:
				if value == 1 {
					penyakit["6"] += 3
				}
			case 13:
				if value == 1 {
					penyakit["7"] += 3
				}
			case 14:
				if value == 1 {
					penyakit["8"] += 3
				}
			case 15:
				if value == 1 {
					penyakit["9"] += 3
				}
			case 16:
				if value == 1 {
					penyakit["1"] += 3
				}
			case 17:
				if value == 1 {
					penyakit["1"] += 1
					penyakit["3"] += 1
					penyakit["4"] += 1
					penyakit["6"] += 4
					penyakit["7"] += 1
				}
			case 18:
				if value == 1 {
					penyakit["1"] += 1
					penyakit["2"] += 1
					penyakit["6"] += 1
					penyakit["8"] += 1
					penyakit["9"] += 1
				}
			case 19:
				if value == 1 {
					penyakit["2"] += 1
					penyakit["3"] += 1
					penyakit["9"] += 1
				}
			case 20:
				if value == 1 {
					penyakit["2"] += 3
				}
			case 21:
				if value == 1 {
					penyakit["3"] += 3
				}
			case 22:
				if value == 1 {
					penyakit["9"] += 3
				}
			case 23:
				if value == 1 {
					penyakit["4"] += 3
				}
			case 24:
				if value == 1 {
					penyakit["4"] += 1
					penyakit["6"] += 1
					penyakit["7"] += 1
				}
			case 25:
				if value == 1 {
					penyakit["7"] += 3
				}
			case 26:
				if value == 1 {
					penyakit["8"] += 1
				}
			case 27:
				if value == 1 {
					penyakit["8"] += 3
				}
			}
		}
		tampilkanNilaiPenyakit(penyakit)
		disease := cariNilaiTertinggi(penyakit)

		fmt.Println("========")
		fmt.Println(disease)
		fmt.Println("=========")

		var data entity.Penyakit
		if err := database.DB.Model(&entity.Penyakit{}).Where("id = ?", disease).Preload("Links").Find(&data).Error; err != nil {
			c.JSON(500, gin.H{
				"message": "error",
			})
			return
		}

		gejala := strings.Split(data.Gejala, ";")
		penanganan := strings.Split(data.Penanganan, ";")
		var links []Link
		for _, link := range data.Links {
			links = append(links, Link{Link: link.Link})
		}

		result := Result{
			Nama:       data.Nama,
			Gejala:     gejala,
			Penanganan: penanganan,
			Links:      links,
		}

		c.JSON(200, result)
	})
	r.Run()
}

func cariNilaiTertinggi(penyakit map[string]int) string {
	poinTertinggi := 0
	penyakitTertinggi := ""

	for point, nilai := range penyakit {
		if nilai > poinTertinggi {
			poinTertinggi = nilai
			penyakitTertinggi = point
		}
	}

	return penyakitTertinggi
}

func tampilkanNilaiPenyakit(penyakit map[string]int) {
	for point, nilai := range penyakit {
		fmt.Printf("%s : %d\n", point, nilai)
	}
}
