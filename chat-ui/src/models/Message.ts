export enum Sender {
  Customer = 1,
  Bot = 0,
}

export default class Message {
  id: string;
  text: string;
  dateTime: string;
  sender: Sender;

  constructor() {
    this.id = "";
    this.text = "";
    this.dateTime = "";
    this.sender = Sender.Customer;
  }

  static mockData(): Message[] {
    const messages: Message[] = [];
    for (let i = 0; i < 12; i++) {
      const isEven = i % 2 === 0;
      const sender = isEven ? Sender.Customer : Sender.Bot;
      const text = isEven ? "Hola, envío esto" : "Mi respuesta es: No >: 3";

      messages.push({ id: `$m-{i}`, text, sender, dateTime: "" });
    }

    return messages;
  }
}

