# ai-summary-service — Deployment Notes (Fly.io)

## App Name
`ai-summary-service`

---

## Initial Setup

```bash
flyctl launch
```

**Prompts:**

- **App name:** `ai-summary-service`
- **Postgres:** No
- **Deploy now:** Yes

This command:

- generated `fly.toml`  
- built the Docker image  
- pushed it to Fly  
- deployed the app  
- assigned the app a public URL  

---

## Secrets

```bash
flyctl secrets set OPENAI_API_KEY=sk-xxxx
flyctl secrets set USE_MOCK_LLM=false
flyctl secrets set PORT=8080
```

To list secrets:

```bash
flyctl secrets list
```

---

## Restart After Setting Secrets

```bash
flyctl apps restart ai-summary-service
```

---

## Health Check

```bash
curl https://ai-summary-service.fly.dev/health
```

**Expected Response:**

```json
{"status":"ok","service":"ai-summary"}
```

---

## Summary Endpoint Test

```bash
curl -X POST https://ai-summary-service.fly.dev/summary \
  -H "Content-Type: application/json" \
  -d '{"text":"This is a test of the deployed summary service."}'
```

---

## Logs

```bash
flyctl logs -a ai-summary-service
```

---

## Live URL

```
https://ai-summary-service.fly.dev
```

(A custom domain such as `summary.jeffellis.dev` will be added during Day 9.)

---

## Deployment Complete

The AI Summary Service is now fully deployed, tested, and documented.
ai