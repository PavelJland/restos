package restos

import (
	"errors"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

func Find_and_replace(path string, text string, replace string) error {
	input, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatalln(err)
		return errors.New("Coundn't read a file!")
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, text) {
			lines[i] = replace
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(path, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
		return errors.New("Coundn't write a file!")
	}
	return nil
}

func SetWindows(args []string) error {
	text := "GRUB_DEFAULT=0"
	replace := "GRUB_DEFAULT=2"
	err := Find_and_replace("/etc/default/grub", text, replace)
	if err != nil {
		log.Fatalln(err)
		return errors.New("Coundn't read a file!")
	}
	text = "GRUB_TIMEOUT=10"
	replace = "GRUB_TIMEOUT=1"
	err = Find_and_replace("/etc/default/grub", text, replace)
	if err != nil {
		log.Fatalln(err)
		return errors.New("Coundn't write a file!")
	}
	ccmd := exec.Command("/bin/sh", "-c", "sudo update-grub2")
	ccmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	if (len(args) > 0) && (args[0] == "r") {
		Reboot()
	}
	return nil
}

func Reboot() {

	cmd2 := exec.Command("/bin/sh", "-c", "sudo reboot")
	if err := cmd2.Start(); err != nil {
		log.Fatal(err)
	}
}

func SetLinux(args []string) error {
	// fmt.Printf("%v", args)
	text := "GRUB_DEFAULT=2"
	replace := "GRUB_DEFAULT=0"
	err := Find_and_replace("/etc/default/grub", text, replace)
	if err != nil {
		log.Fatalln(err)
		return errors.New("Coundn't read a file!")
	}
	text = "GRUB_TIMEOUT=1"
	replace = "GRUB_TIMEOUT=10"
	err = Find_and_replace("/etc/default/grub", text, replace)
	if err != nil {
		log.Fatalln(err)
		return errors.New("Coundn't write a file!")
	}
	ccmd := exec.Command("/bin/sh", "-c", "sudo update-grub2")
	ccmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	// ch := make(chan error)

	// go func() {
	// 	ch <- ccmd.Start()
	// }()
	// errUpdate := <-ch
	// if errUpdate != nil {
	// 	log.Fatal(errUpdate)
	// }
	if (len(args) > 0) && (args[0] == "r") {
		Reboot()
	}
	return nil
}
