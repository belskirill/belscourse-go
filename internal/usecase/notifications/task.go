package notifications

type EmailTask struct {
	To         string
	Template   string
	Code       string
	TTLMinutes int
}
