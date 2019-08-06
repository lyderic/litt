package cmd

func oldmain() {
	/*
		flag.StringVar(&jsonPath, "f", "./litt.json", "configuration `file`")
		flag.StringVar(&reference, "m", "1", "create `montage`")
		flag.BoolVar(&nocontent, "b", false, "don't build content")
		flag.BoolVar(&tag, "t", false, "tag final PDF with montage name and timestamp")
		flag.BoolVar(&verbose, "v", false, "be verbose")
		flag.Usage = usage
		flag.Parse()
		if _, err := os.Stat(jsonPath); os.IsNotExist(err) {
			usage()
			log.Fatal(err)
		}
		basedir = getParent(jsonPath)
		configuration.load()
		if len(flag.Args()) == 0 {
			usage()
			fmt.Println("Please provide a command!")
			return
		}
		action := flag.Args()[0]
		found := false
		var command Command
		for _, current := range commands {
			if current.Name == action || current.Letter == action {
				found = true
				command = current
			}
		}
		if !found {
			fmt.Printf("%s: invalid command.\n", action)
			usage()
			return
		}
		command.Action()
	*/
}

func usage() {
	/*
		fmt.Printf("\nlitt v.%s - (c) Lyderic Landry, London 2018\n", VERSION)
		fmt.Println("Usage: litt [<option>] <command>")
		fmt.Println("\n Commands:\n")
		for _, command := range commands {
			fmt.Println(command)
		}
		fmt.Println("\n Options:\n")
		flag.PrintDefaults()
		if len(configuration.Montages) > 0 {
			fmt.Println("\n Montages:\n")
			list()
		}
		fmt.Println()
	*/
}
