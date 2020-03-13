package sshinit

import (
	"time"

	"golang.org/x/crypto/ssh"
)

//SSHInit holds connectivity information
type SSHInit struct {
	Config *ssh.ClientConfig
}

// To be used only in Interactive mode
var passInteractive string

//NewSSHInit instantiate the Sshinit struct
func NewSSHInit(username string, password string, interactive bool) *SSHInit {

	config := &ssh.Config{
		Ciphers: []string{"aes128-gcm@openssh.com", "chacha20-poly1305@openssh.com",
			"aes128-ctr", "aes192-ctr", "aes256-ctr", "aes128-cbc", "3des-cbc"},
		KeyExchanges: []string{"curve25519-sha256@libssh.org", "ecdh-sha2-nistp256",
			"ecdh-sha2-nistp384", "ecdh-sha2-nistp521", "diffie-hellman-group14-sha1", "diffie-hellman-group1-sha1"},
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

	if interactive {
		passInteractive = password
		configClient.Auth = []ssh.AuthMethod{ssh.KeyboardInteractive(SshInteractive)}
	}

	return &SSHInit{
		Config: configClient,
	}

}

func SshInteractive(user, instruction string, questions []string, echos []bool) (answers []string, err error) {
	answers = make([]string, len(questions))
	// The second parameter is unused
	for n := range questions {
		answers[n] = passInteractive
	}

	return answers, nil
}
