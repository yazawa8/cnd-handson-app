FROM python:3.12-slim AS lint

WORKDIR /app
COPY . /app

# black をインストール
RUN pip install --no-cache-dir black
RUN black . --exclude=venv,proto

# flake8 をインストール
RUN pip install --no-cache-dir flake8
RUN flake8 . --exclude=venv,proto --ignore=E501

FROM python:3.12-slim AS runtime

WORKDIR /app
COPY . /app

# 本番用の依存関係をインストール
RUN pip install --no-cache-dir -r requirements.txt

ENV PYTHONUNBUFFERED=1
ENV PYTHONPATH=/app

CMD ["python", "./cmd/server/main.py", "server"]
