package main

import (
	"fmt"
	"github.com/kataras/golog"
	"os"
	"regexp"
)

func Ipv6(parent *Command, cmd *Command) {
	cmd.AddSubCommand("firewall", "Entrer dans la configuration du port forwarding", FirewallV6)

	cmd.Read(parent)
}

func FirewallV6(parent *Command, cmd *Command) {
	cmd.AddSubCommand("add", "Ajouter une règle de port forwarding", AddPortForwardingV6)
	//cmd.AddSubCommand("delete", "Supprimer une règle de port forwarding", DeletePortForwarding)
	//cmd.AddSubCommand("list", "Afficher la liste des règles de port forwarding", ListPortForwarding)

	cmd.Read(parent)
}

func AddPortForwardingV6(parent *Command, cmd *Command) {
	ipv6, port, name, protocol, err := GetFirewallArgs()
	if err != nil {
		golog.Fatal(err)
		return
	}

	err = AddFirewallRuleV6(ipv6, port, name, protocol)
	if err != nil {
		golog.Fatal(err)
		return
	}

	golog.Info("La règle a été ajoutée avec succès!")
}

func GetFirewallArgs() (string, int, string, string, error) {
	var ipv6 string
	var portDst int
	var name string
	var protocol string

	PrintCursor("Entrer l'adresse IPv6 de la Livebox")
	_, err := fmt.Scanln(&ipv6)
	if err != nil {
		golog.Fatal(err)
		return "", 0, "", "", err
	}

	ipv6Regex := regexp.MustCompile(`^([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}$`)
	if !ipv6Regex.MatchString(ipv6) {
		golog.Error("L'adresse IPv6 n'est pas valide")
		return "", 0, "", "", err
	}

	PrintCursor("Entrer le port de destination")
	_, err = fmt.Scanln(&portDst)
	if err != nil {
		golog.Fatal(err)
		return "", 0, "", "", err
	}

	PrintCursor("Entrer le nom de la règle:")
	_, err = fmt.Scanln(&name)
	if err != nil {
		golog.Fatal(err)
		return "", 0, "", "", err
	}
	if name == "" {
		golog.Error("Le nom de la règle ne peut pas être vide")
		return "", 0, "", "", err
	}

	/**
	6,17 = TCP/UDP 6 = TCP 17 = UDP
	*/
	PrintCursor("Entrer le protocole (TCP ou UDP)")
	_, err = fmt.Scanln(&protocol)
	if err != nil {
		golog.Fatal(err)
		return "", 0, "", "", err
	}

	switch protocol {
	case "TCP":
		protocol = "6"
	case "UDP":
		protocol = "17"
	case "TCP/UDP":
		protocol = "6,17"
	default:
		golog.Error("Le protocole n'est pas valide")
		return "", 0, "", "", err
	}

	return ipv6, portDst, name, protocol, nil
}

func ExitTerminal(parent *Command, cmd *Command) {
	golog.Info("exit")
	os.Exit(0)
}
