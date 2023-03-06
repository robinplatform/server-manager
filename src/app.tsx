import "./app.scss";
import "./global.scss";
import { renderApp } from "@robinplatform/toolkit/react";
import { useRemoteAppMethod } from "@robinplatform/toolkit/react/rpc";
import "@robinplatform/toolkit/styles.css";
import React from "react";
import { z } from "zod";

const App = () => {
	const { data, error } = useRemoteAppMethod(
		"SayHello",
		{},
		{
			resultType: z.object({ message: z.string() }),
		},
	);

	return (
		<div className="appContainer robin-pad robin-bg-dark-blue robin-rounded">
			<>
				{error && <p>{String(error)}</p>}
				{data && <p>{data.message}</p>}
			</>
		</div>
	);
};

renderApp(<App />);
