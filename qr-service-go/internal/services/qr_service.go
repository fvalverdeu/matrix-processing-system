package services

import (
	"gonum.org/v1/gonum/mat"
)

type QRService struct {
	validator *MatrixValidator
}

func NewQRService(validator *MatrixValidator) *QRService {
	return &QRService{validator: validator}
}

func (s *QRService) Decompose(matrix [][]float64) ([][]float64, [][]float64, error) {
	if err := s.validator.Validate(matrix); err != nil {
		return nil, nil, err
	}

	rows := len(matrix)
	cols := len(matrix[0])

	data := make([]float64, 0, rows*cols)
	for _, row := range matrix {
		data = append(data, row...)
	}

	a := mat.NewDense(rows, cols, data)

	var factorization mat.QR
	factorization.Factorize(a)

	var q mat.Dense
	var r mat.Dense
	factorization.QTo(&q)
	factorization.RTo(&r)

	return denseToSlice(&q), denseToSlice(&r), nil
}

func denseToSlice(d *mat.Dense) [][]float64 {
	rows, cols := d.Dims()
	result := make([][]float64, rows)

	for i := 0; i < rows; i++ {
		result[i] = make([]float64, cols)
		for j := 0; j < cols; j++ {
			result[i][j] = d.At(i, j)
		}
	}

	return result
}
