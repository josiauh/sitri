package sitrilib

import (
	"errors"
	"os"
	"strings"
)

func GetSitriInfo() ([]string, error) {
	sitriInfoB, err := os.ReadFile(".sitri/projectInfo")
	if err != nil {
		return nil, errors.New("Could not find project info. Probably not a project?")
	}
	sitriInfoS := string(sitriInfoB)
	sitriInfo := strings.Split(sitriInfoS, "\n")
	return sitriInfo, nil
}

func WriteSitriInfo(info []string) error {
	sitriInfo, err := os.Open(".sitri/projectInfo")

	if err != nil {
		return errors.New("Could not write project info. Probably not a project?")
	}
	sitriInfo.Write([]byte("[sitriInfo v1]\n" + strings.Join(info, "\n")))
	return nil
}
