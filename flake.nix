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
              self.overlay.${system}
            ];
          };
        in
        {
          devShell = import ./shell.nix { inherit pkgs; };

          packages = {
            paste-frontend = import ./frontend.nix {
              inherit pkgs;
              deps = pkgs.napalm.buildPackage ./frontend { };
            };
            paste-backend = import ./backend.nix { inherit pkgs; };
            paste-client = import ./client.nix { inherit pkgs; };
            paste-release = import ./release.nix { inherit pkgs; };
          };

          overlay = (final: prev: with self.packages.${system}; {
            paste-frontend = paste-frontend;
            paste-backend = paste-backend;
            paste-client = paste-client;
            paste-release = paste-release;
          });

          defaultPackage = self.packages.${system}.paste-backend;

          nixosConfigurations.container = import ./container.nix {
            inherit system;
            inherit nixpkgs;

            frontend = pkgs.paste-frontend;
            backend = pkgs.paste-backend;
            release = pkgs.paste-release;
          };
        }
      ) // {
      hydraJobs = {
        container = self.nixosConfigurations.x86_64-linux.container;
        backend = self.packages.x86_64-linux.paste-backend;
        frontend = self.packages.x86_64-linux.paste-frontend;
        client = self.packages.x86_64-linux.paste-client;
      };
    };
}
