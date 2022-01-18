export const handleUser = async (request: Hapi.Request, h: Hapi.ResponseToolkit) => {
  // https://imgur.com/user/MomBotNumber5
  if (!CONFIG.use_api) {
    return 'User page disabled. Rimgu administrator needs to enable API for this to work.';
  }
  const userID = request.params.userID;
  const user = await fetchUserInfo(userID);
  const posts = await fetchUserPosts(userID);
  return h.view('posts', {
    posts,
    user,
    pageTitle: CONFIG.page_title,
    util,
  });
};

export const handleUserCover = async (request: Hapi.Request, h: Hapi.ResponseToolkit) => {
  const userID = request.params.userID;
  const result = await fetchMedia(`/user/${userID}/cover?maxwidth=2560`);
  const response = h.response(result.rawBody)
    .header('Content-Type', result.headers["content-type"] || `image/jpeg`);
  return response;
};

export const handleTag = async (request: Hapi.Request, h: Hapi.ResponseToolkit) => {
  // https://imgur.com/t/funny
  if (!CONFIG.use_api) {
    return 'Tag page disabled. Rimgu administrator needs to enable API for this to work.';
  }
  const tagID = request.params.tagID;
  const result = await fetchTagPosts(tagID);
  return h.view('posts', {
    posts: result.items,
    pageTitle: CONFIG.page_title,
    tag: result,
    util,
  });
};