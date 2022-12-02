let
  rust_overlay = import (builtins.fetchTarball https://github.com/oxalica/rust-overlay/archive/master.tar.gz);
  pkgs = import <nixpkgs> { overlays = [ rust_overlay ]; };
  unstable = import (fetchTarball https://nixos.org/channels/nixos-unstable/nixexprs.tar.xz) {};
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
      ## Nimble fails does not appear to work when installed with nix-shell, so
      ## relying on a non-nix installation for the time being.
      # nim
      nimlsp
    ];
  }

