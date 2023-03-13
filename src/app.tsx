import { ControlPanel } from "./ControlPanel";
import { ServerList } from "./ServerList";
import "./app.scss";
import { useRpcQuery } from "./bridge";
import "./global.scss";
import { useSelectedServer } from "./hooks/useSelectedServer";
import { runAppMethod } from "@robinplatform/toolkit";
import { renderApp } from "@robinplatform/toolkit/react";
import "@robinplatform/toolkit/styles.css";
import { useMutation } from "@tanstack/react-query";
import React from "react";
import { z } from "zod";

const App = () => {
	const { selectedServer } = useSelectedServer();
	const { mutate } = useMutation<string, unknown, unknown, [string]>({
		mutationKey: ["StartServer"],
		mutationFn: () => {
			runAppMethod({
				methodName: "StartServer",
				resultType: z.any(),
				data: {
					server: selectedServer,
				},
			});
			return null;
		},
	});

	return (
		<div className="appContainer robin-pad robin-bg-dark-blue robin-rounded">
			<div className="leftSidebar">
				<button onClick={mutate}>Start Server</button>

				<ServerList />
			</div>

			{selectedServer && <ControlPanel />}
		</div>
	);
};

renderApp(<App />);
