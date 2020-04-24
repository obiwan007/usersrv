// Webpack (static bundling)

// ----------------------------------------------------------------------------
// IMPORTS

/* NPM */

// Plugin for generating `index.html` file for static hosting
import HtmlWebpackPlugin from "html-webpack-plugin";
import { mergeWith } from "lodash";
import webpack from "webpack";
import {} from "webpack-dev-server";
// Get the client-side config as a base to extend
import client from "./client";
/* Local */
// Common config
import { defaultMerger } from "./common";

// ----------------------------------------------------------------------------

// Augment client-side config with HtmlWebPackPlugin
const base: webpack.Configuration = {
  plugins: [
    new HtmlWebpackPlugin({
      inject: false,
      template: "src/views/static.html",
      title: "ReactQL app"
    })
  ]
};

export default mergeWith({}, client, base, defaultMerger);
