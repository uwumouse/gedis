# Gedis - a Redis clone
This project is done within a hour just to practice Go, then I understood that I won't code in this lang.


# Building
First, clone the repo
```bash
git clone https://github.com/uwumouse/gedis.git
```
Build with make
```
make build
```
> Make sure that GoLang enviroment is set up (you have `go` executable on your PC)

# Running
`cd` into `build/` dir.

## Server
`gserv` is executable to run Gedis Server.
## Client
To run Gedis Client run `gedis` executable.

# Syntax
Gedis follows Redis syntax to update values, so:
```
# Save value
] set key value

# Read value
] get key 
```
