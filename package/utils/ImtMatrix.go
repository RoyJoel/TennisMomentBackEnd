package utils

import (
	"bytes"
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
)

type IntMatrix [][][]int

func (m *IntMatrix) Scan(value interface{}) error {
	// 将读取到的字符串转换为 []byte 类型
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("IntMatrix.Scan: expected []byte got %T", value)
	}

	// 用 | 分割字符串，并将每个元素解析为 int 类型，存储到二维切片中
	str := string(bytes)
	rows := strings.Split(str, ";")
	matrix := make([][][]int, len(rows))

	for i, row := range rows {
		xs := strings.Split(row, ",")
		matrix[i] = make([][]int, len(xs))
		for j, x := range xs {
			cols := strings.Split(x, ".")
			matrix[i][j] = make([]int, len(cols))
			for k, col := range cols {
				num, err := strconv.Atoi(col)
				if err != nil {
					return fmt.Errorf("IntMatrix.Scan: failed to unmarshal data: %v", err)
				}
				matrix[i][j][k] = num
			}
		}

	}

	// 将解析得到的二维切片赋值给目标变量
	*m = matrix
	return nil
}

func (m IntMatrix) Value() (driver.Value, error) {
	var buf bytes.Buffer

	for i, row := range m {
		for j, xs := range row {
			for k, x := range xs {
				if k != 0 {
					buf.WriteRune('.')
				}
				fmt.Fprintf(&buf, "%d", x)
			}
			if j != len(row)-1 {
				buf.WriteRune(',')
			}
		}
		if i != len(m)-1 {
			buf.WriteRune(';')
		}
	}

	return buf.Bytes(), nil
}
