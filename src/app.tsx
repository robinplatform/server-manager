import { ServerList } from "./ServerList";
import "./app.scss";
import "./global.scss";
import { renderApp } from "@robinplatform/toolkit/react";
import "@robinplatform/toolkit/styles.css";
import React from "react";

const App = () => {
	const [selectedServer, setSelectedServer] = React.useState<string | null>(
		null,
	);

	return (
		<div className="appContainer robin-pad robin-bg-dark-blue robin-rounded">
			<div className="serverListContainer">
				<ServerList
					selectedServer={selectedServer}
					onSelectServer={(server) => setSelectedServer(server)}
				/>
			</div>

			{selectedServer && (
				<div className='serverControlPanel'>
					<h1 className='serverControlPanelHeading'>{selectedServer}</h1>
				</div>
			)}
		</div>
	);
};

renderApp(<App />);
