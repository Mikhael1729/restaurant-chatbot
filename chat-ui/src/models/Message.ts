export enum Sender {
  Customer = 0,
  Bot = 1,
}

type Message = {
  id: string;
  text: string;
  dateTime: string;
  sender: Sender;
  category: string;
};

export default Message;
