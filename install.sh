#!/bin/bash

# Download latest version
echo "Downloading latest version of corsy..."
curl -s https://api.github.com/repos/themirze/corsy/releases/latest \
| grep "browser_download_url.*corsy" \
| cut -d : -f 2,3 \
| tr -d \" \
| wget -qi -

# Make it executable
chmod +x corsy

# Move to /usr/local/bin
sudo mv corsy /usr/local/bin/

echo "Corsy has been installed successfully!"
echo "You can now run 'corsy' from anywhere." 