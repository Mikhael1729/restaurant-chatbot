type MessageType = "reciver" | "sender";

export default class Message {
  id: string;
  text: string;
  date: Date;
  type: MessageType;

  constructor() {
    this.id = "";
    this.text = "";
    this.date = new Date(Date.now());
    this.type = "sender";
  }

  static mockData(): Message[] {
    const messages: Message[] = [];
    for (let i = 0; i < 6; i++) {
      const isEven = i % 2 === 0;
      const type = !isEven ? "reciver" : "sender";
      const text = isEven ? "Hola, envío esto" : "Mi respuesta es: No >: 3";

      messages.push({ id: `$m-{i}`, text, type, date: new Date(Date.now()) });
    }

    return messages;
  }
}

