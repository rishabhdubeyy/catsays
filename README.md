# catsays 🐱

A tiny Go CLI that makes your cat talk back.  
Think [cowsay](https://en.wikipedia.org/wiki/Cowsay), but with sharper claws.  

```bash
$ catsays "feed me"
feed me 
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
````

---

## Installation

### Option 1: Go users

If you have Go installed, grab the latest version directly:

```bash
go install github.com/rishabhdubeyy/catsays/cmd/catsays@latest
```

Make sure your `$GOPATH/bin` (usually `~/go/bin`) is on your `PATH`.
Then run:

```bash
catsays "hello human"
```

## Usage

```bash
catsays "your text here"
```

If you run it with no arguments:

```bash
catsays
```

You’ll see a short usage message.

## Update

```bash
go clean -modcache
go install github.com/rishabhdubeyy/catsays/cmd/catsays@latest
```