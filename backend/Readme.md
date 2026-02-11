## Login

### Admin
```bash
curl -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email": "admin@sociomile.com", "password": "password123"}'
```

### Agent
```bash
curl -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email": "agent@sociomile.com", "password": "password123"}'
```