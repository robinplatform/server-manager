import { createBrowserHistory } from "history";
import React from "react";

const history = createBrowserHistory();
const pathBasename = "/api/app-resources/server-manager/base";

export function useHistory() {
	const [counter, setCounter] = React.useState(0);
	React.useEffect(() => {
		const unlisten = history.listen(() => {
			setCounter(counter + 1);
		});
		return () => {
			unlisten();
		};
	}, []);

	const historyWrapper = React.useMemo(
		() => ({
			pathname: history.location.pathname.startsWith(pathBasename)
				? history.location.pathname.slice(pathBasename.length)
				: history.location.pathname,
			push: (path: string) => {
				history.push(pathBasename + path);
			},
		}),
		[history.location.pathname, history.location.search],
	);
	return historyWrapper;
}
