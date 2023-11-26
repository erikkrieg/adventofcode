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
    just
    # Go
    go
    gopls
    # Rust
    # rust-bin.stable."1.64.0".default
    rust
    rust-analyzer
    # Nim
    ## Nimble does not appear to work when installed with nix-shell, so I am
    ## relying on a non-nix installation for now.
    # nim
    nimlsp
    # ReasonML
    ## Node is used to get npm, which installs `esy`, the package manager
    ## I'm using for Reason.
    nodejs_20
  ];
}
