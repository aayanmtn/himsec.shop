
package ssh

import (
	"fmt"
	"log"

	"github.com/charmbracelet/wish"
	bm "github.com/charmbracelet/wish/bubbletea"
	"github.com/gliderlabs/ssh"
	tea "github.com/charmbracelet/bubbletea"
)

func makeTeaHandler() wish.Handler {
	return bm.Handler(func() *tea.Program {
		return tea.NewProgram(initialModel())
	})
}

func StartSSHServer() {
	s := &ssh.Server{
		Addr:             ":2222",
		Handler:          makeTeaHandler(),
		PublicKeyHandler: ssh.PublicKeyAuth(func(ctx ssh.Context, key ssh.PublicKey) bool {
			return true // Allow all keys for demo
		}),
	}

	fmt.Println("Starting SSH server on port 2222...")
	if err := s.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
