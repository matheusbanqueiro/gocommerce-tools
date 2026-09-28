package channels

import (
	"github.com/matheusbanqueiro/gocommerce-tools/pkg/events"
)

// PasswordResetChan é o canal assíncrono interno em memória (buffer 100) para tarefas de recuperação de senha.
var PasswordResetChan = make(chan events.PasswordResetEvent, 100)

// DispatchPasswordReset despacha o evento para o canal sem travar a requisição HTTP.
func DispatchPasswordReset(evt events.PasswordResetEvent) {
	select {
	case PasswordResetChan <- evt:
	default:
		// Se o buffer estiver cheio, despacha em goroutine para evitar travamento da API
		go func() {
			PasswordResetChan <- evt
		}()
	}
}
