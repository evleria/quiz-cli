package cmdutils

import (
	"github.com/evleria/quiz-cli/pkg/config"
	"github.com/evleria/quiz-cli/pkg/iostreams"
)

type Factory struct {
	IOStreams iostreams.IOStreams
	Config *config.Config
}

