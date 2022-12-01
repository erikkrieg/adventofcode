let
  rust_overlay = import (builtins.fetchTarball https://github.com/oxalica/rust-overlay/archive/master.tar.gz);
  pkgs = import <nixpkgs> { overlays = [ rust_overlay ]; };
in
  with pkgs;
  mkShell {
    nativeBuildInputs = [
      buildPackages.just
      # Go
      buildPackages.go_1_18
      buildPackages.gopls
      # Rust
      rust-bin.stable."1.64.0".default
      rust-analyzer
      # Nim
      nim
      nimlsp
    ];
  }

