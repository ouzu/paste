{ pkgs }:
with pkgs;

buildGoModule rec {
  name = "paste";
  version = "1.0.0";

  vendorSha256 = "sha256-pJWfzJQwLsch0CkLOhJGsRxV4lzuLIGJr6zFoqb0QOQ=";

  src = ./backend;
}