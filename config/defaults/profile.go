package defaults

func init() {
	murexProfile = append(murexProfile, `
func h {
    # Output the murex history log in a human readable format
    history -> foreach { -> [ Index Block ] -> sprintf: "%6s => %s\n" }
}

func aliases {
	# Output the aliases in human readable format
	runtime: --aliases -> formap name alias {
		$name -> sprintf: "%10s => ${out @alias}\n"
	}
}

autocomplete set cd { [{
    "IncDirs": true
}] }

autocomplete set mkdir { [{
    "IncDirs": true
}] }

autocomplete set rmdir { [{
    "IncDirs": true
}] }

autocomplete set exec { [
    {
        "IncFiles": true,
        "IncDirs": true,
        "IncExePath": true
    },
    {
        "NestedCommand": true
    }
] }

autocomplete set format { [{
    "Dynamic": ({ runtime: --marshallers })
}] }

autocomplete set cast { [{
    "Dynamic": ({ runtime: --unmarshallers })
}] }

autocomplete set tout { [{
    "Dynamic": ({ runtime: --marshallers })
}] }

autocomplete set swivel-datatype { [{
    "Dynamic": ({ runtime: --marshallers })
}] }

autocomplete set config { [{
    "Flags": [ "get", "set" ],
    "FlagValues": {
        "get": [
            { "Dynamic": ({ config: -> formap k v { out $k } -> sort }) },
            { "Dynamic": ({ config: -> [ ${params->[2]} ] -> formap k v { out $k } -> sort }) }
        ],
        "set": [
            { "Dynamic": ({ config: -> formap k v { out $k } -> sort }) },
            { "Dynamic": ({ config: -> [ ${params->[2]} ] -> formap k v { out $k } -> sort }) },
            { "Dynamic":
				({
					params -> set params
					switch {
						case { = `+"`${ config: -> [ $params[2] ] -> [ $params[3] ] -> [ Data-Type ] }`==`bool`"+` } {
							ja [true,false]
						}

						case { config: -> [ $params[2] ] -> [ $params[3] ] -> [ Options ] } {
							config: -> [ $params[2] ] -> [ $params[3] ] -> [ Options ]
						}

						catch {
							out ${config -> [ $params[2] ] -> [ $params[3] ] -> [ Default ]}
						}
					}
				})
			}
        ]
    }
}] }

autocomplete set event { [
    {
        "Dynamic": "{ runtime: --events -> formap k v { out $k } }"
    }
] }

autocomplete set !event { [
    {
        "Dynamic": "{ runtime: --events -> formap k v { out $k } -> sort }"
    },
    {
        "Dynamic": "{ runtime: --events -> [ ${ params->[1] } ] -> formap k v { out $k } -> sort }"
    }
] }

autocomplete set autocomplete { [{
    "Flags" : [ "get", "set" ]
}] }

autocomplete set git { [{
    "Flags": [ "clone", "init", "add", "mv", "reset", "rm", "bisect", "grep", "log", "show", "status", "branch", "checkout", "commit", "diff", "merge", "rebase", "tag", "fetch", "pull", "push", "stash" ],
    "FlagValues": {
        "init": [{ "Flags": ["--bare"] }],
        "add": [{ "IncFiles": true }],
        "mv": [{ "IncFiles": true }],
        "rm": [{ "IncFiles": true }],
        "checkout": [{
            "Dynamic": ({ git branch -> [ :0 ] -> grep -v * }),
            "Flags": [ "-b" ]
        }]
    }
}] }

autocomplete set docker { [
    {
        "Flags": [ "config", "container", "image", "network", "node", "plugin", "secret", "service", "stack", "swarm", "system", "volume", "attach", "build", "commit", "cp", "create", "diff", "events", "exec", "export", "history", "images", "info", "inspect", "kill", "load", "login", "logout", "logs", "pause", "port", "ps", "pull", "push", "rename", "restart", "rm", "rmi", "run", "save", "search", "start", "stats", "stop", "tag", "top", "unpause", "update", "version", "wait" ]
    },
    {
        "Flags": [ "-t" ],
        "Optional": true,
        "AllowMultiple": true,
        "AnyValue": true
    },
    {
        "IncFiles": true
    }
] }

autocomplete set terraform { [{
    "Flags": ["apply","console","destroy","env","fmt","get","graph","import","init","output","plan","providers","push","refresh","show","taint","untaint","validate","version","workspace"],
    "FlagValues": {
        "workspace": [
            {
                "Flags": [ "new", "delete", "select", "list", "show" ]
            }
        ]
    }
}] }

autocomplete set gopass { [
    {
        "Flags": ["--yes","--clip","-c","--help","-h","--version","-v"],
        "AllowMultiple": true,
        "Dynamic": "{ exec: @{params} --generate-bash-completion }",
        "AutoBranch": true
    }
] }

autocomplete set debug { [{
    "Flags": ["on", "off"]
}] }

func progress {
    # Pulls the read progress of a Linux pid via /proc/$pid/fdinfo (only runs on Linux)

    if { = `+"`${os}`==`linux`"+` } then {
        params -> [ 1 ] -> set pid
        
        g <!null> /proc/$pid/fd/* -> regexp <!null> (f#/proc/[0-9]+/fd/([0-9]+)) -> foreach <!null> fd {
            trypipe <!null> {
                open /proc/$pid/fdinfo/$fd -> cast yaml -> [ pos ] -> set pos
                readlink: /proc/$pid/fd/$fd -> set file
                du -b $file -> [ :0 ] -> set int size
                if { = size > 0 } then {
                    = ($pos/$size)*100 -> set int percent
                    out "$percent% ($pos/$size) $file"
                }
            }
        }
    }
}

autocomplete set progress {
    [{
        "DynamicDesc": ({
            ps -A -o pid,cmd --no-headers -> set ps
            map { $ps[:0] } { $ps -> regexp 'f/^[ 0-9]+ (.*)$' }
        }),
        "ListView": true
    }]
}
`)
}
