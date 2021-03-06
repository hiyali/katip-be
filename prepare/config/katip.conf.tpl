upstream back_end {
    ip_hash;
    server 127.0.0.1:5555;
}

server {
  listen        80;
  # listen      [::]:80 default ipv6only=on;
  server_name   {{DOMAIN}} www.{{DOMAIN}} localhost;
  # rewrite ^/(.*)$ https://{{DOMAIN}}/$1 permanent;

  root          /katip/fe;
  index         index.html;

  location /api {
    proxy_pass_header Server;
    proxy_set_header Host $http_host;
    proxy_redirect off;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Scheme $scheme;
    proxy_pass http://back_end;
  }

  location / {
    root        /katip/fe;
    expires     1d;
    add_header  Cache-Control public;
    access_log  off;
  }
}
