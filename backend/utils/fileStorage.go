package utils

import (
	"dot_conf/constants"
	"dot_conf/dto"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func Upload(r *http.Request) (string, error) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		log.Error("Unable to parse the form due to: ", err.Error())
		return "", err
	}

	file, handler, err := r.FormFile("document")
	if err != nil {
		log.Error("Unable to retrieve the file from the form due to: ", err.Error())
		return "", err
	}
	defer file.Close()

	var rc dto.RegisterCompany

	err = json.NewDecoder(r.Body).Decode(&rc)
	if err != nil {
		log.Error("Unable to decode the request body due to: ", err.Error())
		return "", err
	}

	savePath := fmt.Sprintf("%s/%s/%s/%s", constants.Root, constants.Verification, rc.CompanyName, handler.Filename)
	outFile, err := os.Create(savePath)
	if err != nil {
		log.Error("Unable to save the file due to: ", err.Error())
		return "", err
	}
	defer outFile.Close()

	return savePath, nil
}
