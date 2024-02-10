import { useState, useEffect } from "react";
import useWebSocket, { ReadyState } from "react-use-websocket";

import toast, { Toaster } from "react-hot-toast";
import { Chart as ChartJS, ArcElement, Tooltip, Legend } from "chart.js";
import { Pie } from "react-chartjs-2";

import "./App.css";

type Alert = {
  id: number;
  message: string;
  category: "debug" | "info" | "warning" | "error";
};

const AlertType = {
  debug: "alert bg-slate-100  border-1 border-slate-400 text-slate-600",
  info: "alert bg-blue-100 border-1 border-blue-400 text-blue-600",
  warning: "alert bg-amber-100  border-1 border-amber-400 text-amber-600",
  error: "alert bg-red-100 border-1 border-red-400 text-red-600",
};

ChartJS.register(ArcElement, Tooltip, Legend);

function App() {
  const [alertHistory, setAlertHistory] = useState<Alert[]>([]);

  const [socketUrl, setSocketUrl] = useState("ws://localhost:8080/alerts");

  const { sendMessage, lastMessage, readyState } = useWebSocket(socketUrl);

  const connectionStatus = {
    [ReadyState.CONNECTING]: "Connecting",
    [ReadyState.OPEN]: "Open",
    [ReadyState.CLOSING]: "Closing",
    [ReadyState.CLOSED]: "Closed",
    [ReadyState.UNINSTANTIATED]: "Uninstantiated",
  }[readyState];

  useEffect(() => {
    if (lastMessage !== null) {
      const alert: Alert = JSON.parse(lastMessage.data);
      setAlertHistory((prev) => prev.concat(alert));
      toast.custom(
        (t) => (
          <div
            className={`${
              t.visible ? "animate-enter" : "animate-leave"
            } max-w-sm w-full p-2 shadow-md rounded-lg pointer-events-auto flex ring-1 ring-black ring-opacity-5 ${
              AlertType[alert.category]
            }`}
          >
            {alert.message}
          </div>
        ),
        {
          id: alert.id.toString(),
          duration: 1200,
        }
      );
    }
  }, [lastMessage]);

  return (
    <>
      <h1 className="text-5xl font-semibold">Alerts Dashboard</h1>
      <p className="text-xl mt-6">
        Total alerts: <span className="font-bold">{alertHistory.length}</span>
      </p>
      <p className="text-md my-2">
        <span className="font-bold">debug:</span>{" "}
        {alertHistory.filter((alert) => alert.category === "debug").length}
        {" / "}
        <span className="font-bold">info:</span>{" "}
        {alertHistory.filter((alert) => alert.category === "info").length}
        {" / "}
        <span className="font-bold">warning:</span>{" "}
        {alertHistory.filter((alert) => alert.category === "warning").length}
        {" / "}
        <span className="font-bold">error:</span>{" "}
        {alertHistory.filter((alert) => alert.category === "error").length}
      </p>
      <small className="text-xs my-3 italic">
        Connection Status: {connectionStatus}
      </small>
      <div className="h-64 flex column justify-center items-center m-0 text-center">
        <Pie
          data={{
            labels: ["debug", "info", "warning", "error"],
            datasets: [
              {
                label: "Alerts",
                data: [
                  alertHistory.filter((alert) => alert.category === "debug")
                    .length,
                  alertHistory.filter((alert) => alert.category === "info")
                    .length,
                  alertHistory.filter((alert) => alert.category === "warning")
                    .length,
                  alertHistory.filter((alert) => alert.category === "error")
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
      <div className="text-xs flex flex-col items-center w-full h-48 overflow-auto my-8">
        {alertHistory.reverse().map((alert) => (
          <div
            key={alert.id}
            className="flex flex-row items-center border-t gap-2 w-96 border-t-slate-200 py-2"
          >
            <span className={`w-2 h-2 p-0 ${AlertType[alert.category]}`}></span>
            {alert.message}
          </div>
        ))}
      </div>
      <Toaster position="bottom-left" />
    </>
  );
}

export default App;
