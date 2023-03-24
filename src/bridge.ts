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
  CheckServerHealth: {
    input: z.object({
      name: z.string(),
    }),
    output: z.string(),
  },
});

export { useRpcQuery };
