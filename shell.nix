{ pkgs ? import <nixpkgs> { } }:
with pkgs;
mkShell {
  name = "paste-shell";
  buildInputs = [
    # go
    go_1_17
    go-rice

    # nodejs
    nodejs
    nodePackages.npm
    nodePackages.node2nix
    nodePackages.rollup

    # other
    xh
    nixpkgs-fmt
  ];
}
