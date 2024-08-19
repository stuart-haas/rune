package services

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"

	"golang.org/x/crypto/ssh"
)

type SSHClientConfig struct {
	User string
	Hostname string
}

type SSHClient struct {
	Conf SSHClientConfig
}

type SSHConnection struct {
	*ssh.Client
	*ssh.Session
	StdOut io.Reader
	StdErr io.Reader
}

func NewSSHClient (conf SSHClientConfig) (*SSHClient, error) {
	return &SSHClient{
		Conf: conf,
	}, nil
}

func (s *SSHClient) Run() error {
	conn, err := s.Connect()
	if err != nil {
		return err
	}

	defer s.Close(*conn)
	
	if err := conn.Session.Run("ls"); err != nil {
		return err
	}

	output, err := io.ReadAll(conn.StdOut)
	if err != nil {
		return err
	}

	errorOutput, err := io.ReadAll(conn.StdErr)
	if err != nil {
		return err
	}

	fmt.Printf("Output:\n%s\n", string(output))
  fmt.Printf("Error Output:\n%s\n", string(errorOutput))

	return nil
}

func (s *SSHClient) Start() error {
	conn, err := s.Connect()
	if err != nil {
		return err
	}

	defer s.Close(*conn)

	if err := conn.Session.Start("while true; do sudo hostname -f; sleep 2; done"); err != nil {
		return err
	}

	go func() {
		scanner := bufio.NewScanner(conn.StdOut)
		for scanner.Scan() {
				fmt.Println(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
				log.Printf("Error reading stdout: %v", err)
		}
	}()

	go func() {
		scanner := bufio.NewScanner(conn.StdErr)
		for scanner.Scan() {
				fmt.Println("STDERR:", scanner.Text())
		}
		if err := scanner.Err(); err != nil {
				log.Printf("Error reading stderr: %v", err)
		}
	}()

	if err := conn.Session.Wait(); err != nil {
			return err
	}

	return nil
}

func (s *SSHClient) Connect() (*SSHConnection, error) {
	conf, err := s.GetConfig()
	if err != nil {
		return nil, err
	}

	client, err := ssh.Dial("tcp", net.JoinHostPort(s.Conf.Hostname, "22"), conf)
	if err != nil {
		return nil, err
	}

	session, err := client.NewSession()
	if err != nil {
		return nil, err
	}

	stdout, err := session.StdoutPipe()
	if err != nil {
		return nil, err
	}

	stderr, err := session.StderrPipe()
	if err != nil {
		return nil, err
	}

	return &SSHConnection{
		Client: client,
		Session: session,
		StdOut: stdout,
		StdErr: stderr,
	}, nil
}

func (s *SSHClient) Close(conn SSHConnection) {
	conn.Client.Close()
	conn.Session.Close()
}

func (s *SSHClient) GetPrivateKey() (*ssh.Signer, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	key, err := os.ReadFile(filepath.Join(homeDir, ".ssh", "id_rsa"))
	if err != nil {
		return nil, err
	}
	privateKey, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return nil, err
	}
	return &privateKey, nil
}

func (s *SSHClient) GetConfig() (*ssh.ClientConfig, error) {
	privateKey, err := s.GetPrivateKey()
	if err != nil {
		return nil, err
	}
	return &ssh.ClientConfig{
		User: s.Conf.User,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(*privateKey),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}, nil
}