package utils

import (
	"encoding/csv"
	"fmt"
	"io"
)

// ReadCSVFile 从io.Reader中读取CSV文件并返回所有行
func ReadCSVFile(reader io.Reader) ([][]string, error) {
	csvReader := csv.NewReader(reader)
	csvReader.Comma = ','

	lines, err := csvReader.ReadAll()
	if err != nil {
		msg := fmt.Sprintf("读取CSV文件数据错误 error %s", err.Error())
		Intolog(msg)
		return nil, err
	}

	return lines, nil
}
