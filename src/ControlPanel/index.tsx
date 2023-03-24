import { useHistory } from "../hooks/useHistory";
import { useSelectedServer } from "../hooks/useSelectedServer";
import { runAppMethod } from "@robinplatform/toolkit";
import { useMutation } from "@tanstack/react-query";
import React from "react";
import { z } from "zod";

export const ControlPanel: React.FC = () => {
	const { selectedServer } = useSelectedServer();
	const history = useHistory();
	const { mutate } = useMutation<null, unknown, void, [string]>({
		mutationKey: ["StartServer", selectedServer],
		mutationFn: async (): Promise<null> => {
			await runAppMethod({
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
		<div className='serverControlPanel'>
			<h1 className='serverControlPanelHeading'>{selectedServer}</h1>

			<p>{history.pathname}</p>

			<button onClick={() => mutate()}>Start</button>
		</div>
	);
};
