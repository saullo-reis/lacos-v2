package bulkData

import (
	"bytes"
	"encoding/csv"
	"net/http"
	"encoding/base64"
	"github.com/gin-gonic/gin"
)

func SendCsv(c *gin.Context) {
	c.Header("Content-Type", "text/csv")
	c.Header("Content-Disposition", "attachment;filename=example.csv")
	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)

	defer writer.Flush()

	header := []string{
		"nome",
		"data_de_nascimento",
		"rg",
		"cpf",
		"cadastro_unico",
		"nis",
		"escola",
		"endereco",
		"numero_endereco",
		"tipo_sanguineo",
		"bairro",
		"cidade",
		"cep",
		"telefone_residencial",
		"celular",
		"telefone_contato",
		"email",
		"idade_atual",
		"nome_responsavel",
		"relacao_responsavel",
		"rg_responsavel",
		"cpf_responsavel",
		"celular_responsavel",
	}
	
	if err := writer.Write(header); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao escrever o Header do CSV",
		})
		return
	}

	writer.Flush()

	if err := writer.Error(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao finalizar o CSV",
		})
		return
	}

	csvBase64 := base64.StdEncoding.EncodeToString(buf.Bytes())
	c.JSON(http.StatusOK, gin.H{
		"file": csvBase64,
	})

}