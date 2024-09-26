package utils

import "todo_list_api/pkg/models"

func ValidateStatus(status string) error {
	validStatuses := map[string]bool{
		models.StatusPending:    true,
		models.StatusInProgress: true,
		models.StatusCompleted:  true,
	}

	if !validStatuses[status] {
		return ErrInvalidStatus
	}

	return nil
}
