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

	config := &ssh.Config{
		Ciphers: []string{"aes128-gcm@openssh.com", "chacha20-poly1305@openssh.com", "aes128-ctr", "aes192-ctr", "aes256-ctr", "aes128-cbc", "3des-cbc"},
	}

	m, _ := time.ParseDuration("3s")

	configClient := &ssh.ClientConfig{
		Config: *config,
		User:   username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         m,
	}

	return &SSHInit{
		Config: configClient,
	}

}
