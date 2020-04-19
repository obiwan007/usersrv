// Runner (production)

// ----------------------------------------------------------------------------
// IMPORTS

/* Node */
/* NPM */
import chalk from "chalk";
import fs from "fs";
/* Local */
import Output from "../lib/output";
import Stats, { IStats } from "../lib/stats";
import { app, build, common } from "./app";
const proxy = require("koa-proxies");

// ----------------------------------------------------------------------------

function getStats(file: string): IStats {
  return JSON.parse(fs.readFileSync(file, "utf8")) as IStats;
}

common.spinner.info(chalk.green("Production mode"));

void (async () => {
  // Get a list of file accessibility
  const files = Object.values(common.compiled).map(file => {
    try {
      fs.accessSync(file);
      return true;
    } catch (_e) {
      return false;
    }
  });

  // Compile the server if we don't have all the expected files
  if (!files.every(file => file)) {
    common.spinner.info("Building production server...");
    await build();
  } else {
    common.spinner.info("Using cached build files");
  }

  // Create an Output
  const output = new Output({
    client: new Stats(getStats(common.compiled.clientStats)),
    server: new Stats(getStats(common.compiled.serverStats))
  });

  // Attach middleware
  app.use(require(common.compiled.server).default(output));

  app.use(
    proxy("/auth/login", {
      target: "http://gqlsrv:8090",
      changeOrigin: true,
      // agent: new httpsProxyAgent('http://1.2.3.4:88'), // if you need or just delete this line
      // rewrite: path => path.replace(/^\/octocat(\/|\/\w+)?$/, '/vagusx'),
      logs: true
    })
  );
  app.use(
    proxy("/auth/refresh", {
      target: "http://gqlsrv:8090",
      changeOrigin: true,
      logs: true
    })
  );

  app.use(
    proxy("/query", {
      target: "http://gqlsrv:8090",
      changeOrigin: true,
      // agent: new httpsProxyAgent('http://1.2.3.4:88'), // if you need or just delete this line
      // rewrite: path => path.replace(/^\/octocat(\/|\/\w+)?$/, '/vagusx'),
      logs: true
    })
  );

  app.listen(common.port, () => {
    common.spinner.succeed(`Running on http://localhost:${common.port}`);
  });
})();
