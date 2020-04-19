// Runner (static)

// ----------------------------------------------------------------------------
// IMPORTS

/* Node */
/* NPM */
import chalk from "chalk";
import path from "path";
import clientConfig from "../webpack/client";
/* Local */
import { app, build, common, devServer, staticCompiler } from "./app";

const proxy = require("koa-proxies");
// const httpsProxyAgent = require('https-proxy-agent')

// var proxy = require('koa-proxy');

// ----------------------------------------------------------------------------

common.spinner.info(chalk.bgBlue("Static mode"));

void (async () => {
  // Production?
  if (common.isProduction) {
    common.spinner.info("Building production files...");
    await build(true /* build in static mode */);
    common.spinner.succeed("Finished building");
    return;
  }

  // Development...
  common.spinner.info("Building development server...");

  // const options.events = {
  //   error(err, req, res) { },
  //   proxyReq(proxyReq, req, res) { },
  //   proxyRes(proxyRes, req, res) { }
  // }

  // middleware
  app.use(
    proxy("/auth/login", {
      target: "http://localhost:8090",
      changeOrigin: true,
      // agent: new httpsProxyAgent('http://1.2.3.4:88'), // if you need or just delete this line
      // rewrite: path => path.replace(/^\/octocat(\/|\/\w+)?$/, '/vagusx'),
      logs: true
    })
  );
  app.use(
    proxy("/auth/refresh", {
      target: "http://localhost:8090",
      changeOrigin: true,
      logs: true
    })
  );

  app.use(
    proxy("/query", {
      target: "http://localhost:8090",
      changeOrigin: true,
      // agent: new httpsProxyAgent('http://1.2.3.4:88'), // if you need or just delete this line
      // rewrite: path => path.replace(/^\/octocat(\/|\/\w+)?$/, '/vagusx'),
      logs: true
    })
  );
  // app.use(proxy({
  //   host: 'http://localhost:8090/login'
  // }));

  // app.use(proxy({
  //   host: 'http://localhost:8090/login', // proxy alicdn.com...
  //   match: /^\/auth\/login/,        // ...just the /static folder
  //   jar: true,
  // }));

  // app.use(proxy({
  //   host: 'http://localhost:8090/query'
  // }));

  app.listen({ port: common.port, host: common.host }, async () => {
    // Build the static dev server
    const middleware = await devServer(app, staticCompiler);

    // Fallback to /index.html on 404 routes, for client-side SPAs
    app.use(async ctx => {
      const filename = path.resolve(clientConfig.output.path, "index.html");
      ctx.response.type = "html";
      ctx.response.body = middleware.devMiddleware.fileSystem.createReadStream(
        filename
      );
    });
  });
})();
