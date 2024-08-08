FROM ollama/ollama:latest

# Install dependencies for Intel GPU support
RUN apt-get update && \
    apt-get install -y \
    intel-gpu-tools \
    vainfo \
    libva2 \
    libva-drm2 \
    libva-x11-2 \
    i965-va-driver \
    gstreamer1.0-vaapi && \
    rm -rf /var/lib/apt/lists/*

# Add any additional setup commands if necessary
# Example: Copy application files or set environment variables

# Set the entrypoint or command if different from the base image
# CMD ["your_startup_command"]