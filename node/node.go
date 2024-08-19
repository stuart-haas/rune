package node

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

type Host struct {
	Hostname string
	User string
	Password string
}

func GetPrivateKey() (*ssh.Signer, error) {
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

func ExecuteSSH(host Host) error {
	privateKey, err := GetPrivateKey()
	if err != nil {
		return err
	}

	conf := &ssh.ClientConfig{
		User: host.User,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(*privateKey),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", net.JoinHostPort(host.Hostname, "22"), conf)
	if err != nil {
		return err
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	stdout, err := session.StdoutPipe()
	if err != nil {
		return err
	}

	stderr, err := session.StderrPipe()
	if err != nil {
		return err
	}

	if err := session.Run("ls"); err != nil {
		return err
	}

	output, err := io.ReadAll(stdout)
	if err != nil {
		return err
	}

	errorOutput, err := io.ReadAll(stderr)
	if err != nil {
		return err
	}

	fmt.Printf("Output:\n%s\n", string(output))
  fmt.Printf("Error Output:\n%s\n", string(errorOutput))

	return nil
}

func StreamSSH(host Host) error {
	privateKey, err := GetPrivateKey()
	if err != nil {
		return err
	}

	conf := &ssh.ClientConfig{
		User: host.User,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(*privateKey),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", net.JoinHostPort(host.Hostname, "22"), conf)
	if err != nil {
		return err
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	stdout, err := session.StdoutPipe()
	if err != nil {
		return err
	}

	stderr, err := session.StderrPipe()
	if err != nil {
		return err
	}

	if err := session.Start("while true; do sudo echo hello world; sleep 2; done"); err != nil {
		return err
	}

	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
				fmt.Println(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
				log.Printf("Error reading stdout: %v", err)
		}
	}()

	go func() {
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
				fmt.Println("STDERR:", scanner.Text())
		}
		if err := scanner.Err(); err != nil {
				log.Printf("Error reading stderr: %v", err)
		}
	}()

	if err := session.Wait(); err != nil {
			log.Fatalf("Command finished with error: %v", err)
	}

	return nil
}