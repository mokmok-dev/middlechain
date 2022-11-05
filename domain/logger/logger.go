//go:generate go run github.com/golang/mock/mockgen -source=$GOFILE -destination=mock/$GOFILE -package=mock

package logger

type Logger interface {
	Field(string, interface{}) Field
	Info(string, ...Field)
	Error(string, error, ...Field)
}

type (
	Field struct {
		Key       string
		Interface interface{}
	}
)
