package main

import (
	"bufio"
	"github.com/antonlindstrom/sysbuildworker/config"
	"github.com/benmanns/goworker"
	"io"
	"log"
	"os/exec"
)

func CreateVM(queue string, args ...interface{}) error {

	hostname, ok := args[0].(string)

	if !ok {
		return nil
	}

	log.Printf("Starting build of %s (template: %s)\n", hostname, config.Template())
	log.Printf("Saving image in %s\n", config.StoragePath())
	cmd := exec.Command(config.CreateVMCommand(), hostname, config.Template(), config.StoragePath())

	stdout, _ := cmd.StdoutPipe()
	process := bufio.NewReader(stdout)
	err := cmd.Start()

	if err != nil {
		log.Fatal(err)
	}

	for {
		line, _, err := process.ReadLine()

		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
				return err
			}
			break
		}

		log.Printf("%s - %s: %s\n", queue, config.CreateVMCommand(), line)
	}

	return cmd.Wait()
}

func init() {
	goworker.Register("CreateVM", CreateVM)
}

func main() {
	if err := goworker.Work(); err != nil {
		log.Println("Error:", err)
	}
}
