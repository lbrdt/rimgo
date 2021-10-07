import Hapi = require('@hapi/hapi');
import '@hapi/vision';
import { fetchAlbum, fetchAlbumURL, fetchComments, fetchGallery, fetchMedia, fetchUserPosts } from './fetchers';
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
  const albumID = request.params.albumID;
  if (CONFIG.use_api) {
    const album = await fetchAlbum(albumID);
    return h.view('gallery', {
      ...album,
      pageTitle: CONFIG.page_title,
      util,
    });
  }
  const url = await fetchAlbumURL(albumID);
  return h.view('bare-album', {
    url,
    pageTitle: CONFIG.page_title,
    util,
  });
};

export const handleUser = async (request: Hapi.Request, h: Hapi.ResponseToolkit) => {
  // https://imgur.com/user/MomBotNumber5
  if (!CONFIG.use_api) {
    return 'User page disabled. Rimgu administrator needs to enable API for this to work.';
  }
  const userID = request.params.userID;
  const userPosts = await fetchUserPosts(userID);
  return h.view('user-posts', {
    userPosts,
    pageTitle: CONFIG.page_title,
    util,
  });
};

export const handleTag = (request: Hapi.Request, h: Hapi.ResponseToolkit) => {
  // https://imgur.com/t/funny
  throw new Error('not implemented');
};

export const handleGallery = async (request: Hapi.Request, h: Hapi.ResponseToolkit) => {
  const galleryID = request.params.galleryID;
  const gallery = await fetchGallery(galleryID);
  const comments = CONFIG.use_api
    ? await fetchComments(galleryID)
    : null;
  return h.view('gallery', {
    ...gallery,
    comments,
    pageTitle: CONFIG.page_title,
    util,
  });
};
