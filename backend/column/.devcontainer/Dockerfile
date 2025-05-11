FROM mcr.microsoft.com/devcontainers/python:3.12


# [Optional] Uncomment this section to install additional OS packages.
RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends protobuf-compiler build-essential autoconf libtool pkg-config
# [Optional] Uncomment the next lines to use Python get to install anything else you need
# USER vscode
RUN python -m pip install --upgrade pip \
    && python -m pip install grpcio grpcio-tools psycopg2-binary SQLAlchemy

# USER root
# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1
