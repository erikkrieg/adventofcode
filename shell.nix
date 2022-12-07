let
  rust_overlay = import (builtins.fetchTarball https://github.com/oxalica/rust-overlay/archive/master.tar.gz);
  pkgs = import <nixpkgs> { overlays = [ rust_overlay ]; };
  rust = pkgs.rust-bin.selectLatestNightlyWith (toolchain: toolchain.default.override {
    extensions = [ "rust-src" ];
  });
in
  with pkgs;
  mkShell {
    nativeBuildInputs = [
      buildPackages.just
      # Go
      buildPackages.go_1_18
      buildPackages.gopls
      # Rust
      # rust-bin.stable."1.64.0".default
      rust
      rust-analyzer
      # Nim
      ## Nimble does not appear to work when installed with nix-shell, so I am
      ## relying on a non-nix installation for now.
      # nim
      nimlsp
    ];
  }

