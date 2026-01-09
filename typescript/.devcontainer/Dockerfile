FROM oven/bun:1

# 開発に便利なツールやGitをインストール (Debianベースの場合)
RUN apt-get update && apt-get install -y \
    git \
    curl \
    && rm -rf /var/lib/apt/lists/*

# COPY .bashrc /root/.bashrc

RUN mkdir -p ~/.config/git
RUN echo "node_modules" >> ~/.config/git/ignore


