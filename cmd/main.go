package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"

	"github.com/juanpablocruz/sim8086/pkg/instruction"
	"github.com/juanpablocruz/sim8086/pkg/lexer"
	"github.com/juanpablocruz/sim8086/pkg/options"
	"github.com/juanpablocruz/sim8086/pkg/reader"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("exactly one argument required")
		fmt.Printf("Usage: sim8086 asmfile")
		os.Exit(1)
	}

	// execFlag := flag.Bool("exec", false, "-exec to interprete the code")
	showClocksFlag := flag.Bool("showclocks", false, "-showclocks to show cycles for each instruction")

	flags := uint32(0)
	if *showClocksFlag {
		flags |= options.SimFlag_ShowClocks
	}

	fileName := os.Args[1]

	rd, err := reader.New(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", rd.Dump())
	defer rd.Close()

	allInstr := []instruction.Instruction{}

	l := lexer.New(rd)
	for {
		in := l.NextInstruction()
		if in.Op == 0 {
			break
		}
		allInstr = append(allInstr, in)
	}

	fmt.Printf("; %s dissasembly:\n", fileName)
	fmt.Println("bits 16")
	fmt.Println("")

	var out bytes.Buffer
	for _, instr := range allInstr {
		out.WriteString(instr.String())
		out.WriteString("\n")
	}

	fmt.Println(out.String())
	// DisAsm8086(uint32(len(inpt)), MainMemory, flags, timing)
}
