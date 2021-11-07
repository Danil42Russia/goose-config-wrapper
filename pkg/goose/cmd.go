package goose

import (
	"fmt"
	"strings"

	"github.com/Danil42Russia/goose-config-wrapper/pkg/config"
)

type Cmd struct {
	cfg *config.Config
}

func New(cfg *config.Config) *Cmd {
	return &Cmd{
		cfg: cfg,
	}
}

func (c Cmd) genOptions() string {
	o := c.cfg.Goose.Options
	d := c.cfg.DefaultGooseOptions()

	var options []string

	if o.Dir != d.Dir {
		options = append(options, fmt.Sprintf("-dir=%s", o.Dir))
	}
	if o.Table != d.Table {
		options = append(options, fmt.Sprintf("-table=%s", o.Table))
	}
	if o.CertFile != d.CertFile {
		options = append(options, fmt.Sprintf("-certfile=%s", o.CertFile))
	}
	if o.Verbose != d.Verbose {
		options = append(options, fmt.Sprintf("-v=%t", o.Verbose))
	}
	if o.Sequential != d.Sequential {
		options = append(options, fmt.Sprintf("-s=%t", o.Sequential))
	}
	if o.AllowMissing != d.AllowMissing {
		options = append(options, fmt.Sprintf("-allow-missing=%t", o.AllowMissing))
	}

	return strings.Join(options, " ")
}

func (c Cmd) GenRunCmd(command string) string {
	g := c.cfg.Goose
	gooseCommand := fmt.Sprintf("goose %s %s %q %s", c.genOptions(), g.Driver, g.DBString, command)

	return gooseCommand
}
