# AoC 2022 - ReasonML 

As I understand it, [Reason](https://reasonml.github.io/docs/en/what-and-why) is a sort of alternative syntax for OCaml that can be compiled to natively binary using OCaml's standard library and ecosystem, _or_ it can be transpiled to JavaScript (Node or browser JS).

Run solutions natively with [`esy`](https://github.com/esy/esy) by passing the date like so:

```
esy run 1
```

_This will install the dependencies required to compile the native binary._

## Project structure

Follows conventions of the [Dune build system](https://dune.readthedocs.io/en/latest/index.html):
```
# Contains the entry point and handling of the date input
bin/

# Contains solutions for each day
lib/

# Symlink to puzzle inputs for the year
input/
```
