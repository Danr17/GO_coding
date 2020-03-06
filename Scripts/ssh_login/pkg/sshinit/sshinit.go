package sshinit

import (
	"time"

	"golang.org/x/crypto/ssh"
)

//SSHInit holds connectivity information
type SSHInit struct {
	Config *ssh.ClientConfig
}

//NewSSHInit instantiate the Sshinit struct
func NewSSHInit(username string, password string) *SSHInit {

	m, _ := time.ParseDuration("3s")

	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         m,
	}

	return &SSHInit{
		Config: config,
	}

}
