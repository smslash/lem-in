package parse

import (
	"os"

	"git/ssengerb/lem-in/internal/handle"
	"git/ssengerb/lem-in/internal/models"
)

func OpenFile(filePath string, g *models.Graph) {
	file, err := os.Open(filePath)
	if err != nil {
		handle.Error("ERROR: can not open " + filePath)
	}
	defer file.Close()

	parseData(file, g)
}
