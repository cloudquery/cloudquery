export const assetPrefix = window.location.origin.includes('localhost')
  ? ''
  : window.location.href.split('/index.html')[0];
