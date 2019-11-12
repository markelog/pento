import fetch from "isomorphic-unfetch";

const API = process.env.API;

function decorate(tracks) {
  return tracks.map(track => {
    return {
      id: track.id,
      title: track.name,
      start: new Date(track.start),
      stop: new Date(track.stop)
    };
  });
}

export function getTracks(email) {
  return fetch(`${API}/tracks/${email}`)
    .then(r => r.json())
    .then(data => decorate(data.payload));
}
