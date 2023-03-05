import React from "react";
import { renderApp } from "@robinplatform/toolkit/react";
import "@robinplatform/toolkit/styles.css";
import "./global.scss";
import "./app.scss";

const App = () => {
  return (
    <div className="appContainer robin-pad robin-bg-dark-blue robin-rounded">
      <p>hello world</p>
    </div>
  );
};

renderApp(<App />);
