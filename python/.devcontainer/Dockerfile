# ベースは好みで 3.12 / 3.11 等に変更してください
FROM python:3.13-slim

# 必要最低限のOSパッケージ
# - git: 依存に VCS 参照がある場合に必要
# - curl, ca-certificates: uv のインストールに使用
RUN apt-get update && apt-get install -y --no-install-recommends \
    git \
    curl \
    ca-certificates \
  && rm -rf /var/lib/apt/lists/*

# uv をインストール（公式インストーラ）
# インストール先は ~/.local/bin に入るので PATH を通します
RUN curl -LsSf https://astral.sh/uv/install.sh | sh
ENV PATH="/root/.local/bin:${PATH}"

# uv の一般的運用: プロジェクト直下に .venv を作る
# これにより VS Code 側の interpreter も合わせやすい
ENV UV_VENV_IN_PROJECT=1

# ワークディレクトリ
WORKDIR /workspaces/app

# 依存解決をレイヤーキャッシュしやすいように先にコピー
# uv.lock がある場合はそれもコピー（再現性が上がる）
COPY pyproject.toml ./
COPY uv.lock* ./

# 依存関係を同期（lockがあれば frozen で厳密に。無ければ通常 sync）
RUN uv sync --frozen || uv sync

# 残りのソースをコピー
COPY . .

# 既定コマンド（任意）
CMD ["bash"]
