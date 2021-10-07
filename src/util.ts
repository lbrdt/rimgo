export const proxyURL = (url: string): string =>
  url.replace(/^https?:\/\/[^.]*\.imgur.com\//, '/');

export const linkify = (content: string) =>
  content.replace(
    /https?:\/\/[^.]*\.imgur.com\/([\/_a-zA-Z0-9-]+)\.(gifv|mp4)/g,
    '<video src="/$1.mp4" class="commentVideo commentObject" loop="" autoplay="" controls=""></video>'
  ).replace(
    /https?:\/\/[^.]*\.imgur.com\/([\/_a-zA-Z0-9-]+\.[a-z0-9A-Z]{2,6})/g,
    '<a href="/$1" target="_blank"><img class="commentImage commentObject" src="/$1" loading="lazy" /></a>'
  );
