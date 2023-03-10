{ nixpkgs, system, frontend, backend, release }:

nixpkgs.lib.nixosSystem {
  inherit system;

  modules = [
    ({ pkgs, ... }:
      {
        system.stateVersion = "22.11";
        boot.isContainer = true;

        networking.useDHCP = false;
        networking.firewall.allowedTCPPorts = [ 80 443 ];

        systemd.services.paste = {
          description = "paste";

          environment = {
            PASTE_DATA_DIR = "/data";
            PASTE_FRONTEND_DIR = "${frontend}/public";
            PASTE_RELEASE_DIR = "${release}";

            GIN_MODE = "release";
            PORT = "80";
            HTTPS_PORT = "443";
          };

          preStart = ''
            mkdir -p /data

            ${pkgs.openssl}/bin/openssl req -new -newkey rsa:4096 -nodes \
              -keyout /data/paste.key -out /data/paste.csr \
              -subj "/C=US/ST=Denial/L=Springfield/O=Dis/CN=paste"
            
            ${pkgs.openssl}/bin/openssl req -new -newkey rsa:4096 -days 365 -nodes -x509 \
              -subj "/C=US/ST=Denial/L=Springfield/O=Dis/CN=paste" \
              -keyout /data/paste.key  -out /data/paste.cert
          '';

          serviceConfig = {
            ExecStart = "${backend}/bin/paste-backend";
          };

          wantedBy = [ "default.target" ];
          enable = true;
        };
      })
  ];
}
