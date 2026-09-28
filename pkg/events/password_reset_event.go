package events

// PasswordResetEvent representa a mensagem enviada à fila SQS contendo a nova senha gerada para envio ao e-mail do usuário.
type PasswordResetEvent struct {
	EventID           string `json:"event_id"`
	Email             string `json:"email"`
	UserName          string `json:"user_name"`
	TemporaryPassword string `json:"temporary_password"`
}
