# Chatbot API

It receives a text message from the customer to the restaurant and returns the classification of the given message.

## Steps

1. Get the training data
2. Clean it
3. Build a classification neural network that recives a text message and outputs the classification for the text.
    - Must be resilient to typographic and spelling mistakes (which implies the network has to understand them)
    - Train the model with the training dataset and save it in a file.
  4. Create the chatbot endpoint

## System overview

Just a fun graphic:

```
     message
User ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
   ↑                                      ↓
   ┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━ Endpoint
                           classification
```
