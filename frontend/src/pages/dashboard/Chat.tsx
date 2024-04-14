import { useState, useEffect, useRef } from "react";
import { FaEllipsisV, FaPaperPlane } from "react-icons/fa";
import axios from "axios";
import {  Message } from "../../components/UI/assets";

interface MessageBlockProps {
  message: Message;
}

const MessageBlock = ({ message }: MessageBlockProps) => {
  const isUser = message.role === "user";
  return (
    <div className={`flex w-full ${isUser ? "justify-end" : "justify-start"}`}>
      <div className={`flex w-fit flex-col `}>
        <div
          className={`rounded bg-gray-200 p-5 ${isUser ? "bg-primary-1000 text-primary-50" : "text-black bg-white"}`}
        >
          {message.content}
        </div>
      </div>
    </div>
  );
};

const Chat = ({ pageTitle }: { pageTitle: string }) => {
  useEffect(() => {
    document.title = pageTitle;
  }, [pageTitle]);

  const [messages, setMessages] = useState<Message[]>([]);
  const dummyMessageLastRef = useRef<
    React.LegacyRef<HTMLDivElement> | undefined
  >(null);
  const [message, setMessage] = useState("");

  const getMessages = async () => {
    try {
      const responce = await axios.get("http://localhost:8082/chat?user_id=1")
      const messages = responce.data.Messages


      setMessages(messages)
    } catch (error:any) {
      console.log(error.message)
    }
  };



  const scrollToBottom = () => {
    if (!dummyMessageLastRef.current) return;
    dummyMessageLastRef.current.scrollIntoView({ behavior: "smooth" });
  };

  const storeMessageOnDb = async () => {
    try {
      const responce = await axios.post("http://localhost:8082/chat", {
        content:message,
        user_id:"1"
    })

    const data = responce.data
    

    if (data.answer) {

      console.log(data)


      const answer = data.answer as string

      const newMessage :Message = {
        role:"assistant",
        content:answer
      }
      setMessages((prev)=>[...prev, newMessage ])
      scrollToBottom();

    }    

    } catch (error) {
      console.log(error)
    }
  };

  const handleSendMessage = () => {
    setMessages((prev)=>[...prev, {
      role:"user",
      content:message
    }])
    storeMessageOnDb();
    setMessage("");
    scrollToBottom();
  };

  useEffect(()=>{
    getMessages()
  },[])

  console.log(messages)

  return (
    <>
      <div className="relative box-border flex h-full flex-col">
        <div className="flex flex-grow flex-col gap-5  p-8  ">
          {messages.map((message)=>
          <MessageBlock message={message} />
          )}
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
