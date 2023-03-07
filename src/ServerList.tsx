import "./ServerList.scss";
import { useHistory } from "./hooks/useHistory";
import { useSelectedServer } from "./hooks/useSelectedServer";
import { useRemoteAppMethod } from "@robinplatform/toolkit/react/rpc";
import cx from "classnames";
import React from "react";
import { z } from "zod";

const ServerType = z.object({
	name: z.string(),
});
type ServerType = z.infer<typeof ServerType>;

const ServerListItem: React.FC<{
	server: ServerType;
	onClick: () => void;
}> = ({ server, onClick }) => {
	const { selectedServer } = useSelectedServer();

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

export const ServerList: React.FC = () => {
	const { selectedServer, setSelectedServer } = useSelectedServer();
	const history = useHistory();
	const { data: servers, error } = useRemoteAppMethod(
		"GetServers",
		{},
		{
			resultType: ServerListType,
		},
	);

	React.useEffect(() => {
		if (servers?.length && !selectedServer) {
			setSelectedServer(servers[0].name);
		}
	}, [servers]);

	return (
		<>
			{error && <p>{String(error)}</p>}

			<div className='serverList'>
				{servers?.map((server) => (
					<ServerListItem
						key={server.name}
						server={server}
						onClick={() => {
							setSelectedServer(server.name);
							history.push(`/server/${server.name}`);
						}}
					/>
				))}
			</div>
		</>
	);
};
