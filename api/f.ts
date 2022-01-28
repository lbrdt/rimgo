export const fetchUserPosts = async (userID: string, sort: Sorting = 'newest'): Promise<Post[]> => {
  /* eslint-disable max-len */
  // https://api.imgur.com/3/account/mombotnumber5/submissions/0/newest?album_previews=1&client_id=${CLIENT_ID}
  const response = await get(
    ``,
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