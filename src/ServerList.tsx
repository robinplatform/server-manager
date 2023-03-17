import "./ServerList.scss";
import { ServerType, useRpcQuery } from "./bridge";
import { useHistory } from "./hooks/useHistory";
import { useSelectedServer } from "./hooks/useSelectedServer";
import cx from "classnames";
import React from "react";

const ServerListItem: React.FC<{
	server: ServerType;
	onClick: () => void;
}> = ({ server, onClick }) => {
	const { selectedServer } = useSelectedServer();
	const { data: health } = useRpcQuery("CheckServerHealth", {
		name: server.name,
	});

	return (
		<button
			type={"button"}
			className={cx("serverListItem", {
				active: server.name === selectedServer,
			})}
			onClick={onClick}
		>
			{server.name} {health}
		</button>
	);
};

export const ServerList: React.FC = () => {
	const { selectedServer, setSelectedServer } = useSelectedServer();
	const history = useHistory();
	const { data: servers, error } = useRpcQuery("GetServers", {});

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
