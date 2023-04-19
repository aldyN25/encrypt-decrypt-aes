package main

import (
	aes "encrypt-decrypt-aes/aes"
	"fmt"
	"log"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type ValidateNameStruct struct {
	Name string `validate:"name"`
}

// custom validator
func ValidateName(fl validator.FieldLevel) bool {
	regexExp := regexp.MustCompile(`^[a-z A-Z,.\-']+$`)
	checking := regexExp.MatchString(fl.Field().String())

	log.Println("String Name : ", fl.Field().String())
	log.Println("Result Name : ", checking)

	// if !checking {
	// 	return false
	// }

	return checking
}

func main() {
	var c *gin.Context
	var name string
	nameEncode := "Jhon Lenon"

	// Encode Name
	encodeName := aes.GcmEncode(nameEncode)

	//Validasi Name and Decode name
	status, errMsg, errMsgLocal, decData := ValidatorName(encodeName)
	if status != 200 {
		c.JSON(400, gin.H{
			"status":  status,
			"message": errMsg,
			"desc":    errMsgLocal,
		})
		return
	}
	name = decData
	log.Println(status)
	log.Println(name)
}

func ValidatorName(name string) (status int, errorMessage string, errorMessageLocal string, decData string) {
	status = 200

	validate := validator.New()

	if name == "" {
		status = 400
		errorMessage = "Maaf, Parameter Nama tidak boleh kosong"
		errorMessageLocal = "Parameter name kosong"

		return status, errorMessage, errorMessageLocal, decData

	}

	// validasi gcm decode
	decData, decStatus := aes.GcmDecode(name)
	fmt.Println("DEC STATUS :", decStatus)
	if decStatus != 200 {
		status = 400
		errorMessage = "Maaf, Parameter Nama tidak sesuai"
		errorMessageLocal = "Parameter name tidak sesuai"

		return status, errorMessage, errorMessageLocal, decData
	}

	//helper validator
	validate.RegisterValidation("name", ValidateName)
	str := ValidateNameStruct{Name: decData}

	//validasi numeric
	errsValid := validate.Struct(str)
	if errsValid != nil {
		status = 400
		errorMessage = "Nama hanya boleh menggunakan huruf"
		errorMessageLocal = "Format Parameter name tidak sesuai"

		return status, errorMessage, errorMessageLocal, decData

	}

	return status, errorMessage, errorMessageLocal, decData
}
