package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	database "lacosv2.com/src/database/config"
)



func Login(c *gin.Context){
	db,err := database.ConnectDB()
	if err !=nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"message": "Error ao conectar com banco de dados "+err.Error(),
		})
		return
	}
	defer db.Close()

	var body payload
	if err = c.ShouldBindJSON(&body);err !=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": 400,
			"message": "JSON Invalido",
		})
		return
	}
	passwordHashed := HasherPassword(body.Password)

	var userQueried User
	rows, err := db.Query("SELECT username, password FROM users WHERE username = $1", body.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": 500,
			"message":    "Erro ao ler dados do usuário no banco de dados",
		})
		return
	}
	defer rows.Close()

	if rows.Next(){
		var username, password string
		if err = rows.Scan(&username, &password); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status_code": 400,
				"message":    "Erro ao escanear usuário",
			})
			return
		}
		userQueried = User{
			Username: username,
			Password: password,
		}
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"status_code": 404,
			"message":    "Usuário não registrado",
		})
		return
	}

	if passwordHashed == userQueried.Password {
		token, err := GenerateJWT(userQueried.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status_code": 500,
				"message": "Error ao gerar o token: "+err.Error(),
			})
			return
		}

		c.JSON(http.StatusAccepted, gin.H{
			"status_code": 202,
			"message":    "Usuário autorizado",
			"token": token,
		})
		return
	}else{
		c.JSON(http.StatusUnauthorized, gin.H{
			"status_code": 401,
			"message": "Usuário não autorizado",
		})
		return
	}
}