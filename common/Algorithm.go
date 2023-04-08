package common

import (
	"log"
	"os/exec"
)

func AlgorithmExec() (int, error) {
	cmd := exec.Command("cpp", "main.cpp", "-o") // 传入参数补充
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	return 0, err
}
