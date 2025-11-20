# AI Summary Service

A lightweight Go microservice that generates text summaries using either real OpenAI LLM responses or a local mock summarizer.  
This service is part of a 4-project AI Backend Portfolio.

---

## Features

- `POST /summary` → returns a summarized version of the input text  
- Go Fiber server  
- Simple logging middleware  
- Mock mode (`USE_LLM_MOCK=true`)  
- Real OpenAI mode (`USE_LLM_MOCK=false`)  
- Environment-based configuration  
- Fly.io deployment ready  
- No secrets committed (local `.env` only)

---

## Endpoint: POST /summary

### Request Body (JSON)

```json
{
  "text": "Explain the difference between HTTP and HTTPS in 3 sentences."
}
```

### Example Response (Real OpenAI)

```json
{
  "summary": "HTTP transfers data without encryption, making it vulnerable. HTTPS encrypts traffic using SSL/TLS. This protects sensitive data from interception."
}
```

### Example Response (Mock Mode)

```json
{
  "summary": "MOCK SUMMARY: Explain the difference between HTTP and HTTPS in 3 sentences."
}
```

---

## Environment Variables

### `.env` (local only — not committed)

```env
OPENAI_API_KEY=your_real_openai_key_here
USE_LLM_MOCK=false
PORT=8080
```

### `.env.example` (safe to commit)

```env
OPENAI_API_KEY=your_openai_api_key_here
USE_LLM_MOCK=false
PORT=8080
```

---

## Run Locally

Install dependencies:

```bash
go mod tidy
```

Start the server:

```bash
go run .
```

Test:

```bash
curl -X POST http://localhost:8080/summary \
  -H "Content-Type: application/json" \
  -d '{"text":"Summarize this paragraph."}'
```

---

## Deploy to Fly.io

Set secrets:

```bash
flyctl secrets set OPENAI_API_KEY="your_real_key" -a ai-summary-service
flyctl secrets set USE_LLM_MOCK=false -a ai-summary-service
```

Deploy:

```bash
flyctl deploy -a ai-summary-service
```

---

## Project Structure

```text
ai-summary-service/
├── main.go
├── go.mod
├── .gitignore
├── .env.example
├── internal/
│   ├── ai/
│   │   ├── openai.go
│   │   └── mock.go
│   ├── handlers/
│   │   └── summary_handler.go
│   └── middleware/
│       └── logger.go
└── Dockerfile
```

---

## Portfolio Context

This service is part of a 4-project AI Backend Portfolio:

1. Notes-Memory-Core  
2. Notes-Memory-Core-RAG  
3. AI Summary Service  
4. AI Embedding Microservice  

Together, these projects demonstrate:

- Go backend development  
- Microservice architecture  
- AI integration (LLMs + embeddings)  
- Secure secret handling practices  
- Production deployments with Fly.io  
