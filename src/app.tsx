import { ControlPanel } from "./ControlPanel";
import { ServerList } from "./ServerList";
import "./app.scss";
import "./global.scss";
import { useSelectedServer } from "./hooks/useSelectedServer";
import { renderApp } from "@robinplatform/toolkit/react";
import "@robinplatform/toolkit/styles.css";
import React from "react";

const App = () => {
	const { selectedServer } = useSelectedServer();

	return (
		<div className="appContainer robin-pad robin-bg-dark-blue robin-rounded">
			<div className="serverListContainer">
				<ServerList />
			</div>

			{selectedServer && <ControlPanel />}
		</div>
	);
};

renderApp(<App />);
