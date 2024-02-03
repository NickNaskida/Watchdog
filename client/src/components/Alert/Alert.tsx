type AlertProps = {
  message: string;
  category: "debug" | "info" | "warning" | "error";
};

const AlertType = {
  debug: "alert bg-slate-100  border-1 border-slate-400 text-slate-600",
  info: "alert bg-blue-100 border-1 border-blue-400 text-blue-600",
  warning: "alert bg-amber-100  border-1 border-amber-400 text-amber-600",
  error: "alert bg-red-100 border-1 border-red-400 text-red-600",
};

const Alert = (alert: AlertProps) => {
  return (
    <div
      role="alert"
      className={`alert my-2 ${AlertType[alert.category]}`}
    >
      <span>{alert.message}</span>
    </div>
  );
};

export default Alert;
