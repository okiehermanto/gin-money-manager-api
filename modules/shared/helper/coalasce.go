package helper

func CoalesceString(v, fallback string) string {
	if v == "" {
		return fallback
	}
	return v
}

func CoalesceInt(v, fallback int) int {
	if v == 0 {
		return fallback
	}
	return v
}
