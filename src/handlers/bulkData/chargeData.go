package bulkData

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ChargeData(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Falha ao carregar o arquivo: " + err.Error(),
		})
		return
	}

	defer file.Close()

	out, err := os.Create(header.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file: "+ err.Error(),})
		return
	}
	defer out.Close()

	_, err = out.ReadFrom(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file: "+ err.Error(),})
		return
	}

	in, err := os.Open(header.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": "Falha ao abrir CSV: "+ err.Error(),
		})
		return
	}

	reader := csv.NewReader(in)
	records, err := reader.ReadAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Falha ao ler o arquivo: "+ err.Error(),
		})
		return
	}

	apiURL := "http://localhost:8080/persons/create"
	errors := []string{}
	for i, record := range records {
		if i == 0 {
			continue
		}
		convertedNumber, err := strconv.Atoi(record[17])
		if err != nil {
			errors = append(errors, fmt.Sprintf("Linha %d: Erro na conversão da idade"))
			continue
		}
		payload := map[string]interface{}{
			"name":           record[0],
			"birth_date":     record[1],
			"rg":             record[2],
			"cpf":            record[3],
			"cad_unico":      record[4],
			"nis":            record[5],
			"school":         record[6],
			"address":        record[7],
			"address_number": record[8],
			"blood_type":     record[9],
			"neighborhood":   record[10],
			"city":           record[11],
			"cep":            record[12],
			"home_phone":     record[13],
			"cell_phone":     record[14],
			"contact_phone":  record[15],
			"email":          record[16],
			"current_age":    convertedNumber,
			"responsible_person": map[string]interface{}{
				"name":         record[18],
				"relationship": record[19],
				"rg":           record[20],
				"cpf":          record[21],
				"cell_phone":   record[22],
			},
		}
		fmt.Println(payload)
		jsonData, err := json.Marshal(payload)
		if err != nil {
			errors = append(errors, fmt.Sprintf("Linha %d: Falha ao criar JSON - %v", i+1, err))
			continue
		}
		fmt.Println("JSON")
		fmt.Println(bytes.NewBuffer(jsonData))

		req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
		if err != nil {
			errors = append(errors, fmt.Sprintf("Linha %d: Falha ao criar requisição - %v", i+1, err))
			continue
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", authHeader)
		fmt.Println(authHeader)
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			errors = append(errors, fmt.Sprintf("Linha %d: Erro ao chamar API - %v", i+1, err))
			continue
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		if resp.StatusCode != http.StatusOK {
			errors = append(errors, fmt.Sprintf("Linha %d: Erro da API - %s", i+1, string(body)))
			continue
		}
	}

	if len(errors) > 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "Processamento concluído com erros",
			"errors":  errors,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Todos os dados foram enviados com sucesso",
	})
}
