{ pkgs ? import <nixpkgs> { } }:
with pkgs;
mkShell {
  name = "waschhausapp-shell";
  buildInputs = [
    # go
    go_1_17
    go-rice

    # nodejs
    nodejs
    nodePackages.npm

    # debugging
    gdb
    sqlite

    # other
    gnumake
    xh
  ];
}
