import run from './run';
import clean from './clean';
import bundle from './bundle';
import Webpack from 'webpack';
import WebpackDevServer from 'webpack-dev-server';
import WebpackConfig from '../webpack.config';

async function server() {
  await run(clean);
  await run(bundle);

  const compiler = Webpack(WebpackConfig);
  const server = new WebpackDevServer(compiler, WebpackConfig.devServer);

  return new Promise((resolve, reject) => {
    server.listen(4000, "0.0.0.0", () => {
      console.log("Starting server on http://localhost:4000");
      resolve();
    });
  });
}

export default server;
