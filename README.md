# 🚗 NAVO - API de Controle de Gastos para Motoristas de Aplicativo

**Navo** é uma API RESTful construída em Go com [Gin](https://github.com/gin-gonic/gin), pensada para facilitar o controle de gastos de motoristas de aplicativo como Uber, 99, InDrive, entre outros.

---

## 📌 Objetivo

Oferecer uma plataforma robusta e simples para que motoristas registrem seus gastos diários, categorizem despesas por veículo, viagem ou abastecimento, e acompanhem relatórios mensais detalhados.

---

## 🧱 Tecnologias

- **Go (Golang)** + Gin (framework web)
- **PostgreSQL** como banco de dados
- **GORM** (ORM para Go)
- **JWT** para autenticação
- **Arquitetura limpa**, com separação em camadas
- **Docker** (em breve)

---

## 📦 Estrutura do Projeto

```bash
navo/
├── api/        # Backend em Go
│   ├── cmd/
│   ├── config/
│   ├── domain/
│   ├── dto/
│   ├── infrastructure/
│   ├── interfaces/
│   ├── internal/
│   ├── pkg/
│   └── go.mod
└── app/        # (Futuro) App mobile em Flutter
