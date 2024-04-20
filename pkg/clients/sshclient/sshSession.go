package sshclient

import (
	"fmt"
	"os"

	"github.com/bramvdbogaerde/go-scp"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

type SshSession struct {
	client  *ssh.Client
	session *ssh.Session
}

func NewSshSession(hostNameAndPort string, userName string, password string, timeout int) *SshSession {
	var hostKey ssh.PublicKey
	config := &ssh.ClientConfig{
		User: userName,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.FixedHostKey(hostKey),
	}
	client, err := ssh.Dial("tcp", hostNameAndPort, config)
	if err != nil {
		return nil
	}
	session, err := client.NewSession()
	if err != nil {
		return nil
	}
	return &SshSession{
		session: session,
		client:  client,
	}
}

func (thisSession *SshSession) ExecuteCommand(command string) (string, error) {
	var output []byte
	var err error
	if thisSession.session != nil {
		output, err = thisSession.session.Output(command)
	}
	return string(output), err
}

func (thisSession *SshSession) Close() {

}

func (thisSession *SshSession) UploadFile(localFilePath string, remoteFilePath string) {
	sftpClient, err := sftp.NewClient(thisSession.client)
	if err != nil {
		return
	}
	defer sftpClient.Close()
	srcFile, err := os.Open(localFilePath)
	if err != nil {
		return
	}
	defer srcFile.Close()
}

func (thisSession *SshSession) DownloadFile(remoteFilePath string, localFilePath string) {
	client, err := scp.NewClientBySSH(thisSession.client)
	if err != nil {
		fmt.Println("Error creating new SSH session from existing connection", err)
	} else {
		localFile, err := os.Open(localFilePath)
		if err != nil {
			fmt.Println("Error opening local file", err)
			return
		}
		client.CopyFromRemote(nil, localFile, remoteFilePath)
	}
}
