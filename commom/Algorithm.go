package commom

import (
	"log"
	"os/exec"
)

func AlgorithmExec() {
	cmd := exec.Command("cpp", "main.cpp", "-o") // 传入参数补充
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
