{
  description = "Declarative local dev environment";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/23.05";
    unstable-nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    cakemix.url = "github:erikkrieg/cakemix";
  };

  outputs = { flake-utils, nixpkgs, unstable-nixpkgs, cakemix, ... }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        stable = import nixpkgs { inherit system; };
        unstable = import unstable-nixpkgs { inherit system; };
      in
      with stable; {
        devShell = mkShell {
          buildInputs = [
            cakemix.packages.${system}.default
            go
            gopls
            just
          ];
          shellHook = ''
            # Keep Go cache and module files in the project
            export GOPATH="$(pwd)/.go"
          '';
        };
      });
}
