{ pkgs }:
with pkgs;

buildGoModule rec {
  name = "paste-release";
  version = "1.0.0";

  vendorSha256 = "sha256-A4Qo3CmtMEhcITxwANAPPqngC3TkFzZjjkbnclyF63Y=";

  src = ./client;

  buildPhase = ''
    ./release.sh
  '';

  installPhase = ''
    cp -r release $out
  '';
}