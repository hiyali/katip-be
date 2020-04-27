#!/bin/bash
# For LightSail Ubuntu 18
CONF_DIR=~/katip-be/config

echo "Installing mux ..."
sudo apt install -yq tmux tmuxinator

echo "Setting mux ..."
ln -fs $CONF_DIR/.tmux.conf ~/.tmux.conf
ln -fs $CONF_DIR/.tmuxinator ~/.tmuxinator
sudo ln -fs /usr/bin/tmuxinator /usr/bin/mux

echo "Opening .."
mux a
