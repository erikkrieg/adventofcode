{
  description = "Declarative local dev environment";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/23.11";
    flake-utils.url = "github:numtide/flake-utils";
    cakemix.url = "github:erikkrieg/cakemix";
    envim.url = "github:erikkrieg/envim";
    envim.inputs.nixpkgs.follows = "nixpkgs";
  };

  outputs = { flake-utils, nixpkgs, cakemix, envim, ... }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
      in
      with pkgs; {
        devShell = mkShell {
          buildInputs = [
            cakemix.packages.${system}.default
            envim.packages.${system}.default
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
