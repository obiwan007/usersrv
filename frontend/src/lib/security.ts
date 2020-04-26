import { zipkinFetch } from "./tracing";

class Security {
  inMemoryToken: string = "";
  username: string = "";
  picture: string = "";
  async login(username: string, password: string) {
    const uri = window.location.origin; // uri = GRAPHQL ? GRAPHQL : "http://gqlsrv:8090";
    console.log("Login clicked");
    const response = await zipkinFetch(`${uri}/auth/login`, {
      method: "POST",
      credentials: "same-origin", // include, *same-origin, omit
      // credentials: "include", // include, *same-origin, omit
      headers: {
        "Content-Type": "application/json"
        // 'Content-Type': 'application/x-www-form-urlencoded',
      },
      body: JSON.stringify({ username, password })
    });
    //...
    // Extract the JWT from the response
    const { token } = await response.json();
    console.log("The Token", token);
    this.inMemoryToken = token;
    return true;
  }

  async refresh() {
    const uri = window.location.origin; // GRAPHQL ? GRAPHQL : "http://gqlsrv:8090";
    console.log("Login clicked");
    const response = await zipkinFetch(`${uri}/auth/refresh`, {
      method: "POST",
      credentials: "same-origin", // include, *same-origin, omit
      // credentials: "include", // include, *same-origin, omit
      headers: {
        "Content-Type": "application/json"
        // 'Content-Type': 'application/x-www-form-urlencoded',
      }
    });
    //...
    // Extract the JWT from the response
    if (response.status !== 200) {
      throw response.statusText;
    }

    const { token } = await response.json();
    console.log("The Token", token);
    this.inMemoryToken = token;
    const decoded = this.parseJwt(token)
    console.log('Token:', decoded);
    if (decoded) {
      this.username = decoded.name;
      this.picture = decoded.picture;
    }
    return true;
  }

  parseJwt(token: string) {
    var base64Url = token.split('.')[1];
    var base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
    var jsonPayload = decodeURIComponent(atob(base64).split('').map(function (c) {
      return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
    }).join(''));

    return JSON.parse(jsonPayload);
  };
}

export default new Security();
