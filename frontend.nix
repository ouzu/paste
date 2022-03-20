{ pkgs, deps }:
with pkgs;

stdenv.mkDerivation {
  name = "paste-frontend";

  src = ./frontend;

  buildInputs = [
    nodejs
    nodePackages.rollup
  ];

  buildPhase = ''
    ln -s ${deps}/_napalm-install/node_modules ./node_modules

    rollup -c
  '';

  installPhase = ''
    mkdir -p $out/
    cp -r public $out/
    
    mkdir -p $out/public/build/files
    cp node_modules/@fontsource/fira-code/files/* $out/public/build/files
    cp node_modules/@fontsource/fira-sans/files/* $out/public/build/files
  '';
}
