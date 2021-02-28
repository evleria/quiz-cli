package main

import (
	"github.com/evleria/quiz-cli/pkg/cmd/root"
	"github.com/evleria/quiz-cli/pkg/cmdutils"
)

func main() {
	cmd := root.NewRootCmd()
	cmdutils.CheckError(cmd.Execute())
}
