import Hapi = require('@hapi/hapi');
/* eslint-disable @typescript-eslint/no-var-requires */
const Exiting = require('exiting');
import Path = require('path');
import { handleAlbum, handleGallery, handleMedia, handleTag, handleUser } from './handlers';

import CONFIG from './config';

const server = Hapi.server({
  port: CONFIG.port,
  host: CONFIG.host,
  address: CONFIG.address,
  routes: {
    files: {
      relativeTo: Path.join(__dirname, 'static')
    }
  }
});
server.events.on('stop', () => {
  console.log('Server stopped.');
});

const manager = Exiting.createManager(server);

const init = async () => {
  await server.register(require('@hapi/vision'));
  await server.register(require('@hapi/inert'));

  server.route({
    method: 'GET',
    path: '/css/{param*}',
    handler: ({
      directory: {
        path: Path.join(__dirname, 'static/css')
      }
    } as any) // eslint-disable-line @typescript-eslint/no-explicit-any
  });
  server.views({
    engines: {
      pug: require('pug')
    },
    relativeTo: __dirname,
    path: 'templates',
  });
  server.route({
    method: 'GET',
    path: '/{baseName}.{extension}',
    handler: handleMedia,
  });
  server.route({
    method: 'GET',
    path: '/a/{albumID?}',
    handler: handleAlbum,
  });
  server.route({
    method: 'GET',
    path: '/t/{tagID?}',
    handler: handleTag,
  });
  server.route({
    method: 'GET',
    path: '/user/{userID?}',
    handler: handleUser,
  });
  server.route({
    method: 'GET',
    path: '/gallery/{galleryID}',
    handler: handleGallery,
  });

  await manager.start();
  console.log('Server running on %s', server.info.uri);
};

process.on('unhandledRejection', (err) => {
  console.error(err);
  process.exit(1);
});

init();
