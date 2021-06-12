import Message from "../models/Message";
import SendMessageResponse from "../models/SendMessageResponse";

const apiUrl = process.env.VUE_APP_API_URL;

export async function sendMessage(messageText: string): Promise<SendMessageResponse> {
  const response = await fetch(`${apiUrl}/messages`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ text: messageText }),
  });

  return response.json();
}

export async function getMessages(): Promise<Message[]> {
  const response = await fetch(`${apiUrl}/messages`);
  const data: Message[] = await response.json();

  return data;
}
