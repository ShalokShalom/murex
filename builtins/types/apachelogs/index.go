package apachelogs

import (
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types/define"
	"github.com/lmorg/murex/utils"
)

func index(p *proc.Process, params []string) error {
	jInterface, err := unmarshal(p)
	if err != nil {
		return err
	}

	marshaller := func(iface interface{}) ([]byte, error) {
		return utils.JsonMarshal(iface, p.Stdout.IsTTY())
	}

	return define.IndexTemplateObject(p, params, &jInterface, marshaller)
}