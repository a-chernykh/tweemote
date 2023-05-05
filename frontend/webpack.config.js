import path from 'path';
import HtmlWebpackPlugin from 'html-webpack-plugin';
import webpack from 'webpack';
import merge from 'webpack-merge';

const isDebug = false,
      isVerbose = false;

if (!process.env.CONFIG) {
  process.env.CONFIG = "dev"
}

let config = {
  context: path.resolve(__dirname, 'src'),

  output: {
    filename: 'bundle.js',
    path: path.resolve(__dirname, 'dist'),
    publicPath: '/'
  },

  module: {
    rules: [
      {
        test: /\.js$/,
        exclude: /(node_modules|bower_components)/,
        use: {
          loader: 'babel-loader',
					options: {
						presets: ['env'],
            plugins: ["transform-es2015-destructuring", "transform-object-rest-spread"]
					}
        }
      },

      {
        test: /\.css$/,
        use: [ 'style-loader', 'css-loader' ]
      },

      {
        test: /\.less$/,
        use: [ 'style-loader', 'css-loader', 'less-loader' ]
      },

      {
        test: /\.(png|jpg|jpeg|gif|woff|woff2|eot|ttf|svg)$/,
        use: [ 'url-loader' ]
      }
    ]
  },

  plugins: [
    new webpack.HotModuleReplacementPlugin(),
    new webpack.NamedModulesPlugin(),
    new HtmlWebpackPlugin({
      template: 'index.ejs'
    }),
    new webpack.ProvidePlugin({
      $: "jquery",
      jQuery: "jquery",
      "window.jQuery": "jquery"
    })
  ],

  resolve: {
    modules: [
      path.resolve('./src'),
      path.resolve('./node_modules')
    ],
    alias: {
      config: path.join(__dirname, "src/config/" + process.env.CONFIG + ".js"),
    }
  },

  stats: {
    colors: true,
    reasons: isDebug,
    hash: isVerbose,
    version: isVerbose,
    timings: true,
    chunks: isVerbose,
    chunkModules: isVerbose,
    cached: isVerbose,
    cachedAssets: isVerbose,
  },
};

if (process.env.NODE_ENV != 'production') {
  config = merge(config, {
    devtool: 'inline-source-map',

    devServer: {
      contentBase: path.join(__dirname, "dist"),
      compress: true,
      hot: true,
      publicPath: '/',
      historyApiFallback: true,
      disableHostCheck: true,
    },

    entry: [
      'babel-polyfill',

      'react-hot-loader/patch',
      'webpack-dev-server/client',
      'webpack/hot/only-dev-server',

      './app.js'
    ],
  });
} else {
  config.plugins.push(new webpack.DefinePlugin({'process.env': { NODE_ENV: JSON.stringify('production') }}));
  config.plugins.push(new webpack.optimize.UglifyJsPlugin());
  config.entry = [
    'babel-polyfill',
    './app.js'
  ];
}

const configConst = config;

export default configConst;
