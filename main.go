package main

import (
	"flag"
	"fmt"
	"github.com/enescakir/emoji"
	"github.com/kianooshaz/repogen/parser"
	"os"
	"path"
	"strings"
)

func main() {
	sourcePtr := flag.String("source", "", "a path of source sql file")

	flag.Parse()

	if *sourcePtr == "" {
		fmt.Printf("%v️ Error: The --source flag with the path of the SQL queries file is required!\n", emoji.PoutingFace)
		return
	}

	if ext := path.Ext(*sourcePtr); ext != ".sql" {
		fmt.Printf("%v️ Error: Source file should have a .sql extension!\n", emoji.FaceWithSteamFromNose)
		return
	}

	sourceFile, err := os.ReadFile(*sourcePtr)
	if err != nil {
		fmt.Printf("%v️ Error: Source file '%s' not found!\n", emoji.FaceWithSymbolsOnMouth, *sourcePtr)
		return
	}

	fmt.Printf("%s File read successfully.\n", emoji.StarStruck)

	source := string(sourceFile)
	source = strings.NewReplacer(`\n`, "\n", `\t`, "\t", `\r`, "\r").Replace(source)

	// TODO handle multi database
	p := parser.New(source)
	if err := p.Parse(); err != nil {
		fmt.Printf("%v️ Error: %v\n", emoji.FrowningFace, err)
		return
	}

	fmt.Printf("%s Code generation successful. Enjoy it!\n", emoji.SmilingFaceWithSunglasses)
}
