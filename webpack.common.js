const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');

module.exports = {
    entry: {
        client: './client/index.js'
    },
    output: {
        filename: '[name].bundle.[contenthash].js',
        path: path.resolve(__dirname, 'public')
    },
    module: {
        rules: [
            {
                test: /\.js$/,
                exclude: /(node_modules)/,
                use: {
                    loader: 'babel-loader'
                }
            }
        ]
    },
    plugins: [
        new HtmlWebpackPlugin({
            title: "Jira Stats",
            template: "./client/index.html"
        })
    ]
};
