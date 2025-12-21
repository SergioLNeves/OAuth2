# Git Conventions

## Commit Messages
Use **semantic commit format** with a single, concise description:
```
<type>: <description>
```

## Types
- `feat`: New feature
- `fix`: Bug fix
- `refactor`: Code restructuring
- `docs`: Documentation changes
- `chore`: Maintenance tasks
- `test`: Test additions/updates

## Rules
- Write ONE focused commit message (not a list)
- Keep description concise and meaningful
- Use imperative mood (e.g., "add" not "added")

## Examples
```
feat: add OAuth token refresh endpoint
fix: resolve nil pointer in health check handler
refactor: simplify dependency injection registration
docs: update API documentation with new endpoints
chore: update dependencies to latest versions
test: add unit tests for token validation
```
