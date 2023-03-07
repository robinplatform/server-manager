import { useHistory } from "./useHistory";
import { atom, useAtom } from "jotai";
import React from "react";

const selectedServerAtom = atom<string | null>(null);

export function useSelectedServer() {
	const [selectedServer, setSelectedServer] = useAtom(selectedServerAtom);
	const history = useHistory();
	const setSelectedServerWrapper = React.useCallback(
		(server: string) => {
			setSelectedServer(server);
			history.push(`/server/${server}`);
		},
		[history],
	);

	return { selectedServer, setSelectedServer: setSelectedServerWrapper };
}
