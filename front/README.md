# OAuth Frontend Application

Aplicação frontend React para autenticação OAuth 2.0 com PKCE (Proof Key for Code Exchange).

## 📋 Descrição

Este projeto implementa um fluxo completo de autenticação OAuth 2.0 com suporte a PKCE, utilizando React, TanStack Router e TanStack Query. A aplicação fornece uma interface de login segura que se comunica com um servidor OAuth para autenticação de usuários.

### Funcionalidades

- ✅ Autenticação OAuth 2.0 com PKCE
- ✅ Gerenciamento de estado de autenticação
- ✅ Proteção de rotas
- ✅ Armazenamento seguro de tokens
- ✅ Dashboard de usuário autenticado
- ✅ Logout com limpeza de sessão

## 🚀 Como Instalar

### Pré-requisitos

- Node.js (versão 18 ou superior)
- pnpm (gerenciador de pacotes)

### Instalação

```bash
# Clone o repositório
git clone <seu-repositorio>

# Entre na pasta do projeto
cd front

# Instale as dependências
pnpm install
```

### Configuração

Crie um arquivo `.env` na raiz do projeto com as seguintes variáveis:

```env
VITE_OAUTH_CLIENT_ID=teste-front
VITE_OAUTH_AUTH_URL=http://localhost:8080/authorize
VITE_OAUTH_TOKEN_URL=http://localhost:8080/token
VITE_OAUTH_REDIRECT_URI=http://localhost:3000/callback
```

> **Nota:** Ajuste as URLs e o client_id conforme seu servidor OAuth.

## 🏃 Como Rodar

### Modo Desenvolvimento

```bash
pnpm dev
```

A aplicação estará disponível em `http://localhost:3000`

### Build para Produção

```bash
pnpm build
```

### Servir Build de Produção

```bash
pnpm serve
```

## 🛠️ Tecnologias Utilizadas

- **React 19** - Biblioteca UI
- **TypeScript** - Tipagem estática
- **Vite** - Build tool e dev server
- **TanStack Router** - Roteamento file-based
- **TanStack Query** - Gerenciamento de estado assíncrono
- **Tailwind CSS** - Estilização
- **Biome** - Linting e formatação
- **Vitest** - Testes unitários

## 📁 Estrutura do Projeto

```
src/
├── api/           # Chamadas à API OAuth
├── components/    # Componentes React (UI e templates)
├── config/        # Configurações do OAuth
├── hooks/         # Custom hooks (useAuth)
├── lib/           # Utilitários (PKCE, utils)
├── routes/        # Rotas da aplicação
│   ├── index.tsx      # Página de login
│   ├── callback/      # Callback OAuth
│   └── dashboard/     # Dashboard protegido
└── types/         # Definições TypeScript
```

## 🔒 Fluxo de Autenticação

1. **Login** - Usuário clica em "Login" na página inicial
2. **PKCE** - Aplicação gera code_verifier, state e code_challenge
3. **Redirect** - Usuário é redirecionado para o servidor OAuth
4. **Autorização** - Usuário autoriza no servidor OAuth
5. **Callback** - Servidor redireciona para `/callback` com código
6. **Token Exchange** - Aplicação troca código por tokens
7. **Dashboard** - Usuário autenticado acessa o dashboard

## 📝 Scripts Disponíveis

```bash
pnpm dev      # Inicia servidor de desenvolvimento
pnpm build    # Build de produção
pnpm serve    # Serve build de produção
pnpm test     # Executa testes
pnpm lint     # Verifica código com Biome
pnpm format   # Formata código com Biome
pnpm check    # Lint e format juntos
```
