<div align="center">

![UniAPI](/web/public/logo.png)

# UniAPI

**Next-Generation LLM Gateway and AI Asset Management System**

<p align="center">
  <strong>English</strong> |
  <a href="./README.zh.md">中文</a>
</p>

<p align="center">
  <a href="https://raw.githubusercontent.com/Zer0Echo/uniapi/main/LICENSE">
    <img src="https://img.shields.io/github/license/Zer0Echo/uniapi?color=brightgreen" alt="license">
  </a>
  <a href="https://github.com/Zer0Echo/uniapi/releases/latest">
    <img src="https://img.shields.io/github/v/release/Zer0Echo/uniapi?color=brightgreen&include_prereleases" alt="release">
  </a>
  <a href="https://hub.docker.com/r/zer0echo/uniapi">
    <img src="https://img.shields.io/badge/docker-dockerHub-blue" alt="docker">
  </a>
</p>

</div>

## About

UniAPI is a fork of [New API v0.10.8](https://github.com/Calcium-Ion/new-api) (which itself builds on [One API](https://github.com/songquanpeng/one-api)). It aggregates 40+ upstream AI providers (OpenAI, Claude, Gemini, Azure, AWS Bedrock, etc.) behind a unified API with user management, billing, and rate limiting.

### What UniAPI adds on top of New API

- **Subscription & Activation Codes** — Users can activate subscriptions via redemption codes; admin can manage wallet balance and subscription balance separately
- **Admin Balance Management** — Dashboard for managing user balances with dollar-based pricing
- **Ticket System** — Built-in support ticket system with balance validity periods
- **Backend Performance Optimizations** — Async logging, dashboard query parallelization, composite DB indexes, Redis pipeline rate limiting, fine-grained in-memory rate limiter, connection pool tuning
- **Modernized UI** — Warm-toned theme with CSS variable extraction, updated dashboard layout
- **Key Management for Subscriptions** — Dedicated subscription key management interface

## Quick Start

### Docker Compose (Recommended)

```bash
git clone https://github.com/Zer0Echo/uniapi.git
cd uniapi
# Edit docker-compose.yml (change default passwords!)
docker-compose up -d
```

### Docker Run

```bash
docker run --name uniapi -d --restart always \
  -p 3000:3000 \
  -e TZ=Asia/Shanghai \
  -v ./data:/data \
  zer0echo/uniapi:latest
```

Visit `http://localhost:3000` after deployment. Default admin credentials: `root` / `123456`.

## Deployment

| Component | Requirement |
|-----------|-------------|
| **Database** | SQLite (default), MySQL >= 5.7.8, or PostgreSQL >= 9.6 |
| **Cache** | Redis (recommended) or in-memory |
| **Container** | Docker / Docker Compose |

### Key Environment Variables

| Variable | Description |
|----------|-------------|
| `SQL_DSN` | Database connection string (omit for SQLite) |
| `REDIS_CONN_STRING` | Redis connection (e.g. `redis://localhost:6379`) |
| `SESSION_SECRET` | Required for multi-node deployment |
| `CRYPTO_SECRET` | Required when sharing Redis across instances |

## Upstream Projects

| Project | License |
|---------|---------|
| [New API](https://github.com/Calcium-Ion/new-api) (v0.10.8) | AGPL-3.0 |
| [One API](https://github.com/songquanpeng/one-api) | MIT |

## License

[GNU Affero General Public License v3.0 (AGPL-3.0)](./LICENSE)

See [NOTICE](./NOTICE) for full attribution.
