
FROM node:24

ENV DEBIAN_FRONTEND=noninteractive
ENV GOLANG_VERSION=1.25.6

#-------------------------------
# Install runtime dependencies #
#-------------------------------
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    curl \
    netcat-openbsd \
    git \
    jq \
    tree \
    wget \
    gnupg \
    groff \
    less \
    unzip \
    lsb-release \
    openssh-client \
    make \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

#-----------------------------------
# Install Go with dev dependencies #
#-----------------------------------
RUN wget -q https://go.dev/dl/go${GOLANG_VERSION}.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go${GOLANG_VERSION}.linux-amd64.tar.gz && \
    rm go${GOLANG_VERSION}.linux-amd64.tar.gz

ENV PATH="/usr/local/go/bin:${PATH}"

# Install Go dev tools
RUN go install honnef.co/go/tools/cmd/staticcheck@latest && \
    go install github.com/securego/gosec/v2/cmd/gosec@latest && \
    echo 'export PATH=$PATH:$HOME/go/bin' >> ~/.bashrc

#------------------------
# Install AWS CLI & CDK #
#------------------------
RUN curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip" && \
    unzip awscliv2.zip && \
    ./aws/install && \
    npm install -g aws-cdk && \
    # verify
    aws --version && \
    cdk --version
