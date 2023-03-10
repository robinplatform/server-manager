import { createReactRpcBridge } from "@robinplatform/toolkit/react/rpc";
import { z } from "zod";

export const ServerType = z.object({
  name: z.string(),
});
export type ServerType = z.infer<typeof ServerType>;

const { useRpcQuery } = createReactRpcBridge({
  GetServers: {
    input: z.object({}),
    output: z.array(ServerType),
  },
  StartServer: {
    input: z.object({}),
    output: z.any(),
  },
});

export { useRpcQuery };
