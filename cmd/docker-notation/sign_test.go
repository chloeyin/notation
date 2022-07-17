package main

import (
	"testing"
	"time"

	"github.com/notaryproject/notation/internal/cmd"
)

func TestSignCommand(t *testing.T) {
	command := signCommand()
	err := command.ParseFlags([]string{
		"ref",
		"--key", "kkk",
		"--key-file", "fff",
		"--cert-file", "ccc",
		"-r", "ref",
		"--origin",
		"--timestamp", "0000",
		"--expiry", "365s"})
	if err != nil {
		t.Fatalf("Parse Flag failed: %v", err)
	}
	if ref := command.Flags().Arg(0); ref != "ref" {
		t.Fatalf("Expect reference: %v, got: %v", "ref", ref)
	}
	if name, _ := command.Flags().GetString(cmd.PflagKey.Name); name != "kkk" {
		t.Fatalf("Expect %v: %v, got: %v", cmd.PflagKey.Name, "kkk", name)
	}
	if keyFile, _ := command.Flags().GetString(cmd.PflagKeyFile.Name); keyFile != "fff" {
		t.Fatalf("Expect %v: %v, got: %v", cmd.PflagKeyFile.Name, "fff", keyFile)
	}
	if certFile, _ := command.Flags().GetString(cmd.PflagCertFile.Name); certFile != "ccc" {
		t.Fatalf("Expect %v: %v, got: %v", cmd.PflagCertFile.Name, "ccc", certFile)
	}
	if origin, _ := command.Flags().GetBool("origin"); !origin {
		t.Fatalf("Expect %v: %v, got: %v", "origin", true, origin)
	}
	if tm, _ := command.Flags().GetString(cmd.PflagTimestamp.Name); tm != "0000" {
		t.Fatalf("Expect %v: %v, got: %v", cmd.PflagTimestamp.Name, "0000", tm)
	}
	if expiry, _ := command.Flags().GetDuration(cmd.PflagExpiry.Name); expiry != time.Second*365 {
		t.Fatalf("Expect %v: %v, got: %v", cmd.PflagExpiry.Name, time.Second*365, expiry)
	}
}
