import { useEffect, useState } from "react";
import toast, { Toaster } from "react-hot-toast";
import { Chart as ChartJS, ArcElement, Tooltip, Legend } from "chart.js";
import { Pie } from "react-chartjs-2";
import "./App.css";

type Message = {
  text: string;
  category: "debug" | "info" | "warning" | "error";
};

const AlertType = {
  debug: "alert bg-slate-100  border-1 border-slate-400 text-slate-600",
  info: "alert bg-blue-100 border-1 border-blue-400 text-blue-600",
  warning: "alert bg-amber-100  border-1 border-amber-400 text-amber-600",
  error: "alert bg-red-100 border-1 border-red-400 text-red-600",
};

function getRandomElement(array) {
  const randomIndex = Math.floor(Math.random() * array.length);
  return array[randomIndex];
}

ChartJS.register(ArcElement, Tooltip, Legend);

function App() {
  const [alertArray, setAlertArray] = useState<Message[]>([]);

  const addMessage = (message: Message) => {
    setAlertArray([...alertArray, message]);
  };

  const buttonHandler = () => {
    const categories = ["debug", "info", "warning", "error"];
    const category = getRandomElement(categories);
    const toastId = Math.random().toString(36).substring(3);
    addMessage({
      text: `New ${category} alert #${toastId}`,
      category: category,
    });
    toast.custom(
      (t) => (
        <div
          className={`${
            t.visible ? "animate-enter" : "animate-leave"
          } max-w-sm w-full p-2 shadow-md rounded-lg pointer-events-auto flex ring-1 ring-black ring-opacity-5 ${
            AlertType[category]
          }`}
        >
          {`New ${category} alert #${toastId}`}
        </div>
      ),
      {
        id: toastId,
        duration: 800,
      }
    );
  };

  return (
    <>
      <h1 className="text-5xl font-semibold">Alerts Dashboard</h1>
      <p className="text-xl mt-6">
        Total alerts: <span className="font-bold">{alertArray.length}</span>
      </p>
      <p className="text-md my-2">
        <span className="font-bold">debug:</span>{" "}
        {alertArray.filter((alert) => alert.category === "debug").length}
        {" / "}
        <span className="font-bold">info:</span>{" "}
        {alertArray.filter((alert) => alert.category === "info").length}
        {" / "}
        <span className="font-bold">warning:</span>{" "}
        {alertArray.filter((alert) => alert.category === "warning").length}
        {" / "}
        <span className="font-bold">error:</span>{" "}
        {alertArray.filter((alert) => alert.category === "error").length}
      </p>
      <div className="h-64 flex column justify-center items-center m-0 text-center">
        <Pie
          data={{
            labels: ["debug", "info", "warning", "error"],
            datasets: [
              {
                label: "Alerts",
                data: [
                  alertArray.filter((alert) => alert.category === "debug")
                    .length,
                  alertArray.filter((alert) => alert.category === "info")
                    .length,
                  alertArray.filter((alert) => alert.category === "warning")
                    .length,
                  alertArray.filter((alert) => alert.category === "error")
                    .length,
                ],
                backgroundColor: [
                  "rgb(71, 85, 105)",
                  "rgb(37, 99, 235)",
                  "rgb(251, 191, 36)",
                  "rgb(220, 38, 38)",
                ],
              },
            ],
          }}
        />
      </div>
      <button
        className="btn btn-wide btn-primary mt-6"
        onClick={() => buttonHandler()}
      >
        Add Alert
      </button>
      <Toaster position="bottom-left" />
    </>
  );
}

export default App;
