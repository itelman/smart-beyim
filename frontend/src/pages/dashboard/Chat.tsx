import { useState, useEffect, useRef } from "react";
import { FaEllipsisV, FaPaperPlane } from "react-icons/fa";
import axios from "axios";
import { staticMessages, Message } from "../../components/UI/assets";

interface MessageBlockProps {
  message: Message;
}

const MessageBlock = ({ message }: MessageBlockProps) => {
  const isUser = message.sender._id === 0;
  return (
    <div className={`flex w-full ${isUser ? "justify-end" : "justify-start"}`}>
      <div className={`flex w-fit flex-col `}>
        <div
          className={`rounded bg-gray-200 p-5 ${isUser ? "bg-primary-1000 text-primary-50" : "text-black bg-white"}`}
        >
          {message.text}
        </div>
        <span
          className={`text-sm  text-secondary-400 ${isUser ? "text-right" : "text-left"}`}
        >
          {message?.date.split("T")[1].slice(0, 5)}
        </span>
      </div>
    </div>
  );
};

const Chat = ({ pageTitle }: { pageTitle: string }) => {
  useEffect(() => {
    document.title = pageTitle;
  }, [pageTitle]);

  const [messages, setMessages] = useState([...staticMessages]);
  const dummyMessageLastRef = useRef<
    React.LegacyRef<HTMLDivElement> | undefined
  >(null);
  const [message, setMessage] = useState("");

  const getMessages = async () => {};

  const scrollToBottom = () => {
    if (!dummyMessageLastRef.current) return;
    dummyMessageLastRef.current.scrollIntoView({ behavior: "smooth" });
  };

  const storeMessageOnDb = () => {
    setMessages((prev) => [
      ...prev,
      { text: message, date: new Date().toISOString(), sender: { _id: 0 } },
    ]);
  };

  const handleSendMessage = () => {
    storeMessageOnDb();
    setMessage("");
    scrollToBottom();
  };

  const getDateToDisplay = (date: any) => {
    const today = new Date();
    const messageDate = new Date(date);
    if (
      today.getFullYear() === messageDate.getFullYear() &&
      today.getMonth() === messageDate.getMonth() &&
      today.getDate() === messageDate.getDate()
    ) {
      return "Today";
    }
    const messageDateDay = messageDate.getDate();
    const messageDateMonth = messageDate.toLocaleString("default", {
      month: "long",
    });
    return messageDateDay + " " + messageDateMonth;
  };
  const getChatBodyContent = () => {
    let prevDay = new Date("12/10/2022").getDay();
    const content = messages.map((message) => {
      if (prevDay !== new Date(message.date).getDay()) {
        prevDay = new Date(message.date).getDay();
        return (
          <>
            <div className="my-5 flex items-center justify-center gap-2.5">
              <div className="h-px flex-grow bg-gray-300"></div>
              <span className="text-lg whitespace-nowrap text-gray-600">
                {getDateToDisplay(message.date)}
              </span>
              <div className="h-px flex-grow bg-gray-300"></div>
            </div>
            <MessageBlock message={message} />
          </>
        );
      } else if (prevDay === new Date(message.date).getDay()) {
        prevDay = new Date(message.date).getDay();

        return <MessageBlock message={message} />;
      }
    });
    return content;
  };

  return (
    <>
      <div className="relative box-border flex h-full flex-col">
        <div className="flex flex-grow flex-col gap-5  p-8  ">
          {getChatBodyContent()}
          <div ref={dummyMessageLastRef}></div>
        </div>
        <div className="sticky bottom-0 left-0 flex flex-row justify-between gap-2.5   px-8 py-4">
          <input
            type="text"
            placeholder="Write message..."
            className="w-full rounded-xl border border-gray-300 py-3 pl-5 pr-3 text-subHeading3  outline-none focus:border-primary-1000"
            value={message}
            onChange={(e) => {
              setMessage(e.target.value);
            }}
            onKeyDown={(e) => {
              if (e.key === "Enter") {
                handleSendMessage();
              }
            }}
            // ... rest of the props
          />
          <button
            type="button"
            className="flex flex-row items-center gap-2.5 rounded-xl bg-primary-1000 px-5 py-2.5 text-white"
            onClick={handleSendMessage}
          >
            <FaPaperPlane className="h-5 w-5 text-white" />
          </button>
        </div>
      </div>
    </>
  );
};

export default Chat;
