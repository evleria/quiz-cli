package main

import (
	"github.com/evleria/quiz-cli/pkg/cmd/root"
	"github.com/evleria/quiz-cli/pkg/cmdutils"
	"github.com/evleria/quiz-cli/pkg/iostreams"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	factory := &cmdutils.Factory{
		IOStreams: iostreams.IOStreams{
			In:     os.Stdin,
			Out:    os.Stdout,
			ErrOut: os.Stderr,
		},
	}

	cmd := root.NewRootCmd(factory)

	cmd.SetIn(factory.IOStreams.In)
	cmd.SetOut(factory.IOStreams.Out)
	cmd.SetErr(factory.IOStreams.ErrOut)

	cmdutils.CheckError(cmd.Execute())
}
