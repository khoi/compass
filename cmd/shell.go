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
eval "$(sextant shell --type "$shell" --binding {{.Binding}})"
`

const zsh = `__sextant_chpwd() {
	[[ "$(pwd)" == "$HOME" ]] && return
    (sextant add "$(pwd)" &)
}
[[ -n "${precmd_functions[(r)__sextant_chpwd]}" ]] || {
	precmd_functions[$(($#precmd_functions+1))]=__sextant_chpwd
}
{{.Binding}}() {
	local output="$(sextant cd $@)"
	if [ -d "$output" ]; then
		builtin cd "$output"
	else
		sextant cleanup && false
	fi
}
__sextant_completion() {
	reply=(${(f)"$(sextant ls --path-only "$1")"})
}
compctl -U -K __sextant_completion {{.Binding}}
`

const bash = `__sextant_chpwd() {
	[[ "$(pwd)" == "$HOME" ]] && return
    (sextant add "$(pwd)" &)
}
grep "sextant add" <<< "$PROMPT_COMMAND" >/dev/null || {
	PROMPT_COMMAND="$PROMPT_COMMAND"$'\n''(__sextant_chpwd 2>/dev/null &);'
}
{{.Binding}}() {
	local output="$(sextant cd $@)"
	if [ -d "$output" ]; then
		builtin cd "$output"
	else
		sextant cleanup && false
	fi
}
complete -o dirnames -C 'sextant ls --path-only "${COMP_LINE/#{{.Binding}} /}"' {{.Binding}}
`

const fish = `function {{.Binding}}
	set -l output (sextant cd $argv)
	if test -d "$output" 
		cd $output
	else
		sextant cleanup; false
	end
end

function __sextant_add --on-variable PWD
    status --is-command-substitution; and return
    sextant add (pwd)
end

complete -c {{.Binding}} -x -a '(sextant ls --path-only (commandline -t))'
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

	tmpl, err := template.New("sextant").Parse(shellType)

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
