package provisioner


import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
	"golang.org/x/crypto/ssh"
)

type ServerConnInfo struct {
	Server string
	Port   string
	User   string
	Key    string
}

func (c *ServerConnInfo) Socket() string {
	return fmt.Sprintf("%s:%s", c.Server, c.Port)
}

func publicKeyFile(file string) (ssh.AuthMethod, error) {
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		return nil, err
	}
	return ssh.PublicKeys(key), nil
}

func generateSession(s ServerConnInfo) (*ssh.Session, ssh.Conn, error) {
	publicKey, err := publicKeyFile(s.Key)
	if err != nil {
		return nil, nil, err
	}

	config := &ssh.ClientConfig{
		User: s.User,
		Auth: []ssh.AuthMethod{
			publicKey,
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", s.Socket(), config)
	if err != nil {
		return nil, nil, err
	}

	
	session, err := conn.NewSession()
	if err != nil {
		return nil, conn, err
	}

	return session, conn, nil
}

func SSHCommandBool(command string, sci ServerConnInfo) (bool, error) {
	session, conn, err := generateSession(sci)
	if err != nil {
		if conn != nil {
			conn.Close()
		}

		return false, err
	}
	err = session.Run(command)

	session.Close()
	conn.Close()

	if err != nil {
		return false, err
	}
	return true, nil
}

func SSHCommandString(command string, sci ServerConnInfo) (string, error) {
	session, conn, err := generateSession(sci)
	if err != nil {
		if conn != nil {
			conn.Close()
		}

		return "", err
	}

	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf

	err = session.Run(command)

	session.Close()
	conn.Close()

	if err != nil {
		return "", err
	}
	return strings.TrimSuffix(stdoutBuf.String(), "\n"), nil
}