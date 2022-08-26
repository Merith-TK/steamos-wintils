package main

import (
	"flag"
	"fmt"

	"steamos-wintils/config"
	"steamos-wintils/util"
)

func main() {
	err := config.SetupConfig()
	if err != nil {
		fmt.Println(err)
	}
	flagged := false
	// register -c as bool flag
	var c bool
	var l bool
	flag.BoolVar(&c, "c", false, "")
	flag.BoolVar(&l, "l", false, "")

	flag.Parse()
	if c && !l {
		flagged = true
		fmt.Println(setBranch("get", ""))
		return
	}
	if l && !flagged {
		fmt.Println("rel\nrc\nbeta\nbc\nmain")
		return
	}
	if flag.Arg(0) == "" && !flagged {
		fmt.Println("Usage: steamos-select-branch <-c|-l|rel|rc|beta|bc|main>")
		return
	} else if flag.Arg(0) != "" && !flagged {
		setBranch("set", flag.Arg(0))
		return
	}

}

func setBranch(mode string, branch string) string {
	conf := config.Config
	if mode == "set" {
		conf.SelectBranch.Branch = branch
		config.WriteConfig(conf)
		return conf.SelectBranch.Branch
	} else if mode == "get" || mode == "" {
		util.DebugPrint("Config:", conf)
		return conf.SelectBranch.Branch
	}
	return ""
}
