#!/bin/bash

# Define the version tag
VERSION="v0.1.3-beta-1"

# Navigate to the current directory
cd "$(pwd)" || exit

# Tag the version
git tag "$VERSION"

# Push the tag to the remote repository
git push origin "$VERSION"

# Check if the tag was pushed successfully
if [ $? -eq 0 ]; then
    echo "Successfully pushed tag $VERSION"
else
    echo "Failed to push tag $VERSION"
    exit 1
fi

# Set GOPROXY and list the module to verify the tag
GOPROXY=proxy.golang.org go list -m github.com/ayushs-2k4/go-consul-service-manager@"$VERSION"

# Check if the go list command was successful
if [ $? -eq 0 ]; then
    echo "Module version $VERSION is available"
else
    echo "Module version $VERSION not found"
    exit 1
fi
