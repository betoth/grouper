package errors

import (
	"fmt"
	"grouper/config/logger"
)

// HandleServiceError centraliza o tratamento de erros nos serviços
func HandleServiceError(err error, serviceName, operation string) error {
	logger.Error(fmt.Sprintf("Error in %s service during %s", serviceName, operation), err)

	// Verificar se o erro é um erro de negócio e adicionar contexto
	switch {
	case IsDuplicateEntryError(err):
		return fmt.Errorf("%s service: %s failed: %w", serviceName, operation, ErrGroupAlreadyExists)
	case IsForeignKeyViolation(err):
		return fmt.Errorf("%s service: %s failed: %w", serviceName, operation, ErrNotFound)
	case err == ErrNotFound:
		return fmt.Errorf("%s service: %s failed: %w", serviceName, operation, ErrNotFound)
	}

	// Para todos os outros erros, tratá-los como erros de aplicação
	return ErrInternalServerError
}
