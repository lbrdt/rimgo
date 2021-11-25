import cheerio from 'cheerio';
import got, { Response } from 'got';
import { HttpsProxyAgent, HttpProxyAgent } from 'hpagent';
import { globalAgent as httpGlobalAgent } from 'http';
import { globalAgent as httpsGlobalAgent } from 'https';

import CONFIG from './config';

const GALLERY_JSON_REGEX = /window\.postDataJSON=(".*")$/;

const agent = {
  http: CONFIG.http_proxy
    ? new HttpProxyAgent({
      keepAlive: true,
      keepAliveMsecs: 1000,
      maxSockets: 256,
      maxFreeSockets: 256,
      scheduling: 'lifo',
      proxy: CONFIG.http_proxy,
    })
    : httpGlobalAgent,
  https: CONFIG.https_proxy
    ? new HttpsProxyAgent({
      keepAlive: true,
      keepAliveMsecs: 1000,
      maxSockets: 256,
      maxFreeSockets: 256,
      scheduling: 'lifo',
      proxy: CONFIG.https_proxy,
    })
    : httpsGlobalAgent
};

const get = async (url: string): Promise<Response<string>> => {
  try {
    return got(url, { agent });
  }
  catch (err) {
    console.error(`Error getting ${url}`);
    throw err;
  }
};

export const fetchAlbumURL = async (albumID: string): Promise<string> => {
  // https://imgur.com/a/DfEsrAB
  const response = await get(`https://imgur.com/a/${albumID}`);
  const $ = cheerio.load(response.body);
  const url = $('head meta[property="og:image"]').attr('content')?.replace(/\/\?.*$/, '');
  if (!url) {
    throw new Error(`Could not read image url for album ${albumID}`);
  }
  return url;
};

export const fetchAlbum = async (albumID: string): Promise<Media[]> => {
  // https://api.imgur.com/post/v1/albums/zk7mdKH?client_id=${CLIENT_ID}&include=media%2Caccount
  const response = await get(
    `https://api.imgur.com/post/v1/albums/${albumID}?client_id=${CONFIG.imgur_client_id}&include=media%2Caccount`,
  );
  return JSON.parse(response.body);
}

export const fetchComments = async (galleryID: string): Promise<Comment[]> => {
  /* eslint-disable max-len */
  // https://api.imgur.com/comment/v1/comments?client_id=${CLIENT_ID}%5Bpost%5D=eq%3Ag1bk7CB&include=account%2Cadconfig&per_page=30&sort=best
  const response = await get(
    `https://api.imgur.com/comment/v1/comments?client_id=${CONFIG.imgur_client_id}&filter%5Bpost%5D=eq%3A${galleryID}&include=account%2Cadconfig&per_page=30&sort=best`,
  );
  return JSON.parse(response.body).data;
  /* eslint-enable max-len */
}

export const fetchUserInfo = async (userID: string): Promise<UserResult> => {
  // https://api.imgur.com/account/v1/accounts/hughjaniss?client_id=${CLIENT_ID}
  const response = await get(
    `https://api.imgur.com/account/v1/accounts/${userID.toLowerCase()}?client_id=${CONFIG.imgur_client_id}&include=`,
  );
  return JSON.parse(response.body);
}

export const fetchUserPosts = async (userID: string, sort: Sorting = 'newest'): Promise<Post[]> => {
  /* eslint-disable max-len */
  // https://api.imgur.com/3/account/mombotnumber5/submissions/0/newest?album_previews=1&client_id=${CLIENT_ID}
  const response = await get(
    `https://api.imgur.com/3/account/${userID.toLowerCase()}/submissions/0/${sort}?album_previews=1&client_id=${CONFIG.imgur_client_id}`,
  );
  return JSON.parse(response.body).data;
  /* eslint-enable max-len */
}

export const fetchTagPosts = async (tagID: string, sort: Sorting = 'viral'): Promise<TagResult> => {
  /* eslint-disable max-len */
  // https://api.imgur.com/3/account/mombotnumber5/submissions/0/newest?album_previews=1&client_id=${CLIENT_ID}
  const response = await get(
    `https://api.imgur.com/3/gallery/t/${tagID.toLowerCase()}/${sort}/week/0?client_id=${CONFIG.imgur_client_id}`,
  );
  return JSON.parse(response.body).data;
  /* eslint-enable max-len */
}

export const fetchGallery = async (galleryID: string): Promise<Gallery> => {
  // https://imgur.com/gallery/g1bk7CB
  const response = await get(`https://imgur.com/gallery/${galleryID}`);
  const $ = cheerio.load(response.body);
  const postDataScript = $('head script:first-of-type').html();
  if (!postDataScript) {
    throw new Error('Could not find gallery data');
  }
  const postDataMatches = postDataScript.match(GALLERY_JSON_REGEX);
  if (!postDataMatches || postDataMatches.length < 2) {
    throw new Error('Could not parse gallery data');
  }
  const body = postDataMatches[1].replace(/\\'/g, "'");
  return JSON.parse(JSON.parse(body));
};

export const fetchMedia = async (filename: string): Promise<Response<string>> =>
  await get(`https://i.imgur.com/${filename}`);
