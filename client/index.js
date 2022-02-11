import { client } from "twirpscript";
import { nodeHttpTransport } from "twirpscript/dist/node/index.js";
import {GetUserProfile, SetUserProfile} from "./service.pb.js";

client.baseURL = "http://localhost:8080";

// This is provided as a convenience for Node.js clients. If you provide `fetch` globally, this isn't necessary and your client can look identical to the browser client above.
client.rpcTransport = nodeHttpTransport;

// const profile = await SetUserProfile({
//     userID: "test-user-id",
//     email: "test@test.com",
//     username: "test name",
// } );

const profile = await GetUserProfile({
    userID: "test-user-id",
} );

console.log(profile);