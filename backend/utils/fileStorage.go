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

	var rc *dto.RegisterCompany

	reqBody := r.FormValue("metadata")
	err = json.Unmarshal([]byte(reqBody), &rc)
	if err != nil {
		log.Error("Unable to decode the request body due to: ", err.Error())
		return "", err
	}

	directory := fmt.Sprintf("%s/%s/%s", constants.Root, constants.Verification, rc.CompanyName)
	if err = os.MkdirAll(directory, os.ModePerm); err != nil {
		log.Error("Unable to create the directory due to: ", err.Error())
		return "", err
	}

	savePath := fmt.Sprintf("%s/%s", directory, handler.Filename)
	outFile, err := os.Create(savePath)
	if err != nil {
		log.Error("Unable to save the file due to: ", err.Error())
		return "", err
	}
	defer outFile.Close()

	return savePath, nil
}
