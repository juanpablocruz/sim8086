# sim8086

sim8086 is a prototype simulator for the 8086 family of Intel processors.
It handles binary decoding into 8086 instructions. And (__WIP__) emulates
a 8086 machine that can then execute it.

## Installation

To install `sim8086` clone this repository and build it with go.

```bash
git clone https://github.com/juanpablocruz/sim8086.git
cd sim8086
go run cmd/main.go binarycode
```

There are some sample binary files in `part1/*`
