import "./ServerList.scss";
import { useRemoteAppMethod } from "@robinplatform/toolkit/react/rpc";
import cx from "classnames";
import React from "react";
import { z } from "zod";

const ServerType = z.object({
	name: z.string(),
});
type ServerType = z.infer<typeof ServerType>;

const ServerListItem: React.FC<{
	selectedServer: string | null;
	server: ServerType;
	onClick: () => void;
}> = ({ selectedServer, server, onClick }) => {
	return (
		<button
			type={"button"}
			className={cx("serverListItem", {
				active: server.name === selectedServer,
			})}
			onClick={onClick}
		>
			{server.name}
		</button>
	);
};

const ServerListType = z.array(ServerType);

export const ServerList: React.FC<{
	selectedServer: string | null;
	onSelectServer(server: string): void;
}> = ({ selectedServer, onSelectServer }) => {
	const { data: servers, error } = useRemoteAppMethod(
		"GetServers",
		{},
		{
			resultType: ServerListType,
		},
	);

	React.useEffect(() => {
		if (servers?.length && !selectedServer) {
			onSelectServer(servers[0].name);
		}
	}, [servers]);

	return (
		<>
			{error && <p>{String(error)}</p>}

			<div className='serverList'>
				{servers?.map((server) => (
					<ServerListItem
						key={server.name}
						selectedServer={selectedServer}
						server={server}
						onClick={() => onSelectServer(server.name)}
					/>
				))}
			</div>
		</>
	);
};
