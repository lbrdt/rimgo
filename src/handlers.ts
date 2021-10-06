import Hapi = require('@hapi/hapi');
import '@hapi/vision';
import { fetchAlbumURL, fetchComments, fetchGallery, fetchMedia } from './fetchers';
import * as util from './util';

import CONFIG from './config';

export const handleMedia = async (request: Hapi.Request, h: Hapi.ResponseToolkit) => {
  const {
    baseName,
    extension,
  } = request.params;
  const result = await fetchMedia(`${baseName}.${extension}`);
  const response = h.response(result.rawBody)
    .header('Content-Type', result.headers["content-type"] || `image/${extension}`);
  return response;
};

export const handleAlbum = async (request: Hapi.Request, h: Hapi.ResponseToolkit) => {
  // https://imgur.com/a/DfEsrAB
  const url = await fetchAlbumURL(request.params.albumID);
  return h.view('album', {
    url,
    title: CONFIG.page_title,
    util,
  });
};

export const handleUser = (request: Hapi.Request, h: Hapi.ResponseToolkit) => {
  // https://imgur.com/user/MomBotNumber5
  throw new Error('not implemented');
};

export const handleTag = (request: Hapi.Request, h: Hapi.ResponseToolkit) => {
  // https://imgur.com/t/funny
  throw new Error('not implemented');
};

export const handleGallery = async (request: Hapi.Request, h: Hapi.ResponseToolkit) => {
  const galleryID = request.params.galleryID;
  const gallery = await fetchGallery(galleryID);
  const comments = CONFIG.disable_comments
    ? null
    : await fetchComments(galleryID);
  return h.view('gallery', {
    ...gallery,
    comments,
    title: CONFIG.page_title,
    util,
  });
};
