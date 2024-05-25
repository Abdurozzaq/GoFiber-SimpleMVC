const path = require('path');

module.exports = {
    entry: './src/views/js/main.js',  // Entry point for your JavaScript
    output: {
        filename: 'bundle.js',  // Output bundle filename
        path: path.resolve(__dirname, 'src/public/js'),  // Output directory
    },
    // Add rules for handling different file types (e.g., JS, CSS, etc.)
    module: {
        rules: [
            {
                test: /\.js$/,
                exclude: /node_modules/,
                use: 'babel-loader',  // Use Babel for transpilation
            },
            // Add more rules for CSS, images, etc.
        ],
    },
};