package types

import (
	"exec"
)

type Command struct {
	Stdout, Stderr io.Writer

	cmd *exec.Cmd
}
