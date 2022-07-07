{ pkgs }:
with pkgs;

buildGoModule rec {
  name = "pst";
  version = "1.0.0";

  vendorSha256 = "sha256-qT5xLA2ZbYZEzajLarmaZrk23yxonmChC1uxH7hWwJA=";

  src = ./client;
}