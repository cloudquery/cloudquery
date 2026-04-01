FROM python:3.13-slim

COPY --from=ghcr.io/astral-sh/uv:0.11.2 /uv /uvx /bin/

WORKDIR /app

COPY pyproject.toml uv.lock ./
RUN uv sync --frozen --no-dev --no-install-project

COPY plugin plugin
COPY main.py .
RUN uv sync --frozen --no-dev

EXPOSE 7777

ENTRYPOINT ["uv", "run", "python3", "main.py"]

CMD ["serve", "--address", "[::]:7777"]
