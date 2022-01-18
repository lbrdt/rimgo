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