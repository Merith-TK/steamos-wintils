package main

import (
	"fmt"
	"os"
	"os/exec"

	_ "embed"
)

var (
	steamdir = "C:/Program Files (x86)/Steam/"
)

//go:embed assets/shortcut.lnk
var shortcutLnk []byte

func main() {

	// if steamdir/beta file doesn't exist, create it
	if _, err := os.Stat(steamdir + "package/beta"); os.IsNotExist(err) {
		file, err := os.OpenFile(steamdir+"package/beta", os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			fmt.Println(err)
		}
		file.WriteString("steampal_stable_9a24a2bf68596b860cb6710d9ea307a76c29a04d")
		file.Close()
	}

	// stop steam.exe if it's running
	cmd := exec.Command("taskkill", "/f", "/im", "steam.exe")
	cmd.Run()
	cmd = exec.Command(steamdir+"steam.exe", "-gamepadui")
	cmd.Start()

	// write shortcut to desktop
	file, err := os.OpenFile(os.ExpandEnv("${USERPROFILE}/Desktop/DeckUI.lnk"), os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println(err)
	}
	file.Write(shortcutLnk)
	file.Close()

}
