package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const sh = `#!bin/sh
if [ -n "${BASH}" ]; then
	shell="bash"
elif [ -n "${ZSH_NAME}" ]; then
	shell="zsh"
else
	shell=$(echo "${SHELL}" | awk -F/ '{ print $NF }')
fi
if [ "${shell}" = "sh" ]; then
	return 0
fi
eval "$(sextant shell --type "$shell")"
`

const zsh = `__sextant_chpwd() {
	[[ "$(pwd)" == "$HOME" ]] && return
    (sextant add "$(pwd)" &)
}
[[ -n "${precmd_functions[(r)__sextant_chpwd]}" ]] || {
	precmd_functions[$(($#precmd_functions+1))]=__sextant_chpwd
}
s() {
	local output="$(sextant cd $@)"
	if [ -d "$output" ]; then
		test -d "$output" && builtin cd "$output"
	else
		sextant cleanup && false
	fi
}
__sextant_completion() {
	reply=(${(f)"$(sextant ls --path-only "$1")"})
}
compctl -U -K __sextant_completion s
`

const bash = `__sextant_chpwd() {
	[[ "$(pwd)" == "$HOME" ]] && return
    sextant add "$(pwd)"
}
grep "sextant add" <<< "$PROMPT_COMMAND" >/dev/null || {
	PROMPT_COMMAND="$PROMPT_COMMAND"$'\n''(__sextant_chpwd 2>/dev/null &);'
}
s() {
	local output="$(sextant cd $@)"
	if [ -d "$output" ]; then
		test -d "$output" && builtin cd "$output"
	else
		sextant cleanup && false
	fi
}
complete -o dirnames -C 'sextant ls --path-only "${COMP_LINE/#s /}"' s
`

func scriptForShell(shell string) string {
	switch shell {
	case "sh":
		return sh
	case "bash":
		return bash
	case "zsh":
		return zsh
	default:
		return fmt.Sprintf("echo Sextant: We don't support %s shell yet :(", shell)
	}
}

var shellType string

var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "Prints out the shell integration scripts.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(scriptForShell(shellType))
	},
}

func init() {
	rootCmd.AddCommand(shellCmd)
	shellCmd.Flags().StringVarP(&shellType, "type", "t", "sh", "Type of the shell (bash|zsh)")
}
