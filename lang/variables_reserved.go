package lang

import (
	"os"

	"github.com/lmorg/murex/utils/envvars"
	"github.com/lmorg/murex/utils/path"
)

func getVarSelf(p *Process) interface{} {
	return map[string]interface{}{
		"Parent":      int(p.Scope.Parent.Id),
		"Scope":       int(p.Scope.Id),
		"TTY":         p.Scope.Stdout.IsTTY(),
		"Method":      p.Scope.IsMethod,
		"Interactive": Interactive,
		"Not":         p.Scope.IsNot,
		"Background":  p.Scope.Background.Get(),
		"Module":      p.Scope.FileRef.Source.Module,
	}
}

func getVarArgs(p *Process) interface{} {
	return append([]string{p.Scope.Name.String()}, p.Scope.Parameters.StringArray()...)
}

func getVarMurexExeValue() (interface{}, error) {
	pwd, err := os.Executable()
	if err != nil {
		return nil, err
	}

	return path.Unmarshal([]byte(pwd))
}

func getHostname() string {
	name, _ := os.Hostname()
	return name
}

func getPwdValue() (interface{}, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	return path.Unmarshal([]byte(pwd))
}

func getEnvVarValue() interface{} {
	v := make(map[string]interface{})
	envvars.All(v)
	return v
}

func getGlobalValues() interface{} {
	m := make(map[string]interface{})

	GlobalVariables.mutex.Lock()
	for name, v := range GlobalVariables.vars {
		m[name] = v.Value
	}
	GlobalVariables.mutex.Unlock()

	return m
}
