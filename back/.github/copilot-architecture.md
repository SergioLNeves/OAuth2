# Architecture - OAuth 2.0 Backend

## Stack
- **Go 1.25.4**
- **Echo v4** - Web framework
- **Uber Dig** - Dependency injection
- **go-env** - Environment management
- **go-playground/validator** - Request validation with i18n support

## Architecture Pattern
**Hexagonal Architecture (Ports & Adapters)**

Organização que mantém o core da aplicação isolado de detalhes de infraestrutura.

## Project Structure
```
cmd/app/main.go              # Entry point
internal/
├── core/                    # Application Core (business logic)
│   ├── domain/             # Business entities
│   ├── ports/              # Interfaces (contracts)
│   └── services/           # Use cases / Application services
├── adapters/               # External adapters
│   ├── http/              # HTTP handlers & router
│   └── config/            # Configuration & API server
│       ├── api.go         # Graceful shutdown server
│       ├── config.go      # Config loader
│       └── model.go       # Config structs
├── pkg/                   # Shared utilities
│   └── validator.go       # Custom validator with translations
└── dependencies.go        # DI container (Uber Dig)
```

## Hexagonal Architecture Layers

### Core (Business Logic)
- **domain/** - Entidades de negócio puras, sem dependências externas
- **ports/** - Interfaces que definem contratos (Service interfaces, Handler interfaces)
- **services/** - Implementação de casos de uso e lógica de aplicação

### Adapters (Infrastructure)
- **http/** - Adaptadores HTTP (handlers implementam ports, router agrupa rotas)
- **config/** - Adaptador de configuração e servidor

### Dependency Flow
```
Adapters → Ports → Core
```
- Core não conhece adapters
- Adapters dependem de ports (interfaces)
- DI injeta adapters nas portas

## Module Path
`github.com/SergioLNeves/OAuth2/back`
