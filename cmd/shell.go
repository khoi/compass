package cmd

import (
	"fmt"

	"html/template"

	"bytes"

	"github.com/spf13/cobra"
)

const sh = `#!bin/sh
if [ -n "${BASH}" ]; then
	shell="bash"
elif [ -n "${ZSH_NAME}" ]; then
	shell="zsh"
elif [ -n "${__fish_datadir}" ]; then
    shell="fish"
else
	shell=$(echo "${SHELL}" | awk -F/ '{ print $NF }')
fi
if [ "${shell}" = "sh" ]; then
	return 0
fi
eval "$(compass shell --type "$shell" --bind-to {{.Binding}})"
`

const zsh = `__compass_chpwd() {
	[[ "$(pwd)" == "$HOME" ]] && return
    (compass add "$(pwd)" &)
}
[[ -n "${precmd_functions[(r)__compass_chpwd]}" ]] || {
	precmd_functions[$(($#precmd_functions+1))]=__compass_chpwd
}
{{.Binding}}() {
	local output="$(compass cd $@)"
	if [ -d "$output" ]; then
		builtin cd "$output"
	else
		compass cleanup && false
	fi
}
__compass_completion() {
	reply=(${(f)"$(compass ls --path-only "$1")"})
}
compctl -U -K __compass_completion {{.Binding}}
`

const bash = `__compass_chpwd() {
	[[ "$(pwd)" == "$HOME" ]] && return
    (compass add "$(pwd)" &)
}
grep "compass add" <<< "$PROMPT_COMMAND" >/dev/null || {
	PROMPT_COMMAND="$PROMPT_COMMAND"$'\n''(__compass_chpwd 2>/dev/null &);'
}
{{.Binding}}() {
	local output="$(compass cd $@)"
	if [ -d "$output" ]; then
		builtin cd "$output"
	else
		compass cleanup && false
	fi
}
complete -o dirnames -C 'compass ls --path-only "${COMP_LINE/#{{.Binding}} /}"' {{.Binding}}
`

const fish = `function {{.Binding}}
	set -l output (compass cd $argv)
	if test -d "$output" 
		cd $output
	else
		compass cleanup; false
	end
end

function __compass_add --on-variable PWD
    status --is-command-substitution; and return
    compass add (pwd)
end

complete -c {{.Binding}} -x -a '(compass ls --path-only (commandline -t))'
`

func scriptForShell(shell string, keyBinding string) string {
	var b struct {
		Binding string
	}
	b.Binding = keyBinding
	shellType := func() string {
		switch shell {
		case "bash":
			return bash
		case "zsh":
			return zsh
		case "fish":
			return fish
		default:
			return sh
		}
	}()

	tmpl, err := template.New("compass").Parse(shellType)

	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, b)

	if err != nil {
		panic(err)
	}

	return buf.String()
}

var shellType string
var keyBinding string

var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "Prints out the shell integration scripts.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(scriptForShell(shellType, keyBinding))
	},
}

func init() {
	rootCmd.AddCommand(shellCmd)
	shellCmd.Flags().StringVarP(&shellType, "type", "t", "sh", "Type of the shell (bash|zsh|fish)")
	shellCmd.Flags().StringVarP(&keyBinding, "bind-to", "b", "s", "Key binding (default is `s`)")
}
