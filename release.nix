{ pkgs }:
with pkgs;

buildGoModule rec {
  name = "paste-release";
  version = "1.0.0";

  vendorSha256 = "sha256-qT5xLA2ZbYZEzajLarmaZrk23yxonmChC1uxH7hWwJA=";

  src = ./client;

  buildPhase = ''
    ./release.sh
  '';

  installPhase = ''
    cp -r release $out
  '';
}