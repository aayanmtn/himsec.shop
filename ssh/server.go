
package ssh

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/wish"
	"github.com/gliderlabs/ssh"
)

func StartSSHServer() {
	s, err := wish.NewServer(
		wish.WithAddress(":2222"),
		wish.WithHostKeyPath(".ssh/term_info_ed25519"),
		wish.WithMiddleware(
			wish.WithBubbleTea(),
		),
	)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Starting SSH server on port 2222...")
	if err := s.ListenAndServe(); err != ssh.ErrServerClosed {
		log.Fatalln(err)
	}
}
