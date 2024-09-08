package utils

func ValidateStatus(status string) error {
	validStatuses := map[string]bool{
		StatusPending:    true,
		StatusInProgress: true,
		StatusCompleted:  true,
	}

	if !validStatuses[status] {
		return ErrInvalidStatus
	}

	return nil
}
