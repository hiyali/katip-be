upstream back_end {
    ip_hash;
    server 127.0.0.1:5555;
}

server {
  listen        443 ssl http2;
  # listen      [::]:443 default ipv6only=on;
  server_name   {{DOMAIN}} www.{{DOMAIN}};

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

  ssl on;
  ssl_certificate     /etc/letsencrypt/live/{{DOMAIN}}/fullchain.pem;
  ssl_certificate_key /etc/letsencrypt/live/{{DOMAIN}}/privkey.pem;

  ssl_session_timeout 5m;
  ssl_protocols SSLv3 TLSv1 TLSv1.1 TLSv1.2;
  ssl_ciphers "HIGH:!aNULL:!MD5 or HIGH:!aNULL:!MD5:!3DES";
  ssl_prefer_server_ciphers on;
}
