# Docker Configuration

## 📋 Estrutura

```
OAuth/
├── docker-compose.yml          # Orquestração dos serviços
├── back/
│   ├── .env                    # Variáveis de ambiente do backend
│   └── .docker/
│       └── dockerfile.dev      # Dockerfile do backend (Go + Air)
└── front/
    ├── .env                    # Variáveis de ambiente do frontend
    └── .docker/
        └── dockerfile.dev      # Dockerfile do frontend (Node + Vite)
```

## 🔧 Serviços

### Backend (Go + Air)
- **Container**: `oauth-backend`
- **Porta exposta**: `8080` (host) → `8080` (container)
- **Hot-reload**: Air v1.63.4
- **Network**: `oauth-network`
- **Volume**: `./back:/app` (sincronização de código)
- **Env file**: `./back/.env`

### Frontend (Node + Vite)
- **Container**: `oauth-frontend`
- **Porta exposta**: `5173` (host) → `3000` (container)
- **Hot-reload**: Vite dev server com `--host`
- **Network**: `oauth-network`
- **Volume**: `./front:/app` (sincronização de código)
- **Env file**: `./front/.env`

## 🔒 Segurança e Otimizações

### Network Interna (`oauth-network`)
Os containers se comunicam internamente via DNS automático do Docker:
- Frontend acessa backend usando `http://backend:8080` (dentro da network)
- Comunicação direta entre containers (mais rápida e segura)
- Isolamento de outros containers do sistema
- Reduz latência ao evitar passar pelo host

### Configuração de Variáveis de Ambiente

**Backend** (`./back/.env`):
- Todas as configurações do OAuth 2.0
- Database paths e timeouts
- CORS configurado para aceitar o frontend

**Frontend** (`./front/.env`):
- `VITE_OAUTH_AUTH_URL=http://backend:8080/authorize` ← **DNS interno do Docker**
- `VITE_OAUTH_TOKEN_URL=http://backend:8080/token` ← **DNS interno do Docker**
- `VITE_OAUTH_REDIRECT_URI=http://localhost:5173/callback` ← **Callback externo (browser)**

> ⚠️ **Importante**: 
> - O redirect URI usa `localhost` porque é executado no browser do usuário, não no container!
> - As URLs de auth e token usam `backend` (nome do serviço) para comunicação interna entre containers
> - **Os arquivos `.env` não devem subir para produção** - use variáveis de ambiente do sistema em prod

## 🚀 Comandos

### Iniciar os serviços
```bash
docker-compose up --build
```

### Iniciar em background (detached)
```bash
docker-compose up -d --build
```

### Ver logs em tempo real
```bash
docker-compose logs -f
```

### Ver logs de um serviço específico
```bash
docker-compose logs -f backend
docker-compose logs -f frontend
```

### Parar os serviços
```bash
docker-compose down
```

### Parar e remover volumes
```bash
docker-compose down -v
```

### Reconstruir imagens sem cache
```bash
docker-compose build --no-cache
```

### Reiniciar um serviço específico
```bash
docker-compose restart backend
docker-compose restart frontend
```

## 🔄 Hot Reload

Ambos os serviços possuem hot-reload configurado:

- **Backend (Air)**: Detecta mudanças em arquivos `.go` e recompila automaticamente
  - Monitora: `*.go`, `*.tpl`, `*.tmpl`, `*.html`
  - Exclui: `*_test.go`, `tmp/`, `vendor/`
  
- **Frontend (Vite)**: Detecta mudanças em arquivos do projeto e atualiza o browser instantaneamente
  - HMR (Hot Module Replacement) ativo
  - Atualização instantânea sem reload completo da página

Basta editar os arquivos localmente que as mudanças serão refletidas automaticamente!

## 📡 Portas Expostas

| Serviço  | Host         | Container | Acesso Externo              |
|----------|--------------|-----------|------------------------------|
| Backend  | 8080         | 8080      | http://localhost:8080        |
| Frontend | 5173         | 3000      | http://localhost:5173        |

### Comunicação Interna (dentro da network Docker)
- Backend: `http://backend:8080`
- Frontend: `http://frontend:3000`

## 🐛 Debug

### Acessar o shell do container
```bash
# Backend (Alpine Linux)
docker exec -it oauth-backend sh

# Frontend (Alpine Linux)
docker exec -it oauth-frontend sh
```

### Verificar variáveis de ambiente
```bash
docker exec oauth-backend env
docker exec oauth-frontend env
```

### Verificar conectividade entre containers
```bash
# Do frontend, testar conexão com backend
docker exec oauth-frontend wget -qO- http://backend:8080/health

# Verificar DNS resolution
docker exec oauth-frontend nslookup backend
```

### Inspecionar a network
```bash
docker network inspect oauth_oauth-network
```

## 📦 Volumes

### Volumes montados
- `./back:/app` - Código do backend (hot-reload)
- `./front:/app` - Código do frontend (hot-reload)

### Volumes anônimos (excluídos da sincronização)
- `/app/tmp` - Binários compilados do Go (backend)
- `/app/node_modules` - Dependências do Node (frontend)

> 💡 Os volumes anônimos evitam conflitos entre dependências do host e do container

## 🔐 Segurança em Produção

Para produção, **NÃO use os arquivos `.env`**:

1. Configure variáveis de ambiente no sistema/orquestrador (Kubernetes, Docker Swarm, etc)
2. Use secrets management (AWS Secrets Manager, HashiCorp Vault, etc)
3. Nunca commite arquivos `.env` no Git
4. Use `.env.example` para documentar variáveis necessárias

### Exemplo de `.env.example` para produção:
```bash
# Backend
OAUTH_ISSUER=https://auth.seudominio.com
SESSION_SECRET=<gerar-novo-secret>
# ... outras variáveis

# Frontend
VITE_OAUTH_AUTH_URL=https://auth.seudominio.com/authorize
VITE_OAUTH_TOKEN_URL=https://auth.seudominio.com/token
VITE_OAUTH_REDIRECT_URI=https://app.seudominio.com/callback
```

## 📝 Notas Técnicas

- **Network**: `oauth-network` usa driver `bridge` (padrão Docker para comunicação container-to-container)
- **DNS**: Docker fornece DNS automático - containers podem se acessar pelo nome do serviço
- **Restart policy**: `unless-stopped` garante que containers reiniciem após falhas ou reboot do host
- **Dependency**: `depends_on` garante que backend inicie antes do frontend
- **Build context**: Cada serviço tem seu próprio contexto de build isolado
