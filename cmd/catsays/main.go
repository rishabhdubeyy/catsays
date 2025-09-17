package main

import (
	"fmt"
	"os"
	"strings"
)

func renderCat(input []string) string {
	cat := ` 
                        \'*-.                   
                        )  _'-.                
                        .  : '. .               
                        : _   '  \              
                        ; *' _.   '*-._         
                        '-.-'          '-.      
                          ;       '       '.    
                          :.       .        \   
                          . \  .   :   .-'   .  
                          '  '+.;  ;  '      :  
                          :  '  |    ;       ;-.
                          ; '   : :'-:     _.'* ;
                      .*' /  .*' ; .*'- +'  '*'
                      '*-*   '*-*  '*-*'        
	`
	return strings.Join(input, " ") + cat
}

func main() {
	p := fmt.Print
	n := fmt.Println
	input := os.Args[1:]
	if len(input) < 1 {
		n("Usage: catsays [your text here]")
		return
	}
	p(renderCat(input))
}
