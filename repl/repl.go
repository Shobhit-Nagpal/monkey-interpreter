package repl

import (
    "bufio"
    "fmt"
    "io"
    "monkey-interpreter/lexer"
    "monkey-interpreter/parser"
    "monkey-interpreter/evaluator"
    "monkey-interpreter/object"
)

const PROMPT = ">>"

func Start(in io.Reader, out io.Writer) {
    scanner := bufio.NewScanner(in)
    env := object.NewEnvironment()

    for {

        fmt.Printf(PROMPT)
        scanned := scanner.Scan()

        if !scanned {
            return
        }

        line := scanner.Text()
        l := lexer.New(line)
        p := parser.New(l)

        program := p.ParseProgram()

        if len(p.Errors()) != 0 {
            printParseErrors(out, p.Errors())
            continue
        }

        evaluated := evaluator.Eval(program, env)

        if evaluated != nil {

            io.WriteString(out, evaluated.Inspect())
            io.WriteString(out, "\n")
        }


    }
}


const MONKEY_TEXT = `
 /$$      /$$ /$$   /$$ /$$   /$$ /$$   /$$
| $$$    /$$$| $$  | $$| $$$ | $$| $$  /$$/
| $$$$  /$$$$| $$  | $$| $$$$| $$| $$ /$$/ 
| $$ $$/$$ $$| $$  | $$| $$ $$ $$| $$$$$/  
| $$  $$$| $$| $$  | $$| $$  $$$$| $$  $$  
| $$\  $ | $$| $$  | $$| $$\  $$$| $$\  $$ 
| $$ \/  | $$|  $$$$$$/| $$ \  $$| $$ \  $$
|__/     |__/ \______/ |__/  \__/|__/  \__/
`

func printParseErrors(out io.Writer, errors []string) {
    io.WriteString(out, MONKEY_TEXT)
    io.WriteString(out, "Whoopsy Daisy! Ran into an error :(\n")
    io.WriteString(out, "Parser errors:\n")
    for _, msg := range errors {
        io.WriteString(out, "\t" + msg + "\n")
    }
}
