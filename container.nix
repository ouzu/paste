{ nixpkgs, system, frontend, backend }:

nixpkgs.lib.nixosSystem {
  inherit system;

  modules = [
    ({ pkgs, ... }:
      let
        httpPort = 443;
      in
      {
        boot.isContainer = true;

        networking.useDHCP = false;
        networking.firewall.allowedTCPPorts = [ httpPort ];

        systemd.services.paste = {
          description = "paste";

          environment = {
            PASTE_DATA_DIR = "/data";
            PASTE_ADDR = "0.0.0.0:${toString httpPort}";
            PASTE_FRONTEND_DIR = "${frontend}/public";
            PASTE_SSL = "yes";
          };

          preStart = ''
            mkdir /data

            ${pkgs.openssl}/bin/openssl req -new -newkey rsa:4096 -nodes \
              -keyout /data/paste.key -out /data/paste.csr \
              -subj "/C=US/ST=Denial/L=Springfield/O=Dis/CN=paste"
            
            ${pkgs.openssl}/bin/openssl req -new -newkey rsa:4096 -days 365 -nodes -x509 \
              -subj "/C=US/ST=Denial/L=Springfield/O=Dis/CN=paste" \
              -keyout /data/paste.key  -out /data/paste.cert
          '';

          serviceConfig = {
            ExecStart = "${backend}/bin/paste";
          };

          wantedBy = [ "default.target" ];
          enable = true;
        };
      })
  ];
}
