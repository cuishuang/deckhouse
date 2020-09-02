package app

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/crypto/ssh/terminal"
	"gopkg.in/alecthomas/kingpin.v2"
)

const DefaultSshAgentPrivateKeys = "~/.ssh/id_rsa"

var (
	SshAgentPrivateKeys = make([]string, 0)
	SshPrivateKeys      = make([]string, 0)
	SshBastionHost      = ""
	SshBastionPort      = ""
	SshBastionUser      = os.Getenv("USER")
	SshUser             = os.Getenv("USER")
	SshHost             = ""
	SshPort             = ""
	SshExtraArgs        = ""

	AskBecomePass = false
	BecomePass    = ""
)

func DefineSshFlags(cmd *kingpin.CmdClause) {
	cmd.Flag("ssh-agent-private-keys", "Paths to private keys. Those keys will be used to connect to servers and to the bastion. Can be specified multiple times (default: '~/.ssh/id_rsa')").
		StringsVar(&SshAgentPrivateKeys)
	cmd.Flag("ssh-bastion-host", "Jumper (bastion) host to connect to servers (will be used both by terraform and ansible). Only IPs or hostnames are supported, name from ssh-config will not work.").
		StringVar(&SshBastionHost)
	cmd.Flag("ssh-bastion-port", "SSH destination port").
		StringVar(&SshBastionPort)
	cmd.Flag("ssh-bastion-user", "User to authenticate under when connecting to bastion (default: $USER)").
		Default(SshBastionUser).
		StringVar(&SshBastionUser)
	cmd.Flag("ssh-user", "User to authenticate under (default: $USER)").
		Default(SshUser).
		StringVar(&SshUser)
	cmd.Flag("ssh-host", "SSH destination host").
		StringVar(&SshHost)
	cmd.Flag("ssh-port", "SSH destination port").
		StringVar(&SshPort)
	cmd.Flag("ssh-extra-args", "extra args for ssh commands (-vvv)").
		StringVar(&SshExtraArgs)

	cmd.PreAction(func(c *kingpin.ParseContext) (err error) {
		if len(SshAgentPrivateKeys) == 0 {
			SshAgentPrivateKeys = append(SshAgentPrivateKeys, DefaultSshAgentPrivateKeys)
		}
		SshPrivateKeys, err = ParseSSHPrivateKeyPaths(SshAgentPrivateKeys)
		if err != nil {
			return fmt.Errorf("ssh private keys: %v", err)
		}
		return nil
	})
}

func ParseSSHPrivateKeyPaths(pathSets []string) ([]string, error) {
	res := make([]string, 0)
	if len(pathSets) == 0 || (len(pathSets) == 1 && pathSets[0] == "") {
		return res, nil
	}

	for _, pathSet := range pathSets {
		keys := strings.Split(pathSet, ",")
		for _, k := range keys {
			if strings.HasPrefix(k, "~") {
				home := os.Getenv("HOME")
				if home == "" {
					return nil, fmt.Errorf("HOME is not defined for key '%s'", k)
				}
				k = strings.Replace(k, "~", home, 1)
			}

			keyPath, err := filepath.Abs(k)
			if err != nil {
				return nil, fmt.Errorf("get absolute path for '%s': %v", k, err)
			}
			res = append(res, keyPath)
		}
	}
	return res, nil
}

func DefineBecomeFlags(cmd *kingpin.CmdClause) {
	// Ansible compatible
	cmd.Flag("ask-become-pass", "Ask for sudo password before the installation process.").
		Short('K').
		BoolVar(&AskBecomePass)
}

func AskBecomePassword() (err error) {
	if !AskBecomePass {
		return nil
	}
	var data []byte
	if !terminal.IsTerminal(int(os.Stdin.Fd())) {
		return fmt.Errorf("stdin is not a terminal, error reading password")
	}
	fmt.Print("[sudo] Password: ")
	data, err = terminal.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()
	if err != nil {
		return fmt.Errorf("read password: %v", err)
	}
	BecomePass = string(data)
	return nil
}
