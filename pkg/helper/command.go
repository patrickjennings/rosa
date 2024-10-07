package helper

import (
	"strings"

	"github.com/alessio/shellescape"
)

// Builder defines the structure of the builder.
type Builder struct {
	parts []string
}

// Initializes and returns a new Builder.
func NewBuilder(baseCommand string) *Builder {
	return &Builder{
		parts: []string{baseCommand},
	}
}

// adds a subcommand to the command string.
func (cb *Builder) AddSubCommand(subCommand string) *Builder {
	cb.parts = append(cb.parts, subCommand)
	return cb
}

// adds an argument to the command string.
func (cb *Builder) AddArgument(arg string) *Builder {
	cb.parts = append(cb.parts, shellescape.Quote(arg))
	return cb
}

// adds a flag and its value to the command string.
// i.e. --flag value
func (cb *Builder) AddFlag(flag, value string) *Builder {
	combined := "--" + flag + " " + shellescape.Quote(value)
	cb.parts = append(cb.parts, combined)
	return cb
}

// return the final command string.
func (cb *Builder) Build() string {
	return strings.Join(cb.parts, " ")
}
