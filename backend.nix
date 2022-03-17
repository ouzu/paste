{ pkgs }:
with pkgs;

buildGoModule rec {
  name = "paste";
  version = "1.0.0";

  vendorSha256 = "sha256-EtrmTocz1+7kPlSf8f3NKhndvOsMHzjvriKOuqjx314=";

  src = ./backend;
}