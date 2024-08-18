package errors

import (
	"errors"
	"fmt"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

var (
	ErrDBConnection        = errors.New("failed to connect to the database")
	ErrInternalServerError = errors.New("internal server error")
	ErrGroupAlreadyExists  = errors.New("group already exists")
	ErrNotFound            = errors.New("NotFound")
)

// IsDuplicateEntryError verifica se o erro é devido a uma entrada duplicada
func IsDuplicateEntryError(err error) bool {
	return errors.Is(err, gorm.ErrDuplicatedKey)
}

func IsForeignKeyViolation(err error) bool {
	if err == nil {
		return false
	}

	pgErr, ok := err.(*pq.Error)
	if ok {
		// Imprime o código do erro para depuração
		fmt.Printf("Postgres Error Code: %s\n", pgErr.Code)
		return pgErr.Code == "23503" // Código de violação de chave estrangeira
	}

	// Verifica se o erro é encapsulado (wrap) dentro de outro erro
	var pqErr *pq.Error
	if errors.As(err, &pqErr) {
		fmt.Printf("Encapsulated Postgres Error Code: %s\n", pqErr.Code)
		return pqErr.Code == "23503"
	}

	return false
}
