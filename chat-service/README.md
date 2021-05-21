# Chatbot API

The chatbot service for the restaurant.

## Commands

**Run the code:**

```bash
go run main.go
```

**Test the chat endpoint:**

```bash
curl localhost:9090/chat
```

**Pass data to the chat endpoint:**

```bash
curl -d '{message: "Buenas"}' localhost:9090/chat
```

You can add the flag `-v` to the two previous commands to see more information
