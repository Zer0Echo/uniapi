<div align="center">

![UniAPI](/web/public/logo.png)

# UniAPI

**新一代大模型网关与 AI 资产管理系统**

<p align="center">
  <a href="./README.md">English</a> |
  <strong>中文</strong>
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

## 简介

UniAPI 基于 [New API v0.10.8](https://github.com/Calcium-Ion/new-api)（New API 基于 [One API](https://github.com/songquanpeng/one-api)）二次开发。聚合 40+ 上游 AI 服务商（OpenAI、Claude、Gemini、Azure、AWS Bedrock 等），提供统一 API 接口，内置用户管理、计费和限流。

### UniAPI 在 New API 基础上新增的功能

- **订阅与兑换码** — 用户通过兑换码激活订阅；管理员可分别管理钱包余额和订阅余额
- **管理员余额管理** — 以美元计价的用户余额管理面板
- **工单系统** — 内置工单系统，支持余额有效期
- **后端性能优化** — 异步日志、仪表盘查询并行化、复合数据库索引、Redis Pipeline 限流、细粒度内存限流器、连接池调优
- **UI 现代化** — 暖色调主题、CSS 变量提取、仪表盘布局更新
- **订阅密钥管理** — 专用的订阅密钥管理界面

## 快速开始

### Docker Compose（推荐）

```bash
git clone https://github.com/Zer0Echo/uniapi.git
cd uniapi
# 编辑 docker-compose.yml（务必修改默认密码！）
docker-compose up -d
```

### Docker 命令

```bash
docker run --name uniapi -d --restart always \
  -p 3000:3000 \
  -e TZ=Asia/Shanghai \
  -v ./data:/data \
  zer0echo/uniapi:latest
```

部署完成后访问 `http://localhost:3000`。默认管理员账号：`root` / `123456`。

## 部署要求

| 组件 | 要求 |
|------|------|
| **数据库** | SQLite（默认）、MySQL >= 5.7.8 或 PostgreSQL >= 9.6 |
| **缓存** | Redis（推荐）或内存缓存 |
| **容器** | Docker / Docker Compose |

### 关键环境变量

| 变量 | 说明 |
|------|------|
| `SQL_DSN` | 数据库连接字符串（不设置则使用 SQLite） |
| `REDIS_CONN_STRING` | Redis 连接地址（如 `redis://localhost:6379`） |
| `SESSION_SECRET` | 多机部署时必须设置 |
| `CRYPTO_SECRET` | 共享 Redis 时必须设置 |

## 上游项目

| 项目 | 许可证 |
|------|--------|
| [New API](https://github.com/Calcium-Ion/new-api)（v0.10.8） | AGPL-3.0 |
| [One API](https://github.com/songquanpeng/one-api) | MIT |

## 许可证

[GNU Affero 通用公共许可证 v3.0 (AGPL-3.0)](./LICENSE)

完整归属信息见 [NOTICE](./NOTICE)。
