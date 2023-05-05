export const getCurrentHost = () => `${window.location.protocol}//${window.location.host}`;

export const getParameterByName = (name) => {
  let url = window.location.href;

  name = name.replace(/[\[\]]/g, "\\$&");
  var regex = new RegExp("[?&]" + name + "(=([^&#]*)|&|#|$)"),
    results = regex.exec(url);
  if (!results) return null;
  if (!results[2]) return '';

  return decodeURIComponent(results[2].replace(/\+/g, " "));
};

export function randomDomId() {
  var length = 9;
  var id = Math.random().toString(36).replace(/[^a-z]+/g, '').substr(0, length);
  return id;
};
