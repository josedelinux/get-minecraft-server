# tool to get Minecraft server

### Method 1(For Linux/macOS)

```sh
#!/bin/sh
# For Linux/macOS (using curl & jq)

# Get version manifest
VERSION_MANIFEST=$(curl -s https://launchermeta.mojang.com/mc/game/version_manifest.json)

# Get latest version ID
LATEST_VERSION=$(echo "$VERSION_MANIFEST" | jq -r '.latest.release')

# Get the version URL for the latest release
VERSION_URL=$(echo "$VERSION_MANIFEST" | jq -r ".versions[] | select(.id == \"$LATEST_VERSION\") | .url")

# Get the server download URL
SERVER_URL=$(curl -s "$VERSION_URL" | jq -r '.downloads.server.url')

# Download the server JAR with versioned filename
curl -o "minecraft_server.$LATEST_VERSION.jar" "$SERVER_URL"

echo "Downloaded Minecraft server version $LATEST_VERSION as minecraft_server.$LATEST_VERSION.jar"
```

### method 2 (cross-platform)

```sh
go run .
```

## Contribution

All contributions are welcome
