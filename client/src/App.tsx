import { useState } from "react";
import "./App.css";

import Alert from "./components/Alert/Alert";

type Message = {
  text: string;
  category: "debug" | "info" | "warning" | "error";
};

function getRandomElement(array) {
  const randomIndex = Math.floor(Math.random() * array.length);
  return array[randomIndex];
}

function App() {
  const [alertArray, setAlertArray] = useState<Message[]>([]);

  const addMessage = (message: Message) => {
    setAlertArray([...alertArray, message]);
  };

  const buttonHandler = () => {
    const categories = ["debug", "info", "warning", "error"];
    const category = getRandomElement(categories);
    addMessage({
      text: `New ${category} alert #${(Math.random() + 1)
        .toString(36)
        .substring(3)}`,
      category: category,
    });
  };

  return (
    <>
      <h1 className="text-5xl font-semibold">Alerts Dashboard</h1>
      <div className="w-full my-8 h-96 overflow-y-scroll">
        {alertArray
          .slice(0)
          .reverse()
          .map((message, index) => {
            return (
              <Alert
                key={index}
                message={message.text}
                category={message.category}
              />
            );
          })}
      </div>
      <button
        className="btn btn-wide btn-primary"
        onClick={() => buttonHandler()}
      >
        Add Alert
      </button>
    </>
  );
}

export default App;
