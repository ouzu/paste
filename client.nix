{ pkgs }:
with pkgs;

buildGoModule rec {
  name = "pst";
  version = "1.0.0";

  vendorSha256 = "sha256-A4Qo3CmtMEhcITxwANAPPqngC3TkFzZjjkbnclyF63Y=";

  src = ./client;
}