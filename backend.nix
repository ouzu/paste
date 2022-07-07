{ pkgs }:
with pkgs;

buildGoModule rec {
  name = "paste";
  version = "1.0.0";

  vendorSha256 = "sha256-zsHUV1HQzzUaWUEVW+HoukoIbnUy44FgjleUO3puIoc=";

  src = ./backend;
}