NGINX_CONF_DIR=~/katip-be/config/nginx
NGINX_ROOT_DIR=/etc/nginx

echo "Installing nginx ..."
sudo apt install nginx -yq
sudo service nginx stop

echo "Setting nginx configs ..."
sudo mv $NGINX_ROOT_DIR/nginx.conf $NGINX_ROOT_DIR/nginx.conf.back
sudo ln -fs $NGINX_CONF_DIR/hiyali.org $NGINX_ROOT_DIR/hiyali.org
sudo ln -fs $NGINX_CONF_DIR/nginx.conf $NGINX_ROOT_DIR/nginx.conf

echo "Start nginx service ..."
sudo service nginx start
