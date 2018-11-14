package utils

import (
	"os/exec"
	"strings"
	"utils/log"
)

func Exec(cmdLines string)(msg string, err error){
	cmd := exec.Command("/bin/sh","-c",cmdLines)
	bytes, err := cmd.Output()
	if err != nil {
		return
	}
	msg = string(bytes)
	return
}

func ExecInPut(cmdLines,input string)(msg string, err error){
	cmd := exec.Command("/bin/sh","-c",cmdLines)
	cmd.Stdin =strings.NewReader(input)
	bytes, err := cmd.Output()
	if err != nil {
		log.Infof("Exec err:",cmdLines)
	}
	msg = string(bytes)
	return
}