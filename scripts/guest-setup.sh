#!/bin/sh
# Run once, manually, inside the dev VM after `zed sync` has configured the
# virtiofs share on the host side and the VM has been restarted.
set -e

MOUNT_POINT="/mnt/zedlum"
MOUNT_TAG="zedshare"

sudo mkdir -p "$MOUNT_POINT"
if ! grep -q "$MOUNT_TAG" /etc/fstab; then
	echo "$MOUNT_TAG $MOUNT_POINT virtiofs defaults 0 0" | sudo tee -a /etc/fstab
fi
sudo mount -a

mkdir -p ~/.local/share/themes
ln -sfn "$MOUNT_POINT/z-theme" ~/.local/share/themes/z-theme

sudo apt-get install -y gnome-shell-extensions
gnome-extensions enable user-theme@gnome-shell-extensions.gcampax.github.com || true

echo "done. install 'Theme and Shell Reloading' manually from extensions.gnome.org (no apt package)."
echo "reload the shell with Alt+F2, then type rt, then Enter."
