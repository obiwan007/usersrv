import { zipkinFetch } from "./tracing";

class Security {
  inMemoryToken: string = "";
  async login(username: string, password: string) {
    const uri = GRAPHQL ? GRAPHQL : "http://gqlsrv:8090";
    console.log("Login clicked");
    const response = await zipkinFetch(
      `${uri.replace("/query", "")}/auth/login`,
      {
        method: "POST",
        credentials: "same-origin", // include, *same-origin, omit
        // credentials: "include", // include, *same-origin, omit
        headers: {
          "Content-Type": "application/json"
          // 'Content-Type': 'application/x-www-form-urlencoded',
        },
        body: JSON.stringify({ username: "MyUsername", password: "MyPassword" })
      }
    );
    //...
    // Extract the JWT from the response
    const { token } = await response.json();
    console.log("The Token", token);
    this.inMemoryToken = token;
    return true;
  }
}

export default new Security();
