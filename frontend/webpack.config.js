const path = require("path");
const VueLoaderPlugin = require("vue-loader/lib/plugin");
//const CompressionPlugin = require('compression-webpack-plugin');
const webpack = require("webpack");

module.exports = {
  entry: path.resolve(__dirname, "./src/main.js"),
  context: path.resolve(__dirname, "frontend"),
  output: {
    filename: "bundle.js",
    path: path.resolve(__dirname, "public/javascript"),
  },
  module: {
    rules: [
      {
        test: /\.vue$/,
        loader: "vue-loader",
      },
      {
        test: /\.js$/,
        exclude: /node_modules/,
        loader: "babel-loader",
      },
      {
        test: /\.css$/,
        use: ["vue-style-loader", "css-loader"],
      },
      {
        test: /\.scss$/,
        use: ["vue-style-loader", "css-loader", "sass-loader"],
      },
    ],
  },
  resolve: {
    extensions: [".js", ".vue", ".json"],
    alias: {
      vue$: "vue/dist/vue.esm.js",
    },
  },
  plugins: [
    new VueLoaderPlugin(),
    // Ignore all locale files of moment.js
    new webpack.IgnorePlugin(/^\.\/locale$/, /moment$/),
  ],
};
