package runner

import (
	"ga/generator"
	"ga/parser"
	"log"
	"os"
	"os/exec"
)

func Run(name string, params map[string]string) (bool, error) {
	cmd, err := parser.Find(name, params)
	if err != nil {
		return false, err
	}

	filename := generator.GenerateShFile(cmd.CMDs)
	_, e := exec.Command("./" + filename).Output()
	if e != nil {
		log.Fatalln(e)
	}
	defer os.Remove(filename)

	return true, nil
}
