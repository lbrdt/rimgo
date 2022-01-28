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