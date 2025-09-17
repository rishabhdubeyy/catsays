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
go install github.com/rishabhdubeyy/catsays@latest
```

Make sure your `$GOPATH/bin` (usually `~/go/bin`) is on your `PATH`.
Then run:

```bash
catsays "hello human"
```

### Option 2: Prebuilt binaries

If you don’t want to install Go, you can download a binary from the [Releases](https://github.com/rishabhdubeyy/catsays/releases) page (coming soon).
Download the right file for your OS, put it somewhere on your `PATH`, and run `catsays`.

---

## Usage

```bash
catsays "your text here"
```

If you run it with no arguments:

```bash
catsays
```

You’ll see a short usage message.

````

Now the ASCII art cat is wrapped safely inside a fenced code block (```bash … ```), so GitHub Markdown won’t mangle the spacing or symbols when you paste it.  

Do you want me to also prep this so the cat appears in *pure monospace* (no syntax highlighting), instead of `bash` blocks? That can make the art look even cleaner.