{
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs";
    flake-utils.url = "github:numtide/flake-utils";
    napalm.url = "github:nix-community/napalm";
  };

  outputs = { self, nixpkgs, flake-utils, napalm }:
    flake-utils.lib.eachDefaultSystem
      (system:
        let
          pkgs = import nixpkgs {
            inherit system;
            overlays = [
              napalm.overlay
              (final: prev: with self.packages.${system}; {
                paste-frontend-deps = paste-frontend-deps;
                paste-frontend = paste-frontend;
                paste-backend = paste-backend;
                paste-client = paste-client;
              })
            ];
          };
        in
        {
          devShell = import ./shell.nix { inherit pkgs; };

          packages = {
            paste-frontend-deps = pkgs.napalm.buildPackage ./frontend { };
            paste-frontend = import ./frontend.nix { inherit pkgs; };
            paste-backend = import ./backend.nix { inherit pkgs; };
            paste-client = import ./client.nix { inherit pkgs; };

            nixosConfigurations.container = import ./container.nix {
              inherit system;
              inherit nixpkgs;

              frontend = pkgs.paste-frontend;
              backend = pkgs.paste-backend;
            };
          };

          defaultPackage = self.packages.${system}.paste-backend;
        }
      );
}
