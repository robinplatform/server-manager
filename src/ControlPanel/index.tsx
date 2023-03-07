import { useHistory } from "../hooks/useHistory";
import { useSelectedServer } from "../hooks/useSelectedServer";
import React from "react";

export const ControlPanel: React.FC = () => {
	const { selectedServer } = useSelectedServer();
	const history = useHistory();

	return (
		<div className='serverControlPanel'>
			<h1 className='serverControlPanelHeading'>{selectedServer}</h1>

			<p>{history.pathname}</p>
		</div>
	);
};
