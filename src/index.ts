'use strict';

import Hapi = require('@hapi/hapi');
import Path  = require('path');
import { handleAlbum, handleGallery, handleMedia, handleTag, handleUser } from './handlers';

import CONFIG from './config';

const init = async () => {
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
  await server.register(require('@hapi/vision'));
  await server.register(require('@hapi/inert'));

  server.route({
    method: 'GET',
    path: '/css/{param*}',
    handler: ({
        directory: {
            path: Path.join(__dirname, 'static/css')
        }
    } as any)
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

  await server.start();
  console.log('Server running on %s', server.info.uri);
};

process.on('unhandledRejection', (err) => {
  console.error(err);
  process.exit(1);
});

init();
