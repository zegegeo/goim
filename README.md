# GoIM
一个使用 Go 开发的智能 IM 系统，支持基础聊天、好友/群聊、高并发通信，并接入 GPT 智能助手能力。(现在处于开发第一阶段)

## 🌟 项目亮点

- 使用 **Gin + GORM** 实现后端服务
- 支持 **WebSocket 实时通信**
- 消息持久化，支持历史查询
- **Redis** 实现在线状态与缓存
- **Kafka** 支持高并发削峰填谷
- 支持集成 ChatGPT / 千帆 等 LLM 智能对话能力
- 最终形态：一个具备多插件调度的“超级助理 IM 系统”

---

## 📦 技术栈

| 模块          | 技术栈            |
|---------------|-------------------|
| 后端开发      | Go, Gin, GORM     |
| 实时通信      | WebSocket         |
| 认证授权      | JWT               |
| 数据持久化    | MySQL             |
| 状态缓存      | Redis             |
| 消息中间件    | Kafka             |
| 容器部署      | Docker, Compose   |
| 智能助手      | OpenAI / QianFan  |

---

## 🗂️ 功能分阶段规划

### ✅ 第一阶段：基础私聊功能（MVP）
- [x] 用户注册与登录（JWT）
- [x] WebSocket 实现双向通信
- [x] 点对点私聊发送
- [x] 聊天记录持久化（MySQL）
- [x] 聊天记录查询接口

### ⏭ 第二阶段：社交关系与高并发
- 好友添加/删除/列表
- 群聊创建、加入、消息广播
- 离线消息与用户在线状态缓存（Redis）
- 未读消息统计
- 使用 Kafka 异步处理消息落库

### ⏭ 第三阶段：前端界面与部署
- 使用 Vue / React 实现前端
- Docker 容器部署支持

### ⏭ 第四阶段：GPT 智能助手 + Agent 调用
- 集成 ChatGPT 聊天能力
- 插件调度：天气/翻译/搜索等 Agent
- 构建插件注册 + 工具调用系统

---

## 🚀 快速开始

```bash
git clone https://github.com/liwenze/goim.git
cd goim

# 启动服务
go run main.go